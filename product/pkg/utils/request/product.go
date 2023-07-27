package request

type ProductReq struct {
	Name        string
	Description string
	CategoryID  uint
	Price       uint
	Image       string
}

type UpdateProductReq struct {
	ID          uint
	Name        string
	Description string
	CategoryID  uint
	Price       uint
	Image       string
}
type UpdateProductItemReq struct {
	ID            uint
	ProductId     uint
	QtyInStock    uint
	SKU           string
	DiscountPrice string
	Price         uint
	Image         string
}

type DeleteProductReq struct {
	ID uint
}

type CategoryReq struct {
	ID           uint
	ParentID     uint
	CategoryName string
}

type ProductItemReq struct {
	ProductID      uint
	ProductItemId  uint
	QtyInStock     uint
	Price          uint
	DiscountPrice  uint
	SKU            string
	Configurations map[string]Variation
	Images         []string
}

type Variation struct {
	VariationOptionID uint
}
