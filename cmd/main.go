package main

import (
	"github.com/Sokol111/category-service/cmd/server"
	"github.com/Sokol111/category-service/internal/api"
	"github.com/Sokol111/category-service/internal/repository"
	"github.com/Sokol111/category-service/internal/service"
)

func main() {
	conf := server.LoadConfig(".")
	catRepo := repository.NewCatInmemoryRepo()
	catService := service.NewCatService(catRepo)
	catHandler := api.NewCategoryHandler(catService)
	s := server.NewGrpcServer(conf.Grpc.Server.Port, catHandler)
	s.Start()
}
