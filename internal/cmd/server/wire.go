//go:build wireinject
// +build wireinject

package server

import (
	"github.com/Sokol111/category-service/internal/core/services/categorysrv"
	"github.com/Sokol111/category-service/internal/handlers/cathandler"
	"github.com/Sokol111/category-service/internal/repositories/catrepo"
	"github.com/google/wire"
)

var repositories = wire.NewSet(catrepo.NewCategoryRepository)
var services = wire.NewSet(categorysrv.NewCategoryService)
var handlers = wire.NewSet(cathandler.NewCategoryServiceServer)
var server = wire.NewSet(NewServer)

func InitializeServer(port int) *Server {
	wire.Build(repositories, services, handlers, server)
	return nil
}
