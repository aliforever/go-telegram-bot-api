package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendMessage struct {
	parent                *TelegramBot
	chatId                interface{}
	text                  string
	parseMode             string
	disableWebPagePreview bool
	disableNotification   bool
	replyToMessageId      int64
	replyMarkup           interface{}
}

func (m *sendMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                interface{} `json:"chat_id"`
		Text                  string      `json:"text,omitempty"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		DisableNotification   bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:                m.chatId,
		Text:                  m.text,
		ParseMode:             m.parseMode,
		DisableWebPagePreview: m.disableWebPagePreview,
		DisableNotification:   m.disableNotification,
		ReplyToMessageId:      m.replyToMessageId,
		ReplyMarkup:           m.replyMarkup,
	})
}

func (m *sendMessage) response() interface{} {
	return &structs.Message{}
}

func (m *sendMessage) method() string {
	return "POST"
}

func (m *sendMessage) endpoint() string {
	return "sendMessage"
}

/*type Text string

func (t Text) AddWithNewLine(text string) Text {
	t += "\n" + Text(text)
	return t
}

func (t Text) String() string {
	return string(t)
}*/

func (m *sendMessage) SetChatId(chatId int64) *sendMessage {
	m.chatId = chatId
	return m
}

func (m *sendMessage) SetChatUsername(username string) *sendMessage {
	m.chatId = username
	return m
}

func (m *sendMessage) SetText(message string) *sendMessage {
	m.text = message
	return m
}

func (m *sendMessage) SetParseMode(parseMode string) *sendMessage {
	m.parseMode = parseMode
	return m
}

func (m *sendMessage) SetParseModeHTML() *sendMessage {
	m.parseMode = "HTML"
	return m
}

func (m *sendMessage) SetParseModeMarkdown() *sendMessage {
	m.parseMode = "Markdown"
	return m
}

func (m *sendMessage) SetDisableWebPagePreview() *sendMessage {
	m.disableWebPagePreview = true
	return m
}

func (m *sendMessage) SetEnableWebPagePreview() *sendMessage {
	m.disableWebPagePreview = false
	return m
}

func (m *sendMessage) SetDisableNotification() *sendMessage {
	m.disableNotification = true
	return m
}

func (m *sendMessage) SetEnableNotification() *sendMessage {
	m.disableNotification = false
	return m
}

func (m *sendMessage) SetReplyToMessageId(messageId int64) *sendMessage {
	m.replyToMessageId = messageId
	return m
}

func (m *sendMessage) SetReplyMarkup(markup interface{}) *sendMessage {
	if markup == nil {
		return m
	}

	m.replyMarkup = markup
	return m
}

/*func (m *Send) SetReplyMarkupFromSlicesOfStrings(rows [][]string, resize, oneTime, selective bool) *Send {
	b := structs.ReplyKeyboardMarkup{}.FromSlicesOfStrings(rows)
	b.ResizeKeyboard = resize
	b.OneTimeKeyboard = oneTime
	b.Selective = selective
	var keyboard interface{} = b
	m.replyMarkup = keyboard
	return m
}
*/
/*func (m *sendMessage) Send() (message *structs.Message, err error) {
	message, err = m.parent.Send(m)
	return
}
*/
