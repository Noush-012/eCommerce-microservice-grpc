package api

import (
	"log"
	"net"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/config"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServeServer struct {
	gs     *grpc.Server
	listen net.Listener
}

func NewGRPCServer(cfg *config.Config, server pb.AuthServiceServer) (*ServeServer, error) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	listen, err := net.Listen("tcp", cfg.AUTH_SVC_PORT)
	if err != nil {
		return nil, err
	}

	return &ServeServer{
		gs:     grpcServer,
		listen: listen,
	}, nil
}

func (c *ServeServer) Run(cfg *config.Config) error {
	log.Printf("Auth services listening on port : %v ...", cfg.AUTH_SVC_PORT)
	return c.gs.Serve(c.listen)
}
