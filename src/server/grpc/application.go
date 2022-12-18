package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/blog-service/src/server"
	grpch "github.com/blog-service/src/server/grpc/handler"
	"github.com/blog-service/src/server/grpc/pb"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	bm *grpch.GrpcBeanManager
}

func NewGrpcServerAdapter(bm *grpch.GrpcBeanManager) server.ServerPort {
	return &GrpcServer{bm}
}

func (grpcs GrpcServer) StartApplication(addr string) {
	fmt.Println(addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on address %v with error: %v", addr, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, grpcs.bm.NewPostServiceServerHandler())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("cannot run gRPC server: %v", err)
	}
}
