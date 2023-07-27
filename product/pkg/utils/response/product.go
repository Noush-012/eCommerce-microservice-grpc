package response

import "time"

type ResponseProduct struct {
	ID            uint
	Name          string
	Description   string
	Category_name string
	Price         uint
	DiscountPrice uint
	Image         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProductItemResp struct {
	ProductID      uint
	ProductItemID  uint
	StockAvailable uint
	ProductName    string
	Brand          string
	Description    string
	Color          string
	Storage        uint
	Price          uint
	OfferPrice     uint
	Images         []string
}

type Brand struct {
	ID           uint
	CategoryName string
}
