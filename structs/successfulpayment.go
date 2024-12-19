package structs

type SuccessfulPayment struct {
	Currency                   string    `json:"currency"`
	TotalAmount                int64     `json:"total_amount"`
	InvoicePayload             string    `json:"invoice_payload"`
	SubscriptionExpirationDate int64     `json:"subscription_expiration_date"`
	IsRecurring                bool      `json:"is_recurring"`
	IsFirstRecurring           bool      `json:"is_first_recurring"`
	ShippingOptionId           string    `json:"shipping_option_id"`
	OrderInfo                  OrderInfo `json:"order_info"`
	TelegramPaymentChargeId    string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId    string    `json:"provider_payment_charge_id"`
}

func (sp *SuccessfulPayment) IsStars() bool {
	return sp.Currency == "XTR"
}
