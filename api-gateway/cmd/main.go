package main

import (
	"log"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeApi(&cfg)
	if err != nil {
		log.Fatalf("failed initialize api error:%s", err.Error())
	}

	service.Start()
}
