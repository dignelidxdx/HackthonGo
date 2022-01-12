package web

type ResultTotalQuantity struct {
	Quantity    float64 `json:"quantity"`
	Total       float64 `json:"total"`
	Description string  `json:"description"`
}
