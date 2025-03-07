package dtos

import "github.com/google/uuid"

type Customer struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Mobile   string    `json:"mobile"`
}

type Order struct {
	Id              uuid.UUID
	CustomerId      uuid.UUID
	ProductId       uuid.UUID
	OrderId         string
	QuantityOrdered int
	TotalPrice      float64
}
