module code.cloudfoundry.org/groot-windows

go 1.19

require (
	code.cloudfoundry.org/filelock v0.0.0-20230410204127-470838d066c5
	code.cloudfoundry.org/groot v0.0.0-20230523195242-dfd19bf8bc3b
	code.cloudfoundry.org/hydrator v0.0.0-20230410213017-736b9a742d25
	code.cloudfoundry.org/lager/v3 v3.0.1
	github.com/Microsoft/go-winio v0.6.1
	github.com/Microsoft/hcsshim v0.9.9
	github.com/onsi/ginkgo/v2 v2.9.5
	github.com/onsi/gomega v1.27.7
	github.com/opencontainers/image-spec v1.1.0-rc3
	github.com/opencontainers/runtime-spec v1.1.0-rc.2
	github.com/urfave/cli v1.22.13
	golang.org/x/sys v0.8.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/containers/image/v5 v5.25.0 // indirect
	github.com/containers/libtrust v0.0.0-20230121012942-c1716e8a8d01 // indirect
	github.com/containers/ocicrypt v1.1.7 // indirect
	github.com/containers/storage v1.46.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230510103437-eeec1cb781c3 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/runc v1.1.7 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Microsoft/go-winio => github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5
	github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.8.7
	github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.1.0-rc2
)
