//go:build wireinject

package di

import (
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/api"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/client"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/db"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/repository"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/service"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/usecase"
	"github.com/google/wire"
)

func InitiateAPI(cfg *config.Config) (*api.ServeServer, error) {
	wire.Build(

		db.ConnToDB,
		// Repository
		repository.NewUserRepository,
		// Client
		client.NewOrderClient,

		// Usecase
		usecase.NewUserUseCase,

		// Service
		service.NewUserServiceServer,

		api.NewGRPCServer,
	)
	return &api.ServeServer{}, nil
}
