package interfaces

import (
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
)

type UserClient interface {
	FindUser(user models.UserDataRequest) (models.User, error)
}
