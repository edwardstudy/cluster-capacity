/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	externalclientset "k8s.io/client-go/kubernetes"
	fakeclientset "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/events"
	schedconfig "k8s.io/kubernetes/cmd/kube-scheduler/app/config"
	schedoptions "k8s.io/kubernetes/cmd/kube-scheduler/app/options"
	"k8s.io/kubernetes/pkg/scheduler"
	schedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"

	uuid "github.com/satori/go.uuid"
	localCache "sigs.k8s.io/cluster-capacity/pkg/cache"
	"sigs.k8s.io/cluster-capacity/pkg/framework/strategy"
)

const (
	podProvisioner = "cc.kubernetes.io/provisioned-by"
)

type ClusterCapacity struct {
	// emulation strategy
	strategy strategy.Strategy

	externalkubeclient externalclientset.Interface
	informerFactory    informers.SharedInformerFactory

	// schedulers
	schedulers           map[string]*scheduler.Scheduler
	defaultSchedulerName string
	defaultSchedulerConf *schedconfig.CompletedConfig
	// pod to schedule
	simulatedPod     *corev1.Pod
	lastSimulatedPod *corev1.Pod
	maxSimulated     int
	simulated        int
	status           Status
	report           *ClusterCapacityReview

	// analysis limitation
	informerStopCh chan struct{}
	// schedulers channel
	schedulerCh chan struct{}

	// stop the analysis
	stop      chan struct{}
	stopMux   sync.RWMutex
	stopped   bool
	closedMux sync.RWMutex
	closed    bool

	// localCache provider caching for result of syncing
	localCache localCache.Cache
}

// capture all scheduled pods with reason why the analysis could not continue
type Status struct {
	Pods       []*corev1.Pod
	StopReason string
}

func (c *ClusterCapacity) Report() *ClusterCapacityReview {
	if c.report == nil {
		// Preparation before pod sequence scheduling is done
		pods := make([]*corev1.Pod, 0)
		pods = append(pods, c.simulatedPod)
		c.report = GetReport(pods, c.status)
		c.report.Spec.Replicas = int32(c.maxSimulated)
	}

	return c.report
}

type resources struct {
	PodItems  []corev1.Pod
	NodeItems []corev1.Node
	SVCItems  []corev1.Service
	PVCItems  []corev1.PersistentVolumeClaim
	RCItems   []corev1.ReplicationController
	PDBItems  []v1beta1.PodDisruptionBudget
	RSItems   []appsv1.ReplicaSet
	SSItems   []appsv1.StatefulSet
	SCItems   []storagev1.StorageClass
}

// SyncResources sync resources from API server if no cache exists or refresh needed.
// Otherwise using local cache data.
func (c *ClusterCapacity) SyncResources(client externalclientset.Interface, refresh bool) error {
	var res resources
	var err error
	var needCached bool
	if !refresh {
		// Read from cache
		res, err = c.getResourcesFromCache()
		if err != nil {
			refresh = true
		}
		if reflect.DeepEqual(res, resources{}) {
			refresh = true
		}
	}

	if refresh {
		resUpdated, err := getResourcesFromAPIServer(client)
		if err != nil {
			return fmt.Errorf("unable to list pods: %v", err)
		}
		if !reflect.DeepEqual(res, resUpdated) {
			res = resUpdated
			needCached = true
		}
	}

	err = c.FillWithResources(res)
	if err != nil {
		return fmt.Errorf("unable to fill with resources: %v", err)
	}

	if needCached {
		data, err := json.Marshal(&res)
		if err != nil {
			return fmt.Errorf("unable to marshal items: %v", err)
		}
		err = c.localCache.Set("resources", data)
		if err != nil {
			return fmt.Errorf("unable to cache")
		}
	}

	return nil
}

func (c *ClusterCapacity) FillWithResources(res resources) error {
	// Fill into resources thru externalkubeclient
	for _, item := range res.PodItems {
		if _, err := c.externalkubeclient.CoreV1().Pods(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy pod: %v", err)
		}
	}

	for _, item := range res.NodeItems {
		if _, err := c.externalkubeclient.CoreV1().Nodes().Create(&item); err != nil {
			return fmt.Errorf("unable to copy node: %v", err)
		}
	}

	for _, item := range res.SVCItems {
		if _, err := c.externalkubeclient.CoreV1().Services(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy service: %v", err)
		}
	}

	for _, item := range res.PVCItems {
		if _, err := c.externalkubeclient.CoreV1().PersistentVolumeClaims(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy pvc: %v", err)
		}
	}

	for _, item := range res.RCItems {
		if _, err := c.externalkubeclient.CoreV1().ReplicationControllers(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy RC: %v", err)
		}
	}

	for _, item := range res.PDBItems {
		if _, err := c.externalkubeclient.PolicyV1beta1().PodDisruptionBudgets(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy PDB: %v", err)
		}
	}

	for _, item := range res.RSItems {
		if _, err := c.externalkubeclient.AppsV1().ReplicaSets(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy replica set: %v", err)
		}
	}

	for _, item := range res.SSItems {
		if _, err := c.externalkubeclient.AppsV1().StatefulSets(item.Namespace).Create(&item); err != nil {
			return fmt.Errorf("unable to copy stateful set: %v", err)
		}
	}

	for _, item := range res.SCItems {
		if _, err := c.externalkubeclient.StorageV1().StorageClasses().Create(&item); err != nil {
			return fmt.Errorf("unable to copy storage class: %v", err)
		}
	}

	return nil
}

func (c *ClusterCapacity) Bind(ctx context.Context, state *framework.CycleState, p *corev1.Pod, nodeName string, schedulerName string) *framework.Status {
	// run the pod through strategy
	pod, err := c.externalkubeclient.CoreV1().Pods(p.Namespace).Get(p.Name, metav1.GetOptions{})
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Unable to bind: %v", err))
	}
	updatedPod := pod.DeepCopy()
	updatedPod.Spec.NodeName = nodeName
	updatedPod.Status.Phase = corev1.PodRunning

	// TODO(jchaloup): rename Add to Update as this actually updates the scheduled pod
	if err := c.strategy.Add(updatedPod); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Unable to recompute new cluster state: %v", err))
	}

	c.status.Pods = append(c.status.Pods, updatedPod)

	if c.maxSimulated > 0 && c.simulated >= c.maxSimulated {
		c.status.StopReason = fmt.Sprintf("LimitReached: Maximum number of pods simulated: %v", c.maxSimulated)
		c.Close()
		c.stop <- struct{}{}
		return nil
	}

	// all good, create another pod
	if err := c.nextPod(); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Unable to create next pod to schedule: %v", err))
	}
	return nil
}

func (c *ClusterCapacity) Close() {
	c.closedMux.Lock()
	defer c.closedMux.Unlock()

	if c.closed {
		return
	}

	close(c.schedulerCh)
	close(c.informerStopCh)
	c.closed = true
}

func (c *ClusterCapacity) Update(pod *corev1.Pod, podCondition *corev1.PodCondition, schedulerName string) error {
	stop := podCondition.Type == corev1.PodScheduled && podCondition.Status == corev1.ConditionFalse && podCondition.Reason == "Unschedulable"

	// Only for pending pods provisioned by cluster-capacity
	if stop && metav1.HasAnnotation(pod.ObjectMeta, podProvisioner) {
		c.status.StopReason = fmt.Sprintf("%v: %v", podCondition.Reason, podCondition.Message)
		c.Close()
		// The Update function can be run more than once before any corresponding
		// scheduler is closed. The behaviour is implementation specific
		c.stopMux.Lock()
		defer c.stopMux.Unlock()
		c.stopped = true
		c.stop <- struct{}{}
	}
	return nil
}

func (c *ClusterCapacity) nextPod() error {
	pod := corev1.Pod{}
	pod = *c.simulatedPod.DeepCopy()
	// reset any node designation set
	pod.Spec.NodeName = ""
	// use simulated pod name with an index to construct the name
	pod.ObjectMeta.Name = fmt.Sprintf("%v-%v", c.simulatedPod.Name, c.simulated)
	pod.ObjectMeta.UID = types.UID(uuid.NewV4().String())
	pod.Spec.SchedulerName = c.defaultSchedulerName

	// Add pod provisioner annotation
	if pod.ObjectMeta.Annotations == nil {
		pod.ObjectMeta.Annotations = map[string]string{}
	}

	// Stores the scheduler name
	pod.ObjectMeta.Annotations[podProvisioner] = c.defaultSchedulerName

	c.simulated++
	c.lastSimulatedPod = &pod

	_, err := c.externalkubeclient.CoreV1().Pods(pod.Namespace).Create(&pod)
	return err
}

func (c *ClusterCapacity) getResourcesFromCache() (resources, error) {
	// Read from cache
	var res resources
	data, err := c.localCache.Get("resources")
	if err != nil {
		return res, fmt.Errorf("unable to get local cache: %v", err)
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, fmt.Errorf("unable to unmarshal cache: %v", err)
	}

	return res, nil
}

func getResourcesFromAPIServer(client externalclientset.Interface) (resources, error) {
	var res resources
	pods, err := client.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list pods: %v", err)
	}
	res.PodItems = pods.Items

	nodes, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list nodes: %v", err)
	}
	res.NodeItems = nodes.Items

	svcs, err := client.CoreV1().Services(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list services: %v", err)
	}
	res.SVCItems = svcs.Items

	pvcs, err := client.CoreV1().PersistentVolumeClaims(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list pvcs: %v", err)
	}
	res.PVCItems = pvcs.Items

	rcs, err := client.CoreV1().ReplicationControllers(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list RCs: %v", err)
	}
	res.RCItems = rcs.Items

	pdbs, err := client.PolicyV1beta1().PodDisruptionBudgets(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list PDBs: %v", err)
	}
	res.PDBItems = pdbs.Items

	replicaSetItems, err := client.AppsV1().ReplicaSets(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list replicas sets: %v", err)
	}
	res.RSItems = replicaSetItems.Items

	statefulSetItems, err := client.AppsV1().StatefulSets(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list stateful sets: %v", err)
	}
	res.SSItems = statefulSetItems.Items

	storageClassesItems, err := client.StorageV1().StorageClasses().List(metav1.ListOptions{})
	if err != nil {
		return res, fmt.Errorf("unable to list storage classes: %v", err)
	}
	res.SCItems = storageClassesItems.Items

	return res, nil
}

func (c *ClusterCapacity) Run() error {
	// Start all informers.
	c.informerFactory.Start(c.informerStopCh)
	c.informerFactory.WaitForCacheSync(c.informerStopCh)

	ctx, cancel := context.WithCancel(context.Background())

	// TODO(jchaloup): remove all pods that are not scheduled yet
	for _, scheduler := range c.schedulers {
		go func() {
			scheduler.Run(ctx)
		}()
	}
	// wait some time before at least nodes are populated
	// TODO(jchaloup); find a better way how to do this or at least decrease it to <100ms
	time.Sleep(100 * time.Millisecond)
	// create the first simulated pod
	err := c.nextPod()
	if err != nil {
		cancel()
		c.Close()
		close(c.stop)
		return fmt.Errorf("Unable to create next pod to schedule: %v", err)
	}
	<-c.stop
	cancel()
	close(c.stop)
	return nil
}

type localBinderPodConditionUpdater struct {
	schedulerName string
	c             *ClusterCapacity
}

func (b *localBinderPodConditionUpdater) Name() string {
	return "ClusterCapacityBinder"
}

// TODO(jchaloup): Needs to be locked since the scheduler runs the binding phase in a go routine
func (b *localBinderPodConditionUpdater) Bind(ctx context.Context, state *framework.CycleState, p *corev1.Pod, nodeName string) *framework.Status {
	return b.c.Bind(ctx, state, p, nodeName, b.schedulerName)
}

func (c *ClusterCapacity) NewBindPlugin(schedulerName string, configuration runtime.Object, f framework.FrameworkHandle) (framework.Plugin, error) {
	return &localBinderPodConditionUpdater{
		schedulerName: schedulerName,
		c:             c,
	}, nil
}

func (c *ClusterCapacity) createScheduler(schedulerName string, cc *schedconfig.CompletedConfig) (*scheduler.Scheduler, error) {
	outOfTreeRegistry := framework.Registry{
		"ClusterCapacityBinder": func(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
			return c.NewBindPlugin(schedulerName, configuration, f)
		},
	}
	plugins := &schedulerconfig.Plugins{
		Bind: &schedulerconfig.PluginSet{
			Enabled: []schedulerconfig.Plugin{
				{
					Name: "ClusterCapacityBinder",
				},
			},
			Disabled: []schedulerconfig.Plugin{
				{
					Name: "DefaultBinder",
				},
			},
		},
	}

	c.informerFactory.Core().V1().Pods().Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: func(obj interface{}) bool {
				if pod, ok := obj.(*corev1.Pod); ok && pod.Spec.SchedulerName == schedulerName {
					return true
				}
				return false
			},
			Handler: cache.ResourceEventHandlerFuncs{
				UpdateFunc: func(oldObj, newObj interface{}) {
					if pod, ok := newObj.(*corev1.Pod); ok {
						for _, podCondition := range pod.Status.Conditions {
							if podCondition.Type == corev1.PodScheduled {
								c.Update(pod, &podCondition, schedulerName)
							}
						}
					}
				},
			},
		},
	)

	// Create the scheduler.
	return scheduler.New(
		c.externalkubeclient,
		c.informerFactory,
		c.informerFactory.Core().V1().Pods(),
		cc.Broadcaster.NewRecorder(scheme.Scheme, c.defaultSchedulerName),
		c.schedulerCh,
		scheduler.WithFrameworkPlugins(plugins),
		scheduler.WithAlgorithmSource(cc.ComponentConfig.AlgorithmSource),
		scheduler.WithPercentageOfNodesToScore(cc.ComponentConfig.PercentageOfNodesToScore),
		scheduler.WithFrameworkOutOfTreeRegistry(outOfTreeRegistry),
		scheduler.WithPodMaxBackoffSeconds(cc.ComponentConfig.PodMaxBackoffSeconds),
		scheduler.WithPodInitialBackoffSeconds(cc.ComponentConfig.PodInitialBackoffSeconds),
	)
	// extender from schedulerAlgorithmSource.Policy file
}

// TODO(avesh): enable when support for multiple schedulers is added.
/*func (c *ClusterCapacity) AddScheduler(s *sapps.SchedulerServer) error {
	scheduler, err := c.createScheduler(s)
	if err != nil {
		return err
	}

	c.schedulers[s.SchedulerName] = scheduler
	return nil
}*/

// Create new cluster capacity analysis
// The analysis is completely independent of apiserver so no need
// for kubeconfig nor for apiserver url
func New(kubeSchedulerConfig *schedconfig.CompletedConfig, simulatedPod *corev1.Pod, maxPods int, cache localCache.Cache) (*ClusterCapacity, error) {
	client := fakeclientset.NewSimpleClientset()
	sharedInformerFactory := informers.NewSharedInformerFactory(client, 0)

	kubeSchedulerConfig.Client = client

	cc := &ClusterCapacity{
		strategy:           strategy.NewPredictiveStrategy(client),
		externalkubeclient: client,
		simulatedPod:       simulatedPod,
		simulated:          0,
		maxSimulated:       maxPods,
		stop:               make(chan struct{}),
		informerFactory:    sharedInformerFactory,
		informerStopCh:     make(chan struct{}),
		schedulerCh:        make(chan struct{}),
		localCache:         cache,
	}

	cc.schedulers = make(map[string]*scheduler.Scheduler)

	scheduler, err := cc.createScheduler(corev1.DefaultSchedulerName, kubeSchedulerConfig)
	if err != nil {
		return nil, err
	}

	cc.schedulers[corev1.DefaultSchedulerName] = scheduler
	cc.defaultSchedulerName = corev1.DefaultSchedulerName
	return cc, nil
}

func InitKubeSchedulerConfiguration(opts *schedoptions.Options) (*schedconfig.CompletedConfig, error) {
	c := &schedconfig.Config{}
	// clear out all unnecessary options so no port is bound
	// to allow running multiple instances in a row
	opts.Deprecated = nil
	opts.CombinedInsecureServing = nil
	opts.SecureServing = nil
	if err := opts.ApplyTo(c); err != nil {
		return nil, fmt.Errorf("unable to get scheduler config: %v", err)
	}

	// Get the completed config
	cc := c.Complete()

	// completely ignore the events
	cc.Broadcaster = events.NewBroadcaster(&events.EventSinkImpl{
		Interface: fakeclientset.NewSimpleClientset().EventsV1beta1().Events(corev1.NamespaceAll),
	})

	return &cc, nil
}
