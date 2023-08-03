//go:build wireinject

package di

import (
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/api"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/api/handler"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/google/wire"
)

func InitializeApi(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		client.NewAuthClient,
		// client.NewUserClient,
		// client.NewProductClient,
		// client.NewCartClient,
		// client.NewOrderClient,

		handler.NewAuthHandler,
		// handler.NewUserHandler,
		// handler.NewProductHandler,
		// handler.NewCartHandler,
		// handler.NewOrderHandler,

		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
