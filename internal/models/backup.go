package models

type BackUp struct {
	ID            int    `form:"id" json:"id" csv:"id"`
	IsUpdatedData string `form:"is_update_data" json:"is_update_data" validate:"required,is_update_data"`
	Field         string `form:"field" json:"field" validate:"required,field"`
}
