module helloworld

go 1.13

// replace helloworld => ../proto

require (
	github.com/golang/protobuf v1.4.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.28.1
	google.golang.org/protobuf v1.21.0
)
