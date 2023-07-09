package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/structs"
)

type editMessageReplyMarkup struct {
	parent          *TelegramBot
	chatId          interface{}
	messageId       int64
	inlineMessageId string
	replyMarkup     *structs.InlineKeyboardMarkup
}

func (m *editMessageReplyMarkup) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId          interface{}                   `json:"chat_id"`
		MessageId       int64                         `json:"message_id"`
		InlineMessageId string                        `json:"inline_message_id"`
		ReplyMarkup     *structs.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	}{
		ChatId:          m.chatId,
		MessageId:       m.messageId,
		InlineMessageId: m.inlineMessageId,
		ReplyMarkup:     m.replyMarkup,
	})
}

func (m *editMessageReplyMarkup) response() interface{} {
	return &structs.Message{}
}

func (m *editMessageReplyMarkup) method() string {
	return "POST"
}

func (m *editMessageReplyMarkup) endpoint() string {
	return "editMessageReplyMarkup"
}

func (m *editMessageReplyMarkup) SetChatId(chatId int64) *editMessageReplyMarkup {
	m.chatId = chatId
	return m
}

func (m *editMessageReplyMarkup) SetChatUsername(username string) *editMessageReplyMarkup {
	m.chatId = username
	return m
}

func (m *editMessageReplyMarkup) SetMessageId(messageId int64) *editMessageReplyMarkup {
	m.messageId = messageId
	return m
}

func (m *editMessageReplyMarkup) SetInlineMessageId(inlineMessageId string) *editMessageReplyMarkup {
	m.inlineMessageId = inlineMessageId
	return m
}

func (m *editMessageReplyMarkup) SetReplyMarkup(markup *structs.InlineKeyboardMarkup) *editMessageReplyMarkup {
	m.replyMarkup = markup

	return m
}
