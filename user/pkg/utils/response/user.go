package response

import (
	"time"

	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/models"
)

type Profile struct {
	ID             uint
	UserName       string
	FirstName      string
	LastName       string
	Age            uint
	Email          string
	Phone          string
	DefaultAddress Address
	OrderHistory   []OrderHistory
}

type UserRespStrcut struct {
	ID          uint
	FirstName   string
	LastName    string
	Age         uint
	Email       string
	UserName    string
	Phone       string
	BlockStatus bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// home page response
type ResUserHome struct {
	User UserRespStrcut
}

// cart item reponse
type CartItemResp struct {
	ProductItemID uint
	Name          string
	Price         uint
	DiscountPrice uint
	Quantity      uint
	QtyLeft       uint
	StockStatus   bool
	SubTotal      uint
}

type CartResp struct {
	CartItems         []CartItemResp
	TotalProductItems uint
	TotalQty          uint
	TotalPrice        float64
	DiscountAmount    float64
	AppliedCouponID   uint
	AppliedCouponCode string
	CouponDiscount    float64
	FinalPrice        uint
	DefaultShipping   Address
}
type Address struct {
	ID           uint
	UserID       uint
	House        string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	ZipCode      string
	Country      string
	IsDefault    bool
}

type CheckoutOrder struct {
	UserID         uint
	CartItemResp   []CartItemResp
	TotalQty       uint
	TotalPrice     uint
	Discount       uint
	DefaultAddress models.Address
}

type UserContact struct {
	Email string
	Phone string
}
type Wishlist struct {
	ProductItemId uint
	ProductName   string
	Color         string
	Storage       string
	Price         uint
	Quantity      uint
	Image         string
}
