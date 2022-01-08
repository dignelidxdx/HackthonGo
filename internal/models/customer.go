package models

type Customer struct {
	ID        int    `form:"id" json:"id" csv:"id"`
	LastName  string `form:"last_name" json:"last_name" validate:"required,last_name" csv:"last_name"`
	FirstName string `form:"first_name" json:"first_name" validate:"required,first_name" csv:"first_name"`
	Condition string `form:"condition" json:"condition" validate:"required,condition" csv:"condition"`
}
