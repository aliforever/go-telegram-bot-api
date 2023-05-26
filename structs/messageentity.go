package structs

type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int64  `json:"offset"`
	Length   int64  `json:"length"`
	Url      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}

func (e MessageEntity) IsBotCommand() bool {
	return e.Type == "bot_command"
}
