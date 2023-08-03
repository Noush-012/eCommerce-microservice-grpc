package handler

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/api/handler/interfaces"
	client "github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/client/interfaces"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type authHandler struct {
	client client.AuthClient
}

func NewAuthHandler(client client.AuthClient) interfaces.AuthHandler {
	return &authHandler{client: client}
}

func (u *authHandler) LoginPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// UserSignUp godoc
// @summary api for register user
// @security ApiKeyAuth
// @id UserSignUp
// @tags User Signup
// @Param input body request.SignupUserData{} true "Input Fields"
// @Router /signup [post]
// @Success 200 {object} response.Response{} "Account created successfuly"
// @Failure 400  {object} response.Response{} ""invalid entry"
// @Failure 400 {object} response.Response{}  "user already exist"
func (u *authHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Invalid input", "", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//
	var user models.Users
	// var user domain.Users
	if err := copier.Copy(&user, body); err != nil {
		fmt.Println("Copy failed")
	}

	// Check the user already exist in DB and save user if not
	if err := u.client.SignUp(c, user); err != nil {
		response := response.ErrorResponse(400, "User already exist", "", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// success response
	response := response.SuccessResponse(200, "Account created successfuly", nil)
	c.JSON(http.StatusOK, response)

}

// func (u *authHandler) LoginSubmit(c *gin.Context) {
// 	var body request.LoginData
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	if body.Email == "" && body.Password == "" && body.UserName == "" {
// 		_ = errors.New("please enter user_name and password")
// 		response := "Field should not be empty"
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	// Copying
// 	var user models.Users
// 	copier.Copy(&user, body)

// 	dbUser, err := u.client.Login(c, user)
// 	if err != nil {
// 		response := response.ErrorResponse(500, "Something went wrong !", err.Error(), nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := gin.H{"Successfuly send OTP to registered mobile number": dbUser.ID}
// 	c.JSON(http.StatusOK, response)
// }

// func (u *authHandler) UserOTPVerify(c *gin.Context) {

// 	var body request.OTPVerify
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	var user = models.Users{
// 		ID: body.UserID,
// 	}

// 	usr, err := u.client.OTPLogin(c, user)
// 	if err != nil {
// 		response := response.ErrorResponse(500, "user not registered", err.Error(), user)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	// Verify otp
// 	err = verify.TwilioVerifyOTP("+91"+usr.Phone, body.OTP)
// 	if err != nil {
// 		response := gin.H{"error": err.Error()}
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	// setup JWT
// 	ok := auth.JwtCookieSetup(c, "user-auth", usr.ID)
// 	if !ok {
// 		response := response.ErrorResponse(500, "failed to login", "", nil)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return

// 	}
// 	response := response.SuccessResponse(200, "Successfuly logged in!", nil)
// 	c.JSON(http.StatusOK, response)
// }
