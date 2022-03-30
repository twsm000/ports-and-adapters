package grpc

import (
	"appstruct/internal/adapters/framework/left/grpc/pb"
	"appstruct/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (a Adapter) Run() {
	var err error

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("failed to listen on port 9000:", err)
	}

	arithmeticServiceServer := a
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("failed to serve gRPC over port 9000:", err)
	}
}
