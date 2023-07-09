package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type editMessageCaption struct {
	parent          *TelegramBot
	chatId          interface{}
	messageId       int64
	inlineMessageId string
	caption         string
	parseMode       string
	replyMarkup     interface{}
}

func (m *editMessageCaption) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId          interface{} `json:"chat_id"`
		MessageId       int64       `json:"message_id"`
		InlineMessageId string      `json:"inline_message_id"`
		Caption         string      `json:"caption,omitempty"`
		ParseMode       string      `json:"parse_mode,omitempty"`
		ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:          m.chatId,
		MessageId:       m.messageId,
		InlineMessageId: m.inlineMessageId,
		Caption:         m.caption,
		ParseMode:       m.parseMode,
		ReplyMarkup:     m.replyMarkup,
	})
}

func (m *editMessageCaption) response() interface{} {
	return &structs.Message{}
}

func (m *editMessageCaption) method() string {
	return "POST"
}

func (m *editMessageCaption) endpoint() string {
	return "editMessageCaption"
}

/*type Caption string

func (t Caption) AddWithNewLine(caption string) Caption {
	t += "\n" + Caption(caption)
	return t
}

func (t Caption) String() string {
	return string(t)
}*/

func (m *editMessageCaption) SetChatId(chatId int64) *editMessageCaption {
	m.chatId = chatId
	return m
}

func (m *editMessageCaption) SetChatUsername(username string) *editMessageCaption {
	m.chatId = username
	return m
}

func (m *editMessageCaption) SetCaption(caption string) *editMessageCaption {
	m.caption = caption
	return m
}

func (m *editMessageCaption) SetParseMode(parseMode string) *editMessageCaption {
	m.parseMode = parseMode
	return m
}

func (m *editMessageCaption) SetParseModeHTML() *editMessageCaption {
	m.parseMode = "HTML"
	return m
}

func (m *editMessageCaption) SetParseModeMarkdown() *editMessageCaption {
	m.parseMode = "Markdown"
	return m
}

func (m *editMessageCaption) SetMessageId(messageId int64) *editMessageCaption {
	m.messageId = messageId
	return m
}

func (m *editMessageCaption) SetInlineMessageId(inlineMessageId string) *editMessageCaption {
	m.inlineMessageId = inlineMessageId
	return m
}

func (m *editMessageCaption) SetReplyMarkup(markup interface{}) *editMessageCaption {
	m.replyMarkup = tools.ParseReplyMarkup(markup)

	return m
}
