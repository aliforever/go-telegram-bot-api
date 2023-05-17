package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendPoll struct {
	parent                *TelegramBot
	chatId                interface{}
	question              string
	options               []string
	isAnonymous           bool
	pollType              string
	allowsMultipleAnswers bool
	correctOptionId       int64
	isClosed              bool
	disableNotification   bool
	replyToMessageId      int64
	replyMarkup           interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendPoll) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                interface{} `json:"chat_id"`
		Question              string      `json:"question"`
		Options               []string    `json:"options"`
		IsAnonymous           bool        `json:"is_anonymous"`
		PollType              string      `json:"type,omitempty"`
		AllowsMultipleAnswers bool        `json:"allows_multiple_answers,omitempty"`
		CorrectOptionId       int64       `json:"correct_option_id,omitempty"`
		IsClosed              bool        `json:"is_closed,omitempty"`
		DisableNotification   bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:                sv.chatId,
		Question:              sv.question,
		Options:               sv.options,
		IsAnonymous:           sv.isAnonymous,
		PollType:              sv.pollType,
		AllowsMultipleAnswers: sv.allowsMultipleAnswers,
		CorrectOptionId:       sv.correctOptionId,
		IsClosed:              sv.isClosed,
		DisableNotification:   sv.disableNotification,
		ReplyToMessageId:      sv.replyToMessageId,
		ReplyMarkup:           sv.replyMarkup,
	})
}

func (sv *sendPoll) response() interface{} {
	return &structs.Message{}
}

func (sv *sendPoll) method() string {
	return "POST"
}

func (sv *sendPoll) endpoint() string {
	return "sendPoll"
}

func (sv *sendPoll) SetChatId(chatId int64) *sendPoll {
	sv.chatId = chatId
	return sv
}

func (sv *sendPoll) SetChatUsername(username string) *sendPoll {
	sv.chatId = username
	return sv
}

func (sv *sendPoll) SetQuestion(question string) *sendPoll {
	sv.question = question
	return sv
}

func (sv *sendPoll) SetOptions(options []string) *sendPoll {
	sv.options = options
	return sv
}

func (sv *sendPoll) EnableAnonymous() *sendPoll {
	sv.isAnonymous = true
	return sv
}

func (sv *sendPoll) DisableAnonymous() *sendPoll {
	sv.isAnonymous = false
	return sv
}

func (sv *sendPoll) SetType(pollType string) *sendPoll {
	sv.pollType = pollType
	return sv
}

func (sv *sendPoll) AllowMultipleOptions() *sendPoll {
	sv.allowsMultipleAnswers = true
	return sv
}

func (sv *sendPoll) DisallowMultipleOptions() *sendPoll {
	sv.allowsMultipleAnswers = false
	return sv
}

func (sv *sendPoll) SetCorrectOptionId(optionId int64) *sendPoll {
	sv.correctOptionId = optionId
	return sv
}

func (sv *sendPoll) Close() *sendPoll {
	sv.isClosed = true
	return sv
}

func (sv *sendPoll) Open() *sendPoll {
	sv.isClosed = false
	return sv
}

func (sv *sendPoll) SetDisableNotification() *sendPoll {
	sv.disableNotification = true
	return sv
}

func (sv *sendPoll) SetEnableNotification() *sendPoll {
	sv.disableNotification = false
	return sv
}

func (sv *sendPoll) SetReplyToMessageId(messageId int64) *sendPoll {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendPoll) SetReplyMarkup(markup interface{}) *sendPoll {
	sv.replyMarkup = &markup
	return sv
}

/*func (sv *sendPoll) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
