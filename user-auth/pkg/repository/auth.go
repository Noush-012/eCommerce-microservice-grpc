package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
	repo "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type AuthDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) repo.AuthRepository {
	return &AuthDatabase{DB: DB}
}

func (i *AuthDatabase) SaveUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	dbWithContext := i.DB.WithContext(ctx)
	result := dbWithContext.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age,
		user.Email, user.Phone, user.Password, createdAt)
	if result.Error != nil {
		return fmt.Errorf("failed to save user %s: %v", user.UserName, result.Error)
	}
	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
