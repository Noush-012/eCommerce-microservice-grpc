package request

type LoginData struct {
	UserName string `json:"user_name" binding:"omitempty,min=3,max=15"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"required,min=5,max=30"`
}

type SignupUserData struct {
	UserName        string `json:"user_name"  binding:"required,min=3,max=15"`
	FirstName       string `json:"first_name"  binding:"required,min=2,max=50"`
	LastName        string `json:"last_name"  binding:"required,min=1,max=50"`
	Age             uint   `json:"age"  binding:"required,numeric"`
	Email           string `json:"email" binding:"required,email"`
	Phone           string `json:"phone" binding:"required,min=10,max=10"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type OTPVerify struct {
	OTP    string `json:"otp" binding:"required,min=4,max=8"`
	UserID uint   `json:"user_id" binding:"required,numeric"`
}

// Page number and count of products
type ReqPagination struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}

type UserID struct {
	UserID uint `json:"user_id" binding:"required,numeric"`
}
