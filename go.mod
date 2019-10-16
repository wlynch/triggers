module github.com/tektoncd/triggers

go 1.13

require (
	cloud.google.com/go v0.43.0
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	contrib.go.opencensus.io/exporter/stackdriver v0.12.4
	github.com/aws/aws-sdk-go v1.22.0
	github.com/beorn7/perks v1.0.1
	github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/davecgh/go-spew v1.1.1
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/envy v1.7.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v1.0.0
	github.com/google/go-cmp v0.3.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0
	github.com/google/gofuzz v1.0.0
	github.com/google/licenseclassifier v0.0.0-20190711054124-c3068f13fcc3
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.3.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/hashicorp/golang-lru v0.5.3
	github.com/imdario/mergo v0.3.7
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.7
	github.com/knative/test-infra v0.0.0-20190625174906-69af8af1d3fe
	github.com/markbates/inflect v1.0.4
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/petar/GoLLRB v0.0.0-20190514000832-33fb24c13b99
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90
	github.com/prometheus/common v0.6.0
	github.com/prometheus/procfs v0.0.3
	github.com/rogpeppe/go-internal v1.3.0
	github.com/sergi/go-diff v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/tektoncd/pipeline v0.6.0
	github.com/tektoncd/plumbing v0.0.0-20190723185055-a268939d0be5
	github.com/tidwall/gjson v1.3.2
	github.com/tidwall/match v1.0.1
	github.com/tidwall/pretty v1.0.0
	github.com/tidwall/sjson v1.0.4
	go.opencensus.io v0.22.0
	go.uber.org/atomic v1.4.0
	go.uber.org/multierr v1.1.0
	go.uber.org/zap v1.9.2-0.20180814183419-67bc79d13d15
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools v0.0.0-20190807164442-cae9aa543496
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api v0.7.1-0.20190805211801-b7b1a549a9ef
	google.golang.org/appengine v1.6.1
	google.golang.org/genproto v0.0.0-20190716160619-c506a9f90610
	google.golang.org/grpc v1.22.1
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190528110122-9ad12a4af326
	k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go v0.0.0-20180927013249-97dc22a90a29
	k8s.io/code-generator v0.0.0-20181128191024-b1289fc74931
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a
	k8s.io/klog v0.2.0
	k8s.io/kube-openapi v0.0.0-20190722073852-5e22f3d471e6
	knative.dev/caching v0.0.0-20190719140829-2032732871ff
	knative.dev/pkg v0.0.0-20190719141030-e4bc08cc8ded
)
