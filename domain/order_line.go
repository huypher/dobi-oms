package domain

type OrderLine struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Qty       int     `json:"qty"`
	OrderID   int     `json:"order_id"`
	Price     float64 `json:"price"`
}

type OrderLineUsecase interface {
}

type OrderLineRepository interface {
}
