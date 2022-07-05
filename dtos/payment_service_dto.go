package dtos

type PaymentServiceDto struct {
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
	Code        string `json:"code"`
}
