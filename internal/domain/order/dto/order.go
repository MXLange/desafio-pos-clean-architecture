package dto

type Order struct {
	ID        uint `json:"id"`
	ProductID uint `json:"productId"`
	Quantity  uint `json:"quantity"`
}

type OrderCreateResponse struct {
	ID        uint `json:"id"`
	ProductID uint `json:"productId"`
	Quantity  uint `json:"quantity"`
}

type OrderCreateRequest struct {
	ProductID uint `json:"productId" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
}
