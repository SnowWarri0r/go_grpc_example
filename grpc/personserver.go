package grpc

import (
	"google.golang.org/grpc"
	proto "grpc_test/genproto"
	"log"
	"net"
)

func Init() {
	port := ":8000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := grpc.NewServer()
	serviceRegister(service)
	log.Printf("rpc server running on %v", port)
	if err = service.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

// serviceRegister 注册gRPC服务
func serviceRegister(service *grpc.Server) {
	proto.RegisterPersonServiceServer(service, &PersonService{})
}
