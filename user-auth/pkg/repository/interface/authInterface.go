package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
)

type AuthRepository interface {
	SaveUser(ctx context.Context, user models.User) error
}
