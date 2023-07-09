package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"
	"io"
	"time"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendSticker struct {
	parent              *TelegramBot
	chatId              interface{}
	sticker             string
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	file
	fileInfo *fileInfo
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendSticker) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Sticker             string      `json:"sticker"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sv.chatId,
		Sticker:             sv.sticker,
		DisableNotification: sv.disableNotification,
		ReplyToMessageId:    sv.replyToMessageId,
		ReplyMarkup:         sv.replyMarkup,
	})
}

func (sv *sendSticker) response() interface{} {
	return &structs.Message{}
}

func (sv *sendSticker) method() string {
	return "POST"
}

func (sv *sendSticker) endpoint() string {
	return "sendSticker"
}

func (sv *sendSticker) medias() []fileInfo {
	if sv.fileInfo != nil {
		return []fileInfo{*sv.fileInfo}
	}
	return nil
}

func (sv *sendSticker) SetChatId(chatId int64) *sendSticker {
	sv.chatId = chatId
	return sv
}

func (sv *sendSticker) SetChatUsername(username string) *sendSticker {
	sv.chatId = username
	return sv
}

func (sv *sendSticker) SetStickerFileId(stickerFileId string) *sendSticker {
	sv.sticker = stickerFileId
	return sv
}

func (sv *sendSticker) SetStickerFilePath(stickerFilePath string) *sendSticker {
	sv.fileInfo = &fileInfo{
		Field: "sticker",
		Path:  stickerFilePath,
	}
	return sv
}

func (sv *sendSticker) SetStickerFileReader(stickerFileReader io.Reader, fileName string) *sendSticker {
	if fileName == "" {
		fileName = time.Now().Format("2006_01_02_15_04_05")
	}
	sv.fileInfo = &fileInfo{
		Field:  "sticker",
		Reader: stickerFileReader,
		Name:   fileName,
	}
	return sv
}

func (sv *sendSticker) SetDisableNotification() *sendSticker {
	sv.disableNotification = true
	return sv
}

func (sv *sendSticker) SetEnableNotification() *sendSticker {
	sv.disableNotification = false
	return sv
}

func (sv *sendSticker) SetReplyToMessageId(messageId int64) *sendSticker {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendSticker) SetReplyMarkup(markup interface{}) *sendSticker {
	sv.replyMarkup = tools.ParseReplyMarkup(markup)

	return sv
}

/*func (sv *sendSticker) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
