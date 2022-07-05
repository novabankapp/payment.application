package dtos

type AccountDto struct {
	Identifier     string            `json:"identifier"`
	PaymentService PaymentServiceDto `json:"payment_service"`
}
