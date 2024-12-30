package tests

var (
	ChatID   int64  = 5760288822
	BotToken string = "6789667899:AAGgzbYwp7Cu23DRMCgJIiICd1PNCiNWc3I"
	UserID   int64  = 0
)

type Tests struct {
	BotToken string
	ChatId   int64
	UserId   int64
}

func (t Tests) Defaults() Tests {
	return Tests{
		ChatId:   ChatID,
		BotToken: BotToken,
		UserId:   UserID,
	}
}
