package request

type LoginData struct {
	UserName string
	Email    string
	Password string
}

type SignupUserData struct {
	UserName        string
	FirstName       string
	LastName        string
	Age             uint
	Email           string
	Phone           string
	Password        string
	ConfirmPassword string
}

type OTPVerify struct {
	OTP    string
	UserID uint
}

// Page number and count of products
type ReqPagination struct {
	Count      uint
	PageNumber uint
}

type UserID struct {
	UserID uint
}
