package response

import (
	"time"
)

type ShopOrder struct {
	Id              uint
	OrderID         uint
	OrderDate       time.Time
	OrderTotal      float64
	Shipping_id     uint
	ShippingAddress Address
	OrderStatusID   uint
	OrderStatus     string
	PaymentMethod   string
	PaymentStatus   string
	TransactionID   string
}

type OrderHistory struct {
	ID             uint32
	OrderDate      time.Time
	OrderStatus    string
	DeliveryStatus string
	OrderTotal     float64
	PaymentMethod  string
	PaymentStatus  string
}
type Actions struct {
	Id   uint
	Name string
}

type ReturnRequests struct {
	ReturnID      uint
	UserID        uint
	OrderId       uint
	RequestedAt   time.Time
	OrderDate     time.Time
	DeliveredAt   time.Time
	PaymentMethod string
	PaymentStatus string
	Reason        string
	OrderTotal    uint
	IsApproved    bool
}
