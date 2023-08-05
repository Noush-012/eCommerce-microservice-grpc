package request

import "time"

type Users struct {
	ID          uint      `json:"id"`
	UserName    string    `json:"user_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         uint      `json:"age"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Password    string    `json:"password"`
	BlockStatus bool      `json:"block_status" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AddressPatchReq struct {
	ID           uint      `json:"address_id"`
	UserID       uint      `json:"-"`
	House        string    `json:"house"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
	UpdatedAt    time.Time `json:"-"`
}
type Address struct {
	ID           uint      `json:"address_id"`
	UserID       uint      `json:"-"`
	House        string    `json:"house"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type AddToCartReq struct {
	UserID         uint    `json:"user_id"`
	ProductItemID  uint    `json:"product_item_id" binding:"required"`
	Quantity       uint    `json:"quantity" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}

type UpdateCartReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}

type DeleteCartItemReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
}

type AddToWishlist struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}
