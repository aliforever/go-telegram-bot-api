package tgbotapi

import (
	"encoding/json"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendAudio struct {
	parent              *TelegramBot
	chatId              interface{}
	audio               interface{}
	caption             string
	parseMode           string
	duration            int64
	performer           string
	title               string
	thumb               interface{}
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	file
	audios []fileInfo
}

func (sph *sendAudio) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Audio               interface{} `json:"audio,omitempty"`
		Caption             string      `json:"caption,omitempty"`
		ParseMode           string      `json:"parse_mode,omitempty"`
		Duration            int64       `json:"duration,omitempty"`
		Performer           string      `json:"performer,omitempty"`
		Title               string      `json:"title,omitempty"`
		Thumb               interface{} `json:"thumb,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sph.chatId,
		Audio:               sph.audio,
		Caption:             sph.caption,
		ParseMode:           sph.parseMode,
		Duration:            sph.duration,
		Performer:           sph.performer,
		Title:               sph.title,
		Thumb:               sph.thumb,
		DisableNotification: sph.disableNotification,
		ReplyToMessageId:    sph.replyToMessageId,
		ReplyMarkup:         sph.replyMarkup,
	})
}

func (sph *sendAudio) response() interface{} {
	return &structs.Message{}
}

func (sph *sendAudio) method() string {
	return "POST"
}

func (sph *sendAudio) endpoint() string {
	return "sendAudio"
}

func (sph *sendAudio) medias() []fileInfo {
	return sph.audios
}

func (sph *sendAudio) SetChatId(chatId int64) *sendAudio {
	sph.chatId = chatId
	return sph
}

func (sph *sendAudio) SetChatUsername(username string) *sendAudio {
	sph.chatId = username
	return sph
}

func (sph *sendAudio) SetAudioId(audioId string) *sendAudio {
	sph.audio = audioId
	return sph
}

func (sph *sendAudio) SetAudioFilePath(photoFilePath string) *sendAudio {
	if sph.audios == nil {
		sph.audios = []fileInfo{}
	}
	sph.audios = append(sph.audios, fileInfo{
		Path:  photoFilePath,
		Field: "voice",
	})
	return sph
}

func (sph *sendAudio) SetAudioReader(phr io.Reader, audioName string) *sendAudio {
	if sph.audios == nil {
		sph.audios = []fileInfo{}
	}
	sph.audios = append(sph.audios, fileInfo{
		Reader: phr,
		Name:   audioName,
		Field:  "voice",
	})
	return sph
}

func (sph *sendAudio) SetAudioThumbId(audioThumbId string) *sendAudio {
	sph.thumb = audioThumbId
	return sph
}

func (sph *sendAudio) SetAudioThumbFilePath(audioThumbPath string) *sendAudio {
	if sph.audios == nil {
		sph.audios = []fileInfo{}
	}
	sph.audios = append(sph.audios, fileInfo{
		Path:  audioThumbPath,
		Field: "thumb",
	})
	return sph
}

func (sph *sendAudio) SetAudioThumbReader(phr io.Reader, audioThumbName string) *sendAudio {
	if sph.audios == nil {
		sph.audios = []fileInfo{}
	}
	sph.audios = append(sph.audios, fileInfo{
		Reader: phr,
		Name:   audioThumbName,
		Field:  "thumb",
	})
	return sph
}

func (sph *sendAudio) SetCaption(caption string) *sendAudio {
	sph.caption = caption
	return sph
}

func (sph *sendAudio) SetParseMode(parseMode string) *sendAudio {
	sph.parseMode = parseMode
	return sph
}

func (sph *sendAudio) SetParseModeHTML() *sendAudio {
	sph.parseMode = "HTML"
	return sph
}

func (sph *sendAudio) SetParseModeMarkdown() *sendAudio {
	sph.parseMode = "Markdown"
	return sph
}

/*func (sph *sendAudio) SetDisableWebPagePreview() *sendAudio {
	sph.disableWebPagePreview = true
	return sph
}

func (sph *sendAudio) SetEnableWebPagePreview() *sendAudio {
	sph.disableWebPagePreview = false
	return sph
}*/

func (sph *sendAudio) SetDisableNotification() *sendAudio {
	sph.disableNotification = true
	return sph
}

func (sph *sendAudio) SetEnableNotification() *sendAudio {
	sph.disableNotification = false
	return sph
}

func (sph *sendAudio) SetReplyToMessageId(messageId int64) *sendAudio {
	sph.replyToMessageId = messageId
	return sph
}

func (sph *sendAudio) SetReplyMarkup(markup interface{}) *sendAudio {
	sph.replyMarkup = &markup
	return sph
}

/*func (sph *sendAudio) Send() (message *structs.Message, err error) {
	message, err = sph.parent.Send(sph)
	return
}
*/
