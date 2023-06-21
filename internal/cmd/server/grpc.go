package server

import (
	"github.com/Sokol111/category-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
}

func NewServer(port int, c proto.CategoryServiceServer) *Server {
	s := grpc.NewServer()
	proto.RegisterCategoryServiceServer(s, c)
	return &Server{
		grpcServer: s,
		port:       port,
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Println("Started grpc server on port", s.port)
	if err := s.grpcServer.Serve(ln); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
