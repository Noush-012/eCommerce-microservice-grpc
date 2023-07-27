package api

import (
	"log"
	"net"

	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/config"
	"github.com/Noush-012/project-ecommerce-microservice/product-service/pkg/pb"

	"google.golang.org/grpc"
)

type ServeServer struct {
	gs     *grpc.Server
	listen net.Listener
}

func NewGRPCServer(cfg *config.Config, server pb.ProductServiceServer) (*ServeServer, error) {

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, server)
	listen, err := net.Listen("tcp", cfg.PRODUCT_SVC_PORT)
	if err != nil {
		return nil, err
	}

	return &ServeServer{
		gs:     grpcServer,
		listen: listen,
	}, nil
}

func (c *ServeServer) Run(cfg *config.Config) error {
	log.Printf("Auth services listening on port : %v ...", cfg.PRODUCT_SVC_PORT)
	return c.gs.Serve(c.listen)
}
