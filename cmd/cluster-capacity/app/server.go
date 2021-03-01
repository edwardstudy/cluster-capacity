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

package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	aflag "k8s.io/component-base/cli/flag"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
	kubeschedulerconfigv1alpha1 "k8s.io/kube-scheduler/config/v1alpha1"
	schedconfig "k8s.io/kubernetes/cmd/kube-scheduler/app/config"
	schedoptions "k8s.io/kubernetes/cmd/kube-scheduler/app/options"
	_ "k8s.io/kubernetes/pkg/scheduler/algorithmprovider"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	kubeschedulerscheme "k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	utilpointer "k8s.io/utils/pointer"

	"sigs.k8s.io/cluster-capacity/cmd/cluster-capacity/app/options"
	"sigs.k8s.io/cluster-capacity/pkg/cache"
	"sigs.k8s.io/cluster-capacity/pkg/framework"
	"sigs.k8s.io/cluster-capacity/pkg/utils"
)

var (
	clusterCapacityLong = dedent.Dedent(`
		Cluster-capacity simulates an API server with initial state copied from the Kubernetes enviroment
		with its configuration specified in KUBECONFIG. The simulated API server tries to schedule the number of
		pods specified by --max-limits flag. If the --max-limits flag is not specified, pods are scheduled until
		the simulated API server runs out of resources.
	`)
)

func NewClusterCapacityCommand() *cobra.Command {
	opt := options.NewClusterCapacityOptions()
	cmd := &cobra.Command{
		Use:   "cluster-capacity --kubeconfig KUBECONFIG --podspec PODSPEC",
		Short: "Cluster-capacity is used for simulating scheduling of one or multiple pods",
		Long:  clusterCapacityLong,
		Run: func(cmd *cobra.Command, args []string) {
			err := Validate(opt)
			if err != nil {
				fmt.Println(err)
				cmd.Help()
				return
			}
			err = Run(opt)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.SetNormalizeFunc(aflag.WordSepNormalizeFunc)
	flags.AddGoFlagSet(flag.CommandLine)
	opt.AddFlags(flags)

	return cmd
}

func Validate(opt *options.ClusterCapacityOptions) error {
	if len(opt.PodSpecFile) == 0 {
		return fmt.Errorf("Pod spec file is missing")
	}

	_, present := os.LookupEnv("CC_INCLUSTER")
	if !present {
		if len(opt.Kubeconfig) == 0 {
			return fmt.Errorf("kubeconfig is missing")
		}
	}
	return nil
}

func Run(opt *options.ClusterCapacityOptions) error {
	conf := options.NewClusterCapacityConfig(opt)

	versionedCfg := kubeschedulerconfigv1alpha1.KubeSchedulerConfiguration{}
	versionedCfg.DebuggingConfiguration = configv1alpha1.DebuggingConfiguration{
		EnableProfiling: utilpointer.BoolPtr(true), // profile debugging is cheap to have exposed and standard on kube binaries
	}

	kubeschedulerscheme.Scheme.Default(&versionedCfg)
	kcfg := kubeschedulerconfig.KubeSchedulerConfiguration{}
	if err := kubeschedulerscheme.Scheme.Convert(&versionedCfg, &kcfg, nil); err != nil {
		return err
	}

	// Always set the list of bind plugins to ClusterCapacityBinder
	kcfg.SchedulerName = v1.DefaultSchedulerName
	if kcfg.Plugins == nil {
		kcfg.Plugins = &kubeschedulerconfig.Plugins{}
	}

	// TODO configurable plugins
	kcfg.Plugins.Bind = &kubeschedulerconfig.PluginSet{
		Enabled:  []kubeschedulerconfig.Plugin{{Name: "ClusterCapacityBinder"}},
		Disabled: []kubeschedulerconfig.Plugin{{Name: "DefaultBinder"}},
	}

	opts := &schedoptions.Options{
		ComponentConfig: kcfg,
		ConfigFile:      conf.Options.DefaultSchedulerConfigFile,
	}

	cc, err := framework.InitKubeSchedulerConfiguration(opts)
	if err != nil {
		return fmt.Errorf("failed to init kube scheduler configuration: %v ", err)
	}

	err = conf.ParseAPISpec(v1.DefaultSchedulerName)
	if err != nil {
		return fmt.Errorf("Failed to parse pod spec file: %v ", err)
	}

	var cfg *restclient.Config
	if len(conf.Options.Kubeconfig) != 0 {
		master, err := utils.GetMasterFromKubeConfig(conf.Options.Kubeconfig)
		if err != nil {
			return fmt.Errorf("Failed to parse kubeconfig file: %v ", err)
		}

		cfg, err = clientcmd.BuildConfigFromFlags(master, conf.Options.Kubeconfig)
		if err != nil {
			return fmt.Errorf("Unable to build config: %v", err)
		}

	} else {
		cfg, err = restclient.InClusterConfig()
		if err != nil {
			return fmt.Errorf("Unable to build in cluster config: %v", err)
		}
	}

	conf.KubeClient, err = clientset.NewForConfig(cfg)
	if err != nil {
		return err
	}

	conf.Cache, err = cache.NewDiskCache(opt.CacheDir)
	if err != nil {
		return fmt.Errorf("unable to create disk cache: %v", err)
	}

	report, err := runSimulator(conf, cc)
	if err != nil {
		return err
	}

	//nodes, _ := cc.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	//fmt.Printf("nodes: %v", nodes)
	if err := framework.ClusterCapacityReviewPrint(report, conf.Options.Verbose, conf.Options.OutputFormat); err != nil {
		return fmt.Errorf("Error while printing: %v", err)
	}
	return nil
}

func runSimulator(s *options.ClusterCapacityConfig, kubeSchedulerConfig *schedconfig.CompletedConfig) (*framework.ClusterCapacityReview, error) {
	cc, err := framework.New(kubeSchedulerConfig, s.Pod, s.Options.MaxLimit, s.Cache)
	if err != nil {
		return nil, err
	}

	err = cc.SyncResources(s.KubeClient, s.Options.Refresh)
	if err != nil {
		return nil, err
	}

	// TODO sync with metrics server, store in local cache that serves as loadCache.

	err = cc.Run()
	if err != nil {
		return nil, err
	}

	report := cc.Report()
	return report, nil
}
