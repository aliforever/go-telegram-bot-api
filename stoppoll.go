package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type stopPoll struct {
	parent      *TelegramBot
	chatId      interface{}
	messageId   int64
	replyMarkup *structs.InlineKeyboardMarkup
}

func (m *stopPoll) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId      interface{}                   `json:"chat_id"`
		MessageId   int64                         `json:"message_id"`
		ReplyMarkup *structs.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}{
		ChatId:      m.chatId,
		MessageId:   m.messageId,
		ReplyMarkup: m.replyMarkup,
	})
}

func (m *stopPoll) response() interface{} {
	var resp interface{}
	return &resp
}

func (m *stopPoll) method() string {
	return "POST"
}

func (m *stopPoll) endpoint() string {
	return "stopPoll"
}

func (m *stopPoll) SetChatId(chatId int64) *stopPoll {
	m.chatId = chatId
	return m
}

func (m *stopPoll) SetChatUsername(username string) *stopPoll {
	m.chatId = username
	return m
}

func (m *stopPoll) SetMessageId(messageId int64) *stopPoll {
	m.messageId = messageId
	return m
}

func (m *stopPoll) SetReplyMarkup(markup *structs.InlineKeyboardMarkup) *stopPoll {
	if markup == nil {
		return m
	}

	m.replyMarkup = markup
	return m
}
