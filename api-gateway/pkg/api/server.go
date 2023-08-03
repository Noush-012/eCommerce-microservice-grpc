package api

import (
	handler "github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/api/handler/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/api/routes"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	engine *gin.Engine
}

// NewServerHTTP creates a new server with given handler functions
func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/"), authHandler)
	// routes.SetupAdminRoutes(engine.Group("/admin"), productHandler)

	return &Server{
		engine: engine,
		port:   cfg.PORT,
	}
}

func (c *Server) Start() {
	c.engine.Run(c.port)
}
