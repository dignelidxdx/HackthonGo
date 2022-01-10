package models

type Invoice struct {
	ID       int      `form:"id" json:"id"`
	Datetime string   `form:"datetime" json:"datetime" validate:"required,datetime"`
	Customer Customer `form:"customer" json:"customer" validate:"required,customer"`
	Total    float64  `form:"condition" json:"condition" validate:"required,condition"`
}
