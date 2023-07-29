package models

import "time"

type Admin struct {
	ID        uint
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type SalesReport struct {
	OrderID        int
	UserID         int
	TotalAmount    float64
	CouponCode     string
	PaymentMethod  string
	OrderStatus    string
	DeliveryStatus string
	OrderDate      time.Time
}
