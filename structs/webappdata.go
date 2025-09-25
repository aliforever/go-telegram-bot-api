package structs

// WebAppData describes data sent from a Web App to the bot.
// See https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}
