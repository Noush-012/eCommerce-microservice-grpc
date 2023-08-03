package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	client "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/client/interface"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/models"
	interfaces "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/repository/interface"
	service "github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/usecase/interface"
	"github.com/Noush-012/project-ecommerce-microservice/user-auth-service/pkg/verify"
	"github.com/jinzhu/copier"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	AuthRepository interfaces.AuthRepository
	userClient     client.UserClient
}

func NewAuthUseCase(repo interfaces.AuthRepository, userClient client.UserClient) service.AuthService {
	return &AuthUseCase{AuthRepository: repo,
		userClient: userClient,
	}
}

func (u *AuthUseCase) SignUp(ctx context.Context, user models.User) error {

	// Check if user already exist
	var data models.UserDataRequest
	copier.Copy(&data, user)
	DBUser, err := u.userClient.FindUser(data)

	if err != nil {
		return err
	}

	if DBUser.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.AuthRepository.SaveUser(ctx, user)
		if err != nil {
			return errors.New("failed to save user")
		}

	} else {
		return fmt.Errorf("%v user already exists", DBUser.UserName)
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, user models.Users) (models.Users, error) {
	// Find user in db
	var data models.UserDataRequest
	copier.Copy(&data, user)
	DBUser, err := u.userClient.FindUser(data)

	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	// Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("user blocked by admin")
	}

	if _, err := verify.TwilioSendOTP("+91" + DBUser.Phone); err != nil {
		// response := response.ErrorResponse(500, "failed to send otp", err.Error(), nil)
		// c.JSON(http.StatusInternalServerError, response)
		return user, fmt.Errorf("failed to send otp %v", err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		return user, errors.New("password incorrect")
	}
	var dbUser models.Users
	copier.Copy(&dbUser, DBUser)

	return dbUser, nil

}
func (u *AuthUseCase) OTPLogin(ctx *context.Context, body request.OTPVerify) (models.Users, error) {
	// Find user in db
	var data models.UserDataRequest
	copier.Copy(&data, body)
	DBUser, err := u.userClient.FindUser(data)

	fmt.Println("== == == == ", DBUser)
	if err != nil {
		return models.Users{}, err
	} else if DBUser.ID == 0 {
		return models.Users{}, errors.New("user not exist")
	}

	var dbUser models.Users
	copier.Copy(&dbUser, DBUser)

	// Verify otp
	err = verify.TwilioVerifyOTP("+91"+DBUser.Phone, body.OTP)
	if err != nil {

		return models.Users{}, errors.New(err.Error())
	}
	return dbUser, nil
}
