package models

type Sale struct {
	ID       int     `form:"id" json:"id"`
	Invoice  Invoice `form:"invoice" json:"invoice" validate:"required,invoice"`
	Product  Product `form:"product" json:"product" validate:"required,product"`
	Quantity float64 `form:"quantity" json:"quantity" validate:"required,quantity"`
}
