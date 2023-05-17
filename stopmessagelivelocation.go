package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type stopMessageLiveLocation struct {
	parent          *TelegramBot
	chatId          interface{}
	messageId       int64
	inlineMessageId string
	replyMarkup     interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (smll *stopMessageLiveLocation) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId          interface{} `json:"chat_id,omitempty"`
		MessageId       int64       `json:"message_id,omitempty"`
		InlineMessageId string      `json:"inline_message_id,omitempty"`
		ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:          smll.chatId,
		MessageId:       smll.messageId,
		InlineMessageId: smll.inlineMessageId,
		ReplyMarkup:     smll.replyMarkup,
	})
}

func (smll *stopMessageLiveLocation) response() interface{} {
	return &structs.Message{}
}

func (smll *stopMessageLiveLocation) method() string {
	return "POST"
}

func (smll *stopMessageLiveLocation) endpoint() string {
	return "stopMessageLiveLocation"
}

func (smll *stopMessageLiveLocation) SetChatId(chatId int64) *stopMessageLiveLocation {
	smll.chatId = chatId
	return smll
}

func (smll *stopMessageLiveLocation) SetChatUsername(username string) *stopMessageLiveLocation {
	smll.chatId = username
	return smll
}

func (smll *stopMessageLiveLocation) SetMessageId(messageId int64) *stopMessageLiveLocation {
	smll.messageId = messageId
	return smll
}

func (smll *stopMessageLiveLocation) SetInlineMessageId(inlineMessageId string) *stopMessageLiveLocation {
	smll.inlineMessageId = inlineMessageId
	return smll
}

func (smll *stopMessageLiveLocation) SetReplyMarkup(markup interface{}) *stopMessageLiveLocation {
	smll.replyMarkup = &markup
	return smll
}

/*func (smll *stopMessageLiveLocation) Send() (message *structs.Message, err error) {
	message, err = smll.parent.Send(sd)
	return
}
*/
