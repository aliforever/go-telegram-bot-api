package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type editMessageLiveLocation struct {
	parent          *TelegramBot
	chatId          interface{}
	messageId       int64
	inlineMessageId string
	latitude        float64
	longitude       float64
	replyMarkup     interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (emll *editMessageLiveLocation) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId          interface{} `json:"chat_id,omitempty"`
		MessageId       int64       `json:"message_id,omitempty"`
		InlineMessageId string      `json:"inline_message_id,omitempty"`
		Latitude        float64     `json:"latitude"`
		Longitude       float64     `json:"longitude"`
		ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:          emll.chatId,
		MessageId:       emll.messageId,
		InlineMessageId: emll.inlineMessageId,
		Latitude:        emll.latitude,
		Longitude:       emll.longitude,
		ReplyMarkup:     emll.replyMarkup,
	})
}

func (emll *editMessageLiveLocation) response() interface{} {
	return &structs.Message{}
}

func (emll *editMessageLiveLocation) method() string {
	return "POST"
}

func (emll *editMessageLiveLocation) endpoint() string {
	return "editMessageLiveLocation"
}

func (emll *editMessageLiveLocation) SetChatId(chatId int64) *editMessageLiveLocation {
	emll.chatId = chatId
	return emll
}

func (emll *editMessageLiveLocation) SetChatUsername(username string) *editMessageLiveLocation {
	emll.chatId = username
	return emll
}

func (emll *editMessageLiveLocation) SetLatitude(lat float64) *editMessageLiveLocation {
	emll.latitude = lat
	return emll
}

func (emll *editMessageLiveLocation) SetLongitude(lng float64) *editMessageLiveLocation {
	emll.longitude = lng
	return emll
}

func (emll *editMessageLiveLocation) SetMessageId(messageId int64) *editMessageLiveLocation {
	emll.messageId = messageId
	return emll
}

func (emll *editMessageLiveLocation) SetInlineMessageId(inlineMessageId string) *editMessageLiveLocation {
	emll.inlineMessageId = inlineMessageId
	return emll
}

func (emll *editMessageLiveLocation) SetReplyMarkup(markup interface{}) *editMessageLiveLocation {
	emll.replyMarkup = &markup
	return emll
}

/*func (emll *editMessageLiveLocation) Send() (message *structs.Message, err error) {
	message, err = emll.parent.Send(sd)
	return
}
*/
