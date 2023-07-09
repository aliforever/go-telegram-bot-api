package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendLocation struct {
	parent              *TelegramBot
	chatId              interface{}
	latitude            float64
	longitude           float64
	livePeriod          int64
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sl *sendLocation) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Latitude            float64     `json:"latitude"`
		Longitude           float64     `json:"longitude"`
		LivePeriod          int64       `json:"live_period"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sl.chatId,
		Latitude:            sl.latitude,
		Longitude:           sl.longitude,
		LivePeriod:          sl.livePeriod,
		DisableNotification: sl.disableNotification,
		ReplyToMessageId:    sl.replyToMessageId,
		ReplyMarkup:         sl.replyMarkup,
	})
}

func (sl *sendLocation) response() interface{} {
	return &structs.Message{}
}

func (sl *sendLocation) method() string {
	return "POST"
}

func (sl *sendLocation) endpoint() string {
	return "sendLocation"
}

func (sl *sendLocation) SetChatId(chatId int64) *sendLocation {
	sl.chatId = chatId
	return sl
}

func (sl *sendLocation) SetChatUsername(username string) *sendLocation {
	sl.chatId = username
	return sl
}

func (sl *sendLocation) SetLatitude(lat float64) *sendLocation {
	sl.latitude = lat
	return sl
}

func (sl *sendLocation) SetLongitude(lng float64) *sendLocation {
	sl.longitude = lng
	return sl
}

func (sl *sendLocation) SetLivePeriod(lP int64) *sendLocation {
	sl.livePeriod = lP
	return sl
}

func (sl *sendLocation) SetDisableNotification() *sendLocation {
	sl.disableNotification = true
	return sl
}

func (sl *sendLocation) SetEnableNotification() *sendLocation {
	sl.disableNotification = false
	return sl
}

func (sl *sendLocation) SetReplyToMessageId(messageId int64) *sendLocation {
	sl.replyToMessageId = messageId
	return sl
}

func (sl *sendLocation) SetReplyMarkup(markup interface{}) *sendLocation {
	sl.replyMarkup = tools.ParseReplyMarkup(markup)

	return sl
}

/*func (sl *sendLocation) Send() (message *structs.Message, err error) {
	message, err = sl.parent.Send(sd)
	return
}
*/
