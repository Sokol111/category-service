package main

import (
	"github.com/Sokol111/category-service/internal/cmd/server"
	"github.com/Sokol111/category-service/internal/config"
)

func main() {
	conf := config.LoadConfig(".")
	server.InitializeServer(conf.Grpc.Server.Port).Start()
}
