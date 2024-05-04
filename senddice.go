package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

const (
	emojiDice        = "🎲"
	emojiDart        = "🎯"
	emojiBasketball  = "🏀"
	emojiSoccer      = "⚽"
	emojiBowling     = "🎳"
	emojiSlotMachine = "🎰"
)

type sendDice struct {
	parent *TelegramBot

	chatId                interface{}
	emoji                 string
	parseMode             string
	disableWebPagePreview bool
	disableNotification   bool
	replyToMessageId      int64
	replyMarkup           interface{}
}

func (sph *sendDice) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                interface{} `json:"chat_id"`
		Emoji                 string      `json:"photo,omitempty"`
		Caption               string      `json:"caption,omitempty"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		DisableNotification   bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:                sph.chatId,
		Emoji:                 sph.emoji,
		ParseMode:             sph.parseMode,
		DisableWebPagePreview: sph.disableWebPagePreview,
		DisableNotification:   sph.disableNotification,
		ReplyToMessageId:      sph.replyToMessageId,
		ReplyMarkup:           sph.replyMarkup,
	})
}

func (sph *sendDice) response() interface{} {
	return &structs.Message{}
}

func (sph *sendDice) method() string {
	return "POST"
}

func (sph *sendDice) endpoint() string {
	return "sendDice"
}

func (sph *sendDice) SetChatId(chatId int64) *sendDice {
	sph.chatId = chatId

	return sph
}

func (sph *sendDice) SetDiceEmoji() *sendDice {
	sph.emoji = emojiDice

	return sph
}

func (sph *sendDice) SetDartEmoji() *sendDice {
	sph.emoji = emojiDart

	return sph
}

func (sph *sendDice) SetBasketballEmoji() *sendDice {
	sph.emoji = emojiBasketball

	return sph
}

func (sph *sendDice) SetSoccerEmoji() *sendDice {
	sph.emoji = emojiSoccer

	return sph
}

func (sph *sendDice) SetBowlingEmoji() *sendDice {
	sph.emoji = emojiBowling

	return sph
}

func (sph *sendDice) SetSlotMachineEmoji() *sendDice {
	sph.emoji = emojiSlotMachine

	return sph
}

func (sph *sendDice) SetChatUsername(username string) *sendDice {
	sph.chatId = username

	return sph
}

func (sph *sendDice) SetParseMode(parseMode string) *sendDice {
	sph.parseMode = parseMode
	return sph
}

func (sph *sendDice) SetParseModeHTML() *sendDice {
	sph.parseMode = "HTML"
	return sph
}

func (sph *sendDice) SetParseModeMarkdown() *sendDice {
	sph.parseMode = "Markdown"
	return sph
}

func (sph *sendDice) SetDisableWebPagePreview() *sendDice {
	sph.disableWebPagePreview = true
	return sph
}

func (sph *sendDice) SetEnableWebPagePreview() *sendDice {
	sph.disableWebPagePreview = false
	return sph
}

func (sph *sendDice) SetDisableNotification() *sendDice {
	sph.disableNotification = true
	return sph
}

func (sph *sendDice) SetEnableNotification() *sendDice {
	sph.disableNotification = false
	return sph
}

func (sph *sendDice) SetReplyToMessageId(messageId int64) *sendDice {
	sph.replyToMessageId = messageId
	return sph
}

func (sph *sendDice) SetReplyMarkup(markup interface{}) *sendDice {
	sph.replyMarkup = tools.ParseReplyMarkup(markup)

	return sph
}

func (sph *sendDice) SetReplyMarkupFromSlicesOfStrings(rows [][]string, resize, oneTime, selective bool) *sendDice {
	b := structs.ReplyKeyboardMarkup{}.FromSlicesOfStrings(rows)
	b.SetResizeKeyboard(resize)
	b.SetOneTimeKeyboard(oneTime)
	b.SetSelective(selective)
	var keyboard interface{} = b
	sph.replyMarkup = &keyboard
	return sph
}

/*func (sph *sendDice) Send() (message *structs.Message, err error) {
	message, err = sph.parent.Send(sph)
	return
}
*/
