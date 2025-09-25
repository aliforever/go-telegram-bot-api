package structs

// WebAppInfo describes a Web App.
// See https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}
