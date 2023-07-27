package request

import "time"

type Users struct {
	ID          uint
	UserName    string
	FirstName   string
	LastName    string
	Age         uint
	Email       string
	Phone       string
	Password    string
	BlockStatus bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AddressPatchReq struct {
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
	UpdatedAt    time.Time
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
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AddToCartReq struct {
	UserID         uint
	ProductItemID  uint
	Quantity       uint
	Price          float64
	Discount_price uint
}

type UpdateCartReq struct {
	UserID        uint
	ProductItemID uint
	Quantity      uint
}

type DeleteCartItemReq struct {
	UserID        uint
	ProductItemID uint
}

type AddToWishlist struct {
	UserID        uint
	ProductItemID uint
	Quantity      uint
}
