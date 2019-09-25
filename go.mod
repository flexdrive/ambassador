module github.com/datawire/ambassador

go 1.13

require (
	github.com/envoyproxy/go-control-plane v0.5.0
	github.com/envoyproxy/protoc-gen-validate v0.0.15-0.20190405222122-d6164de49109
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gogo/googleapis v1.1.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.2.1-0.20190205222052-c823c79ea157 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/lyft/protoc-gen-validate v0.0.0-20180626203901-f9d2b11e4414 // indirect
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/sirupsen/logrus v1.0.6
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac // indirect
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 // indirect
	google.golang.org/grpc v1.19.1
	// google.golang.org/grpc v1.14.0
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	istio.io/gogo-genproto v0.0.0-20190614210408-e88dc8b0e4db
)

// Pin versions of external commands (i.e. things that we don't use as
// libraries).  We explicitly 'replace' them so that `go mod tidy`
// can't make the file forget which version we want.
replace (
	// Go a few commits after v0.0.14 to get the
	// github.com/{lyft→envoyproxy}/protoc-gen-validate rename
	github.com/envoyproxy/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate v0.0.15-0.20190405222122-d6164de49109

	github.com/gogo/protobuf => github.com/gogo/protobuf v1.2.1
)
