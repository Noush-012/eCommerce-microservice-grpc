tidy:
	go mod tidy
proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

wire:
	cd pkg/di && wire

server:
	go run cmd/main.go

docker-build:
	docker build -t noush-012/project-ecommerce-microservice-user-svc .