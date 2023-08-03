package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
)

type AuthService interface {
	SignUp(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.Users) (models.Users, error)
	OTPLogin(ctx *context.Context, body request.OTPVerify) (models.Users, error)
}
