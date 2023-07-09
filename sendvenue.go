package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendVenue struct {
	parent              *TelegramBot
	chatId              interface{}
	latitude            float64
	longitude           float64
	title               string
	address             string
	foursquareId        string
	foursquareType      string
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendVenue) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Latitude            float64     `json:"latitude"`
		Longitude           float64     `json:"longitude"`
		Title               string      `json:"title"`
		Address             string      `json:"address"`
		FoursquareId        string      `json:"foursquare_id"`
		FoursquareType      string      `json:"foursquare_type"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sv.chatId,
		Latitude:            sv.latitude,
		Longitude:           sv.longitude,
		Title:               sv.title,
		Address:             sv.address,
		FoursquareId:        sv.foursquareId,
		FoursquareType:      sv.foursquareType,
		DisableNotification: sv.disableNotification,
		ReplyToMessageId:    sv.replyToMessageId,
		ReplyMarkup:         sv.replyMarkup,
	})
}

func (sv *sendVenue) response() interface{} {
	return &structs.Message{}
}

func (sv *sendVenue) method() string {
	return "POST"
}

func (sv *sendVenue) endpoint() string {
	return "sendVenue"
}

func (sv *sendVenue) SetChatId(chatId int64) *sendVenue {
	sv.chatId = chatId
	return sv
}

func (sv *sendVenue) SetChatUsername(username string) *sendVenue {
	sv.chatId = username
	return sv
}

func (sv *sendVenue) SetLatitude(lat float64) *sendVenue {
	sv.latitude = lat
	return sv
}

func (sv *sendVenue) SetLongitude(lng float64) *sendVenue {
	sv.longitude = lng
	return sv
}

func (sv *sendVenue) SetTitle(title string) *sendVenue {
	sv.title = title
	return sv
}

func (sv *sendVenue) SetAddress(address string) *sendVenue {
	sv.address = address
	return sv
}

func (sv *sendVenue) SetFoursquareId(foursquareId string) *sendVenue {
	sv.foursquareId = foursquareId
	return sv
}

func (sv *sendVenue) SetFoursquareType(foursquareType string) *sendVenue {
	sv.foursquareType = foursquareType
	return sv
}

func (sv *sendVenue) SetDisableNotification() *sendVenue {
	sv.disableNotification = true
	return sv
}

func (sv *sendVenue) SetEnableNotification() *sendVenue {
	sv.disableNotification = false
	return sv
}

func (sv *sendVenue) SetReplyToMessageId(messageId int64) *sendVenue {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendVenue) SetReplyMarkup(markup interface{}) *sendVenue {
	sv.replyMarkup = tools.ParseReplyMarkup(markup)

	return sv
}

/*func (sv *sendVenue) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
