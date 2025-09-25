package structs

// WebApp represents a Web App.
// See https://core.telegram.org/bots/api#webapp
type WebApp struct {
	URL string `json:"url"`
}
