package server

import (
	"github.com/Sokol111/category-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type GrpcServer struct {
	grpcServer *grpc.Server
	port       int
}

func NewGrpcServer(port int, c proto.CategoryServiceServer) *GrpcServer {
	s := grpc.NewServer()
	proto.RegisterCategoryServiceServer(s, c)
	return &GrpcServer{
		grpcServer: s,
		port:       port,
	}
}

func (s *GrpcServer) Start() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Println("Started grpc server on port", s.port)
	if err := s.grpcServer.Serve(ln); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
