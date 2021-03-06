package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type leaveChat struct {
	parent *TelegramBot
	chatId interface{}
}

func (sv *leaveChat) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId interface{} `json:"chat_id"`
	}{
		ChatId: sv.chatId,
	})
}

func (sv *leaveChat) response() interface{} {
	return structs.ResponseTypeBool()
}

func (sv *leaveChat) method() string {
	return "POST"
}

func (sv *leaveChat) endpoint() string {
	return "leaveChat"
}

func (sv *leaveChat) SetChatId(chatId int64) *leaveChat {
	sv.chatId = chatId
	return sv
}

func (sv *leaveChat) SetChatUsername(username string) *leaveChat {
	sv.chatId = username
	return sv
}

/*func (sv *leaveChat) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
