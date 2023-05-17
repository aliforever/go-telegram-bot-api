package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type copyMessage struct {
	parent                   *TelegramBot
	chatId                   interface{}
	fromChatId               interface{}
	messageId                int64
	caption                  string
	parseMode                string
	captionEntities          []structs.MessageEntity
	disableNotification      bool
	replyToMessageId         int64
	allowSendingWithoutReply bool
	replyMarkup              interface{}
}

func (cm *copyMessage) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                   interface{}             `json:"chat_id"`
		FromChatId               interface{}             `json:"from_chat_id"`
		MessageId                int64                   `json:"message_id"`
		Caption                  string                  `json:"caption,omitempty"`
		CaptionEntities          []structs.MessageEntity `json:"caption_entities,omitempty"`
		ParseMode                string                  `json:"parse_mode,omitempty"`
		DisableNotification      bool                    `json:"disable_notification,omitempty"`
		ReplyToMessageId         int64                   `json:"reply_to_message_id,omitempty"`
		AllowSendingWithoutReply bool                    `json:"allow_sending_without_reply,omitempty"`
		ReplyMarkup              interface{}             `json:"reply_markup,omitempty"`
	}{
		ChatId:                   cm.chatId,
		FromChatId:               cm.fromChatId,
		MessageId:                cm.messageId,
		Caption:                  cm.caption,
		CaptionEntities:          cm.captionEntities,
		ParseMode:                cm.parseMode,
		DisableNotification:      cm.disableNotification,
		ReplyToMessageId:         cm.replyToMessageId,
		AllowSendingWithoutReply: false,
		ReplyMarkup:              cm.replyMarkup,
	})
}

func (cm *copyMessage) response() interface{} {
	return &structs.MessageId{}
}

func (cm *copyMessage) SetChatId(chatId int64) *copyMessage {
	cm.chatId = chatId
	return cm
}

func (cm *copyMessage) SetChatUsername(username string) *copyMessage {
	cm.chatId = username
	return cm
}

func (cm *copyMessage) SetFromChatId(fromChatId int64) *copyMessage {
	cm.fromChatId = fromChatId
	return cm
}

func (cm *copyMessage) SetFromChatUsername(fromChatUsername string) *copyMessage {
	cm.fromChatId = fromChatUsername
	return cm
}

func (cm *copyMessage) SetMessageId(messageId int64) *copyMessage {
	cm.messageId = messageId
	return cm
}

func (cm *copyMessage) SetCaption(caption string) *copyMessage {
	cm.caption = caption
	return cm
}

func (cm *copyMessage) SetParseModeHTML() *copyMessage {
	cm.parseMode = "HTML"
	return cm
}

func (cm *copyMessage) SetParseModeMarkdown() *copyMessage {
	cm.parseMode = "Markdown"
	return cm
}

func (cm *copyMessage) SetReplyToMessageId(messageId int64) *copyMessage {
	cm.replyToMessageId = messageId
	return cm
}

func (cm *copyMessage) SetAllowSendingWithoutReply() *copyMessage {
	cm.allowSendingWithoutReply = true
	return cm
}

func (cm *copyMessage) SetDisallowSendingWithoutReply() *copyMessage {
	cm.allowSendingWithoutReply = false
	return cm
}

func (cm *copyMessage) SetEnableNotification() *copyMessage {
	cm.disableNotification = false
	return cm
}

func (cm *copyMessage) SetDisableNotification() *copyMessage {
	cm.disableNotification = true
	return cm
}

func (cm *copyMessage) SetMessage(message *structs.Message) *copyMessage {
	cm.messageId = message.MessageId
	cm.fromChatId = message.Chat.Id
	return cm
}

func (cm *copyMessage) SetReplyMarkup(markup interface{}) *copyMessage {
	cm.replyMarkup = markup
	return cm
}

func (cm *copyMessage) method() string {
	return "POST"
}

func (cm *copyMessage) endpoint() string {
	return "copyMessage"
}

/*func (cm *copyMessage) Forward() (message *structs.Message, err error) {
	message, err = cm.parent.Send(fm)
	return
}
*/
