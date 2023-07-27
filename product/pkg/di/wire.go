package di

import (
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/api"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/config"
)

func InitiateAPI(cfg *config.Config) (*api.ServeServer, error) {
	// wire.Build(

	// 	db.ConnToDB,
	// 	// Repository
	// 	repository.,

	// 	// Usecase
	// 	usecase.,

	// 	// Service
	// 	service.,

	// 	api.NewGRPCServer,
	// )
	return &api.ServeServer{}, nil
}
