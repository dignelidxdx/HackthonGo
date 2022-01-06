package models

type Product struct {
	ID          int     `form:"id" json:"id"`
	Description string  `form:"description" json:"description" validate:"required,description"`
	Price       float64 `form:"price" json:"price" validate:"required,price"`
}
