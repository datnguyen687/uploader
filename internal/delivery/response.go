package delivery

import "ecommerce/internal/models"

type BasicResponse struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message"`
	Error   string     `json:"error"`
}

type FilterProductsResponse struct {
	BasicResponse
	Data []models.Product `json:"data"`
}
