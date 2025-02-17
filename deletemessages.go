package tgbotapi

import (
	"encoding/json"
)

type deleteMessages struct {
	parent     *TelegramBot
	chatId     interface{}
	messageIDs []int64
}

func (m *deleteMessages) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId     interface{} `json:"chat_id"`
		MessageIds []int64     `json:"message_ids"`
	}{
		ChatId:     m.chatId,
		MessageIds: m.messageIDs,
	})
}

func (m *deleteMessages) response() interface{} {
	var resp bool
	return &resp
}

func (m *deleteMessages) method() string {
	return "POST"
}

func (m *deleteMessages) endpoint() string {
	return "deleteMessages"
}

func (m *deleteMessages) SetChatId(chatId int64) *deleteMessages {
	m.chatId = chatId
	return m
}

func (m *deleteMessages) SetChatUsername(username string) *deleteMessages {
	m.chatId = username
	return m
}

func (m *deleteMessages) SetMessageIDs(messageIds []int64) *deleteMessages {
	m.messageIDs = messageIds
	return m
}

func (m *deleteMessages) AddMessageID(messageID int64) *deleteMessages {
	m.messageIDs = append(m.messageIDs, messageID)
	return m
}
