package domain

type Employee struct {
	ID        int    `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"Maria"`
	LastName  string `json:"last_name" example:"Morales"`
	Condition string `json:"condition" example:"condition"`
}
