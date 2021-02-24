module sigs.k8s.io/cluster-capacity

go 1.12

require (
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/emicklei/go-restful v2.13.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/lithammer/dedent v1.1.0
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/prometheus/common v0.11.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	go.etcd.io/etcd v0.0.0-20200716221620-18dfb9cca345 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200813001606-1ccf2a5ae4fd // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.21.0-alpha.0
	k8s.io/apiserver v0.20.2 // indirect
	k8s.io/client-go v0.20.2
	k8s.io/cloud-provider v0.20.2 // indirect
	k8s.io/component-base v0.20.2
	k8s.io/csi-translation-lib v0.20.2 // indirect
	k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd // indirect
	k8s.io/kube-scheduler v0.20.2
	k8s.io/kubernetes v1.20.2
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2 // Fixes https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-3121
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.7-rc.0
	k8s.io/apiserver => k8s.io/apiserver v0.17.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
	k8s.io/component-base => k8s.io/component-base v0.17.6
	k8s.io/cri-api => k8s.io/cri-api v0.17.13-rc.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.6
	k8s.io/klog => k8s.io/klog v1.0.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.6
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // release-1.17
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.6
	k8s.io/kubectl => k8s.io/kubectl v0.17.6
	k8s.io/kubelet => k8s.io/kubelet v0.17.6
	k8s.io/kubernetes => k8s.io/kubernetes v1.17.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.6
	k8s.io/metrics => k8s.io/metrics v0.17.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.6
	k8s.io/utils => k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
)
