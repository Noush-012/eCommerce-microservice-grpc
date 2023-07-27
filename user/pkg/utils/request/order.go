package request

type ReturnRequest struct {
	UserID  uint
	OrderID uint
	Reason  string
}
type CancelOrder struct {
	UserID  uint
	OrderID uint
}
