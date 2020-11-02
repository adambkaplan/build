module github.com/shipwright-io/build

go 1.13

require (
	cloud.google.com/go v0.65.0 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.7.1-0.20200615190824-f8c219d2d895 // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.2.1-0.20200609204449-6bcf6f8577f0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.13.2 // indirect
	github.com/Azure/go-autorest/autorest v0.10.2 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.3 // indirect
	github.com/aws/aws-sdk-go v1.31.12 // indirect
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/docker/cli v0.0.0-20200210162036-a4bedce16568 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/logr v0.2.0
	github.com/go-openapi/spec v0.19.6
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.1-0.20191009090205-6c0755d89d1e // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/operator-framework/operator-sdk v0.18.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/client_model v0.2.0
	github.com/rogpeppe/go-internal v1.5.2 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/tektoncd/pipeline v0.11.3
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sys v0.0.0-20200828081204-131dc92a58d5 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.0.0-20200828013309-97019fc2e64b // indirect
	google.golang.org/genproto v0.0.0-20200828030656-73b5761be4c5 // indirect
	google.golang.org/grpc v1.31.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.18.10
	k8s.io/apimachinery v0.18.10
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.18.10
	k8s.io/gengo v0.0.0-20200205140755-e0e292d8aa12 // indirect
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	k8s.io/kubectl v0.18.10
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451
	knative.dev/pkg v0.0.0-20200207155214-fef852970f43
	sigs.k8s.io/controller-runtime v0.6.1
	sigs.k8s.io/structured-merge-diff/v3 v3.0.1-0.20200706213357-43c19bbb7fba // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	// github.com/operator-framework/operator-registry requires Sirupsen/logrus@v1.7.0
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.7.0
	// latest tag resolves to a very old version. this is only used for spinning up local test registries
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/go-logr/logr => github.com/go-logr/logr v0.1.0 // v0.2.0 release doesn't have logr.InfoLogger
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.10
	k8s.io/client-go => k8s.io/client-go v0.18.10 // Required by prometheus-operator
	k8s.io/code-generator => k8s.io/code-generator v0.18.10
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // resolve `case-insensitive import collision` for gnostic/openapiv2 package
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.1

)
