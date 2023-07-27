//go:build wireinject

package di

import (
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/config"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/api"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/client"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/db"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/repository"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/service"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/usecase"
	"github.com/google/wire"
)

func InitiateAPI(cfg *config.Config) (*api.ServeServer, error) {
	wire.Build(

		db.ConnToDB,
		// Repository
		repository.NewAuthRepository,
		// Client
		client.NewUserServiceClient,

		// Usecase
		usecase.NewAuthUseCase,

		// Service
		service.NewAuthServiceServer,

		api.NewGRPCServer,
	)
	return &api.ServeServer{}, nil
}
