module micro-hello

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/v2 v2.0.0 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
replace github.com/micro/go-micro/v2 v2.9.1 => ./v2@v2.9.1