package domain

import "time"

type Order struct {
	ID         int       `json:"id"`
	Code       string    `json:"code"`
	CustomerID int       `json:"customer_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderUsecase interface {
}

type OrderRepository interface {
}
