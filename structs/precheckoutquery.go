package structs

type PreCheckoutQuery struct {
	Id               string    `json:"id"`
	From             User      `json:"from,omitempty"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info,omitempty"`
}
