package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int32
	UserName    string
	FirstName   string
	LastName    string
	Age         int32
	Email       string
	Phone       string
	Password    string
	BlockStatus bool
}

type Wallet struct {
	ID        uint
	UserID    uint
	Balance   float64
	Remark    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Address struct {
	ID           uint
	UserID       uint
	User         User
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

// cart model
type CartItem struct {
	ID            uint
	CartID        uint
	ProductItemID uint
	Quantity      uint
	StockStatus   bool
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Cart struct {
	ID     uint
	UserID uint
	Items  []CartItem
	Total  float64
}

type Wishlist struct {
	ID            uint
	UserID        uint
	ProductItemID uint
	Quantity      uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
