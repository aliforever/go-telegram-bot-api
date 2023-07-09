package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendContact struct {
	parent              *TelegramBot
	chatId              interface{}
	phoneNumber         string
	firstName           string
	lastName            string
	vcard               string
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendContact) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		PhoneNumber         string      `json:"phone_number"`
		FirstName           string      `json:"first_name"`
		LastName            string      `json:"last_name,omitempty"`
		Vcard               string      `json:"vcard,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sv.chatId,
		PhoneNumber:         sv.phoneNumber,
		FirstName:           sv.firstName,
		LastName:            sv.lastName,
		Vcard:               sv.vcard,
		DisableNotification: sv.disableNotification,
		ReplyToMessageId:    sv.replyToMessageId,
		ReplyMarkup:         sv.replyMarkup,
	})
}

func (sv *sendContact) response() interface{} {
	return &structs.Message{}
}

func (sv *sendContact) method() string {
	return "POST"
}

func (sv *sendContact) endpoint() string {
	return "sendContact"
}

func (sv *sendContact) SetChatId(chatId int64) *sendContact {
	sv.chatId = chatId
	return sv
}

func (sv *sendContact) SetChatUsername(username string) *sendContact {
	sv.chatId = username
	return sv
}

func (sv *sendContact) SetPhoneNumber(phone string) *sendContact {
	sv.phoneNumber = phone
	return sv
}

func (sv *sendContact) SetFirstName(firstName string) *sendContact {
	sv.firstName = firstName
	return sv
}

func (sv *sendContact) SetLastName(lastName string) *sendContact {
	sv.lastName = lastName
	return sv
}

func (sv *sendContact) SetVcard(vcard string) *sendContact {
	sv.vcard = vcard
	return sv
}

func (sv *sendContact) SetDisableNotification() *sendContact {
	sv.disableNotification = true
	return sv
}

func (sv *sendContact) SetEnableNotification() *sendContact {
	sv.disableNotification = false
	return sv
}

func (sv *sendContact) SetReplyToMessageId(messageId int64) *sendContact {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendContact) SetReplyMarkup(markup interface{}) *sendContact {
	sv.replyMarkup = tools.ParseReplyMarkup(markup)

	return sv
}

/*func (sv *sendContact) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
