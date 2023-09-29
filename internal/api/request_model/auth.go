package request_model

// this struct is used for request body at endpoint login and register
type Register struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
