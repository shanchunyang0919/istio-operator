module github.com/banzaicloud/istio-operator

go 1.14

require (
	github.com/Masterminds/semver/v3 v3.1.0
	github.com/banzaicloud/istio-client-go v0.0.9
	github.com/banzaicloud/k8s-objectmatcher v1.4.0
	github.com/caddyserver/caddy v1.0.5
	github.com/coreos/go-semver v0.3.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.3.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/goph/emperror v0.17.2
	github.com/lestrrat-go/jwx v1.0.6
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200627165143-92b8a710ab6c
	github.com/stretchr/testify v1.6.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.0.0-20200616195046-dc31b401abb5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	k8s.io/api v0.19.2
	k8s.io/apiextensions-apiserver v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	sigs.k8s.io/controller-runtime v0.7.0
)
