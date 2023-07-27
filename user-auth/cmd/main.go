package main

import (
	"log"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/di"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Failed to config variables !")
	}
	service, err := di.InitiateAPI(cfg)
	if err != nil {
		log.Printf("\nfailed to initiate service error : %s", err.Error())
	}

	err = service.Run(cfg)
	if err != nil {
		log.Printf("\n failed to run the server")
	}
}
