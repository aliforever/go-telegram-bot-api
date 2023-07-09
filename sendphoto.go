package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendPhoto struct {
	parent                *TelegramBot
	chatId                interface{}
	photo                 interface{}
	caption               string
	parseMode             string
	disableWebPagePreview bool
	disableNotification   bool
	replyToMessageId      int64
	replyMarkup           interface{}

	fileInfo *fileInfo
}

func (sph *sendPhoto) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                interface{} `json:"chat_id"`
		Photo                 interface{} `json:"photo,omitempty"`
		Caption               string      `json:"caption,omitempty"`
		ParseMode             string      `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
		DisableNotification   bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:                sph.chatId,
		Photo:                 sph.photo,
		Caption:               sph.caption,
		ParseMode:             sph.parseMode,
		DisableWebPagePreview: sph.disableWebPagePreview,
		DisableNotification:   sph.disableNotification,
		ReplyToMessageId:      sph.replyToMessageId,
		ReplyMarkup:           sph.replyMarkup,
	})
}

func (sph *sendPhoto) response() interface{} {
	return &structs.Message{}
}

func (sph *sendPhoto) method() string {
	return "POST"
}

func (sph *sendPhoto) endpoint() string {
	return "sendPhoto"
}

func (sph *sendPhoto) medias() []fileInfo {
	if sph.fileInfo != nil {
		return []fileInfo{*sph.fileInfo}
	}

	return nil
}

func (sph *sendPhoto) SetChatId(chatId int64) *sendPhoto {
	sph.chatId = chatId
	return sph
}

func (sph *sendPhoto) SetChatUsername(username string) *sendPhoto {
	sph.chatId = username
	return sph
}

func (sph *sendPhoto) SetPhotoId(photoId string) *sendPhoto {
	sph.photo = photoId
	return sph
}

func (sph *sendPhoto) SetPhotoFilePath(photoFilePath string) *sendPhoto {
	sph.fileInfo = &fileInfo{
		Field: "photo",
		Path:  photoFilePath,
	}
	sph.photo = "attach://photo"

	return sph
}

func (sph *sendPhoto) SetPhotoReader(phr io.Reader, photoName string) *sendPhoto {
	sph.fileInfo = &fileInfo{
		Field:  "photo",
		Reader: phr,
		Name:   photoName,
	}
	sph.photo = "attach://photo"

	return sph
}

func (sph *sendPhoto) SetCaption(caption string) *sendPhoto {
	sph.caption = caption
	return sph
}

func (sph *sendPhoto) SetParseMode(parseMode string) *sendPhoto {
	sph.parseMode = parseMode
	return sph
}

func (sph *sendPhoto) SetParseModeHTML() *sendPhoto {
	sph.parseMode = "HTML"
	return sph
}

func (sph *sendPhoto) SetParseModeMarkdown() *sendPhoto {
	sph.parseMode = "Markdown"
	return sph
}

func (sph *sendPhoto) SetDisableWebPagePreview() *sendPhoto {
	sph.disableWebPagePreview = true
	return sph
}

func (sph *sendPhoto) SetEnableWebPagePreview() *sendPhoto {
	sph.disableWebPagePreview = false
	return sph
}

func (sph *sendPhoto) SetDisableNotification() *sendPhoto {
	sph.disableNotification = true
	return sph
}

func (sph *sendPhoto) SetEnableNotification() *sendPhoto {
	sph.disableNotification = false
	return sph
}

func (sph *sendPhoto) SetReplyToMessageId(messageId int64) *sendPhoto {
	sph.replyToMessageId = messageId
	return sph
}

func (sph *sendPhoto) SetReplyMarkup(markup interface{}) *sendPhoto {
	sph.replyMarkup = tools.ParseReplyMarkup(markup)

	return sph
}

func (sph *sendPhoto) SetReplyMarkupFromSlicesOfStrings(rows [][]string, resize, oneTime, selective bool) *sendPhoto {
	b := structs.ReplyKeyboardMarkup{}.FromSlicesOfStrings(rows)
	b.SetResizeKeyboard(resize)
	b.SetOneTimeKeyboard(oneTime)
	b.SetSelective(selective)
	var keyboard interface{} = b
	sph.replyMarkup = &keyboard
	return sph
}

/*func (sph *sendPhoto) Send() (message *structs.Message, err error) {
	message, err = sph.parent.Send(sph)
	return
}
*/
