module kratos-demo

go 1.13

require (
	github.com/go-kratos/kratos v1.0.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.5.0
	github.com/pkg/errors v0.9.1 // indirect
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0