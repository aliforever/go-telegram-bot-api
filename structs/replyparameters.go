package structs

type ReplyParameters struct {
	MessageID                int64           `json:"message_id"`
	ChatID                   any             `json:"chat_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Quote                    string          `json:"quote"`
	QuoteParseMode           string          `json:"quote_parse_mode"`
	QuoteEntities            []MessageEntity `json:"quote_entities"`
	QuotePosition            int64           `json:"quote_position"`
}
