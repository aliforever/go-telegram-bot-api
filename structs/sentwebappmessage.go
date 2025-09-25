package structs

// SentWebAppMessage describes an inline message sent by a Web App on behalf of a user.
// See https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id,omitempty"`
}
