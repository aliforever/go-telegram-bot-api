package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendVoice struct {
	parent              *TelegramBot
	chatId              interface{}
	voice               interface{}
	caption             string
	parseMode           string
	duration            int64
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	file
	voices []fileInfo
}

func (sph *sendVoice) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Voice               interface{} `json:"voice,omitempty"`
		Caption             string      `json:"caption,omitempty"`
		ParseMode           string      `json:"parse_mode,omitempty"`
		Duration            int64       `json:"duration,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sph.chatId,
		Voice:               sph.voice,
		Caption:             sph.caption,
		ParseMode:           sph.parseMode,
		Duration:            sph.duration,
		DisableNotification: sph.disableNotification,
		ReplyToMessageId:    sph.replyToMessageId,
		ReplyMarkup:         sph.replyMarkup,
	})
}

func (sph *sendVoice) method() string {
	return "POST"
}

func (sph *sendVoice) endpoint() string {
	return "sendVoice"
}

func (sph *sendVoice) medias() []fileInfo {
	return sph.voices
}

func (sph *sendVoice) response() interface{} {
	return &structs.Message{}
}

func (sph *sendVoice) SetChatId(chatId int64) *sendVoice {
	sph.chatId = chatId
	return sph
}

func (sph *sendVoice) SetChatUsername(username string) *sendVoice {
	sph.chatId = username
	return sph
}

func (sph *sendVoice) SetVoiceId(audioId string) *sendVoice {
	sph.voice = audioId
	return sph
}

func (sph *sendVoice) SetVoiceFilePath(photoFilePath string) *sendVoice {
	if sph.voices == nil {
		sph.voices = []fileInfo{}
	}
	sph.voices = append(sph.voices, fileInfo{
		Path:  photoFilePath,
		Field: "voice",
	})
	return sph
}

func (sph *sendVoice) SetVoiceReader(phr io.Reader, audioName string) *sendVoice {
	if sph.voices == nil {
		sph.voices = []fileInfo{}
	}
	sph.voices = append(sph.voices, fileInfo{
		Reader: phr,
		Name:   audioName,
		Field:  "voice",
	})
	return sph
}

func (sph *sendVoice) SetVoiceThumbFilePath(audioThumbPath string) *sendVoice {
	if sph.voices == nil {
		sph.voices = []fileInfo{}
	}
	sph.voices = append(sph.voices, fileInfo{
		Path:  audioThumbPath,
		Field: "thumb",
	})
	return sph
}

func (sph *sendVoice) SetVoiceThumbReader(phr io.Reader, audioThumbName string) *sendVoice {
	if sph.voices == nil {
		sph.voices = []fileInfo{}
	}
	sph.voices = append(sph.voices, fileInfo{
		Reader: phr,
		Name:   audioThumbName,
		Field:  "thumb",
	})
	return sph
}

func (sph *sendVoice) SetCaption(caption string) *sendVoice {
	sph.caption = caption
	return sph
}

func (sph *sendVoice) SetParseMode(parseMode string) *sendVoice {
	sph.parseMode = parseMode
	return sph
}

func (sph *sendVoice) SetParseModeHTML() *sendVoice {
	sph.parseMode = "HTML"
	return sph
}

func (sph *sendVoice) SetParseModeMarkdown() *sendVoice {
	sph.parseMode = "Markdown"
	return sph
}

/*func (sph *sendVoice) SetDisableWebPagePreview() *sendVoice {
	sph.disableWebPagePreview = true
	return sph
}

func (sph *sendVoice) SetEnableWebPagePreview() *sendVoice {
	sph.disableWebPagePreview = false
	return sph
}*/

func (sph *sendVoice) SetDisableNotification() *sendVoice {
	sph.disableNotification = true
	return sph
}

func (sph *sendVoice) SetEnableNotification() *sendVoice {
	sph.disableNotification = false
	return sph
}

func (sph *sendVoice) SetReplyToMessageId(messageId int64) *sendVoice {
	sph.replyToMessageId = messageId
	return sph
}

func (sph *sendVoice) SetReplyMarkup(markup interface{}) *sendVoice {
	sph.replyMarkup = tools.ParseReplyMarkup(markup)

	return sph
}

/*func (sph *sendVoice) Send() (message *structs.Message, err error) {
	message, err = sph.parent.Send(sph)
	return
}
*/
