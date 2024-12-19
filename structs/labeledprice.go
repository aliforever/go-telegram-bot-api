package structs

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}
