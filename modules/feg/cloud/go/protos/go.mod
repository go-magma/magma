module github.com/go-magma/magma/modules/feg/cloud/go/protos

replace (
	github.com/go-magma/magma/gateway/go => ../../../../../gateway/go
	github.com/go-magma/magma/lib/go => ../../../../../lib/go
	github.com/go-magma/magma/lib/go/protos => ../../../../../lib/go/protos
	github.com/go-magma/magma/orc8r/cloud/go => ../../../../../orc8r/cloud/go
)

require (
	github.com/go-magma/magma/lib/go/protos v0.0.0
	github.com/go-magma/magma/modules/lte/cloud/go v0.0.0-20200818162846-8db662db3aa7
	github.com/golang/protobuf v1.4.1
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.25.0
)

go 1.13
