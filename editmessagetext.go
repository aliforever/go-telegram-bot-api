package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type editMessageText struct {
	parent                *TelegramBot
	chatId                interface{}
	messageId             int64
	inlineMessageId       string
	text                  string
	parseMode             string
	disableWebPagePreview bool
	replyMarkup           interface{}
}

func (m *editMessageText) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                interface{} `json:"chat_id"`
		MessageId             int64       `json:"message_id"`
		InlineMessageId       string      `json:"inline_message_id"`
		Text                  string      `json:"text,omitempty"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:                m.chatId,
		MessageId:             m.messageId,
		InlineMessageId:       m.inlineMessageId,
		Text:                  m.text,
		ParseMode:             m.parseMode,
		DisableWebPagePreview: m.disableWebPagePreview,
		ReplyMarkup:           m.replyMarkup,
	})
}

func (m *editMessageText) response() interface{} {
	return &structs.Message{}
}

func (m *editMessageText) method() string {
	return "POST"
}

func (m *editMessageText) endpoint() string {
	return "editMessageText"
}

/*type Text string

func (t Text) AddWithNewLine(text string) Text {
	t += "\n" + Text(text)
	return t
}

func (t Text) String() string {
	return string(t)
}*/

func (m *editMessageText) SetChatId(chatId int64) *editMessageText {
	m.chatId = chatId
	return m
}

func (m *editMessageText) SetChatUsername(username string) *editMessageText {
	m.chatId = username
	return m
}

func (m *editMessageText) SetText(message string) *editMessageText {
	m.text = message
	return m
}

func (m *editMessageText) SetParseMode(parseMode string) *editMessageText {
	m.parseMode = parseMode
	return m
}

func (m *editMessageText) SetParseModeHTML() *editMessageText {
	m.parseMode = "HTML"
	return m
}

func (m *editMessageText) SetParseModeMarkdown() *editMessageText {
	m.parseMode = "Markdown"
	return m
}

func (m *editMessageText) SetDisableWebPagePreview() *editMessageText {
	m.disableWebPagePreview = true
	return m
}

func (m *editMessageText) SetEnableWebPagePreview() *editMessageText {
	m.disableWebPagePreview = false
	return m
}

func (m *editMessageText) SetMessageId(messageId int64) *editMessageText {
	m.messageId = messageId
	return m
}

func (m *editMessageText) SetInlineMessageId(inlineMessageId string) *editMessageText {
	m.inlineMessageId = inlineMessageId
	return m
}

func (m *editMessageText) SetReplyMarkup(markup interface{}) *editMessageText {
	m.replyMarkup = markup
	return m
}
