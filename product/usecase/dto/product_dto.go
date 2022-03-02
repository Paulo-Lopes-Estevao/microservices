package dto

type ProductDtoInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price,string"`
}

type ProductDtoRequestInput map[string]interface{}

type ProductDtoOutput struct {
	Data         ProductDtoInput `json:"data"`
	ErrorMessage string          `json:"error_message"`
}
