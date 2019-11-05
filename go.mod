module github.com/tektoncd/triggers

go 1.13

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.12.4 // indirect
	github.com/aws/aws-sdk-go v1.22.0 // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/gobuffalo/envy v1.7.0 // indirect
	github.com/google/go-cmp v0.3.0
	github.com/google/go-containerregistry v0.0.0-20191029173801-50b26ee28691 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/tektoncd/pipeline v0.6.0
	github.com/tektoncd/plumbing v0.0.0-20191026154524-1062bd15a735
	github.com/tidwall/gjson v1.3.2
	github.com/tidwall/sjson v1.0.4
	go.uber.org/zap v1.10.0
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api v0.7.1-0.20190805211801-b7b1a549a9ef // indirect
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/api v0.0.0-20190528110122-9ad12a4af326
	k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go v0.0.0-20180927013249-97dc22a90a29
	k8s.io/code-generator v0.0.0-20191026065352-f361089c127c
	knative.dev/caching v0.0.0-20190719140829-2032732871ff
	knative.dev/pkg v0.0.0-20190719141030-e4bc08cc8ded
)

replace github.com/tektoncd/plumbing => /usr/local/google/home/wlynch/src/github.com/tektoncd/plumbing
