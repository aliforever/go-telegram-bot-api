package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type editMessageMedia struct {
	parent          *TelegramBot
	chatId          interface{}
	messageId       int64
	inlineMessageId string
	media           inputMedia
	replyMarkup     interface{}
	file
	files []fileInfo
}

func (m *editMessageMedia) marshalJSON() ([]byte, error) {
	nM := struct {
		ChatId          interface{} `json:"chat_id"`
		MessageId       int64       `json:"message_id,omitempty"`
		InlineMessageId string      `json:"inline_message_id,omitempty"`
		Media           string      `json:"media,omitempty"`
		ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:          m.chatId,
		MessageId:       m.messageId,
		InlineMessageId: m.inlineMessageId,
		ReplyMarkup:     m.replyMarkup,
	}
	if m.media != nil {
		m.SetMedia(m.media)
		mj, _ := m.media.marshalJSON()
		nM.Media = string(mj)
	}
	return json.Marshal(nM)
}

func (m *editMessageMedia) response() interface{} {
	return &structs.Message{}
}

func (m *editMessageMedia) method() string {
	return "POST"
}

func (m *editMessageMedia) endpoint() string {
	return "editMessageMedia"
}
func (m *editMessageMedia) medias() []fileInfo {
	if m.media != nil {
		m.files = m.media.medias()
	}
	return m.files
}

/*type Caption string

func (t Caption) AddWithNewLine(caption string) Caption {
	t += "\n" + Caption(caption)
	return t
}

func (t Caption) String() string {
	return string(t)
}*/

func (m *editMessageMedia) InputPhoto() *inputMediaPhoto {
	photo := &inputMediaPhoto{}
	m.media = photo
	return photo
}

func (m *editMessageMedia) InputVideo() *inputMediaVideo {
	video := &inputMediaVideo{}
	m.media = video
	return video
}

func (m *editMessageMedia) SetChatId(chatId int64) *editMessageMedia {
	m.chatId = chatId
	return m
}

func (m *editMessageMedia) SetChatUsername(username string) *editMessageMedia {
	m.chatId = username
	return m
}

func (m *editMessageMedia) SetMedia(media inputMedia) *editMessageMedia {
	m.media = media
	m.files = media.medias()
	return m
}

func (m *editMessageMedia) SetMessageId(messageId int64) *editMessageMedia {
	m.messageId = messageId
	return m
}

func (m *editMessageMedia) SetInlineMessageId(inlineMessageId string) *editMessageMedia {
	m.inlineMessageId = inlineMessageId
	return m
}

func (m *editMessageMedia) SetReplyMarkup(markup interface{}) *editMessageMedia {
	if markup == nil {
		return m
	}

	m.replyMarkup = markup
	return m
}
