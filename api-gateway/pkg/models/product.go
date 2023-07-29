package models

import "time"

// Category struct
type Category struct {
	ID           uint
	ParentID     *uint
	Parent       *Category
	CategoryName string
	Products     []*Product
	Children     []*Category
}

// Product struct
type Product struct {
	ID            uint
	Name          string
	Description   string
	CategoryID    uint
	Category      *Category
	Price         uint
	DiscountPrice uint
	Image         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Items         []*ProductItem
}

// ProductItem struct
type ProductItem struct {
	ID             uint
	ProductID      uint
	QtyInStock     uint
	StockStatus    bool
	Price          uint
	SKU            string
	DiscountPrice  uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Configurations []*ProductConfig
	Images         []*ProductImage
}

// Variation struct COLOR or STORAGE
type Variation struct {
	ID            uint
	CategoryID    uint
	Category      *Category
	VariationName string
	Options       []*VariationOption
}

// VariationOption struct eg: 128 , 512, red, black
type VariationOption struct {
	ID             uint
	VariationID    uint
	Variation      *Variation
	OptionValue    string
	Configurations []*ProductConfig
}

// ProductConfig struct
type ProductConfig struct {
	ID                uint
	ProductItemID     uint
	VariationOptionID uint
	ProductItem       *ProductItem
	VariationOption   *VariationOption
}

type ProductImage struct {
	ID            uint
	ProductItemID uint
	ProductItem   ProductItem
	Image         string
}
