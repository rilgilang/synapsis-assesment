package request_model

type AddToCart struct {
	ProductId string `json:"product_id"`
	Total     int    `json:"total"`
}

type DeleteFromCart struct {
	Id string `json:"id"`
}
