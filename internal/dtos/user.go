package dtos

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `jsom:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Mobile   string    `json:"mobile"`
	Role     string    `json:"role"`
}

type LoginRes struct {
	ID       uuid.UUID `jsom:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Mobile   string    `json:"mobile"`
	Role     string    `json:"role"`
	Token    string    `json:"token"`
}

type ProductQuantity struct {
	ProductID            uuid.UUID
	Name                 string
	TotalQuantityOrdered int
}
