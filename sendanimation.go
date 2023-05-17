package tgbotapi

import (
	"encoding/json"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendAnimation struct {
	parent              *TelegramBot
	chatId              interface{}
	animation           interface{}
	duration            int64
	width               int64
	height              int64
	thumb               interface{}
	caption             string
	parseMode           string
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	file
	animations []fileInfo
}

func (sd *sendAnimation) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Animation           interface{} `json:"animation,omitempty"`
		Duration            int64       `json:"duration"`
		Width               int64       `json:"width"`
		Height              int64       `json:"height"`
		Thumb               interface{} `json:"thumb,omitempty"`
		Caption             string      `json:"caption,omitempty"`
		ParseMode           string      `json:"parse_mode,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sd.chatId,
		Animation:           sd.animation,
		Duration:            sd.duration,
		Width:               sd.width,
		Height:              sd.height,
		Thumb:               sd.thumb,
		Caption:             sd.caption,
		ParseMode:           sd.parseMode,
		DisableNotification: sd.disableNotification,
		ReplyToMessageId:    sd.replyToMessageId,
		ReplyMarkup:         sd.replyMarkup,
	})
}

func (sd *sendAnimation) response() interface{} {
	return &structs.Message{}
}

func (sd *sendAnimation) method() string {
	return "POST"
}

func (sd *sendAnimation) endpoint() string {
	return "sendAnimation"
}

func (sd *sendAnimation) medias() []fileInfo {
	for i := range sd.animations {
		sd.animations[i].Field = "animation"
	}
	return sd.animations
}

func (sd *sendAnimation) SetChatId(chatId int64) *sendAnimation {
	sd.chatId = chatId
	return sd
}

func (sd *sendAnimation) SetChatUsername(username string) *sendAnimation {
	sd.chatId = username
	return sd
}

func (sd *sendAnimation) SetWidth(width int64) {
	sd.width = width
}

func (sd *sendAnimation) SetDuration(duration int64) {
	sd.duration = duration
}

func (sd *sendAnimation) SetHeight(height int64) {
	sd.height = height
}

func (sd *sendAnimation) SetAnimationId(animationId string) *sendAnimation {
	sd.animation = animationId
	return sd
}

func (sd *sendAnimation) SetAnimationFilePath(photoFilePath string) *sendAnimation {
	if sd.animations == nil {
		sd.animations = []fileInfo{}
	}
	sd.animations = append(sd.animations, fileInfo{
		Path: photoFilePath,
	})
	return sd
}

func (sd *sendAnimation) SetAnimationReader(phr io.Reader, animationName string) *sendAnimation {
	if sd.animations == nil {
		sd.animations = []fileInfo{}
	}
	sd.animations = append(sd.animations, fileInfo{
		Reader: phr,
		Name:   animationName,
	})
	return sd
}

func (sd *sendAnimation) SetAnimationThumbId(animationThumbId string) *sendAnimation {
	sd.thumb = animationThumbId
	return sd
}

func (sd *sendAnimation) SetAnimationThumbFilePath(animationThumbPath string) *sendAnimation {
	if sd.animations == nil {
		sd.animations = []fileInfo{}
	}
	sd.animations = append(sd.animations, fileInfo{
		Path: animationThumbPath,
	})
	return sd
}

func (sd *sendAnimation) SetAnimationThumbReader(phr io.Reader, animationThumbName string) *sendAnimation {
	if sd.animations == nil {
		sd.animations = []fileInfo{}
	}
	sd.animations = append(sd.animations, fileInfo{
		Reader: phr,
		Name:   animationThumbName,
	})
	return sd
}

func (sd *sendAnimation) SetCaption(caption string) *sendAnimation {
	sd.caption = caption
	return sd
}

func (sd *sendAnimation) SetParseMode(parseMode string) *sendAnimation {
	sd.parseMode = parseMode
	return sd
}

func (sd *sendAnimation) SetParseModeHTML() *sendAnimation {
	sd.parseMode = "HTML"
	return sd
}

func (sd *sendAnimation) SetParseModeMarkdown() *sendAnimation {
	sd.parseMode = "Markdown"
	return sd
}

/*func (sd *sendAnimation) SetDisableWebPagePreview() *sendAnimation {
	sd.disableWebPagePreview = true
	return sd
}

func (sd *sendAnimation) SetEnableWebPagePreview() *sendAnimation {
	sd.disableWebPagePreview = false
	return sd
}*/

func (sd *sendAnimation) SetDisableNotification() *sendAnimation {
	sd.disableNotification = true
	return sd
}

func (sd *sendAnimation) SetEnableNotification() *sendAnimation {
	sd.disableNotification = false
	return sd
}

func (sd *sendAnimation) SetReplyToMessageId(messageId int64) *sendAnimation {
	sd.replyToMessageId = messageId
	return sd
}

func (sd *sendAnimation) SetReplyMarkup(markup interface{}) *sendAnimation {
	sd.replyMarkup = &markup
	return sd
}

/*func (sd *sendAnimation) Send() (message *structs.Message, err error) {
	message, err = sd.parent.Send(sd)
	return
}
*/
