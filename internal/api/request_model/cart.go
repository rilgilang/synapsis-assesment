package request_model

type AddToCart struct {
	ProductId string `json:"product_id"`
	Total     int    `json:"total"`
}
