package domain

import "time"

type Customer struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Ward        string    `json:"ward"`
	District    string    `json:"district"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Rank        string    `json:"rank"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerUsecase interface {
}

type CustomerRepository interface {
}
