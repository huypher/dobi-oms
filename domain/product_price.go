package domain

import "time"

type ProductPrice struct {
	ProductID          int       `json:"product_id"`
	OriginalPrice      float64   `json:"original_price"`
	DiscountPercentage float64   `json:"discount_percentage"`
	DiscountAmount     float64   `json:"discount_amount"`
	Expression         string    `json:"expression"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ProductPriceUsecase interface {
}

type ProductPriceRepository interface {
}
