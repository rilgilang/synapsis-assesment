package request_model

type GetProducts struct {
	Category string `json:"category" db:"product_categories.category_name"`
}
