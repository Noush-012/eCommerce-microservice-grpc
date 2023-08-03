package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
)

type AuthClient interface {
	SignUp(ctx context.Context, user models.Users) error
	Login(ctx context.Context, user models.Users) (models.Users, error)
	OTPLogin(ctx context.Context, body request.OTPVerify) (models.Users, error)
}
