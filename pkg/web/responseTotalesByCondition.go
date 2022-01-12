package web

type ResponseTotalesByCondition struct {
	Condition string  `json:"condition"`
	Total     float64 `json:"total"`
}
