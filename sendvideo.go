package tgbotapi

import (
	"encoding/json"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendVideo struct {
	parent              *TelegramBot
	chatId              interface{}
	video               interface{}
	caption             string
	parseMode           string
	duration            int64
	width               int64
	height              int64
	supportsStreaming   bool
	performer           string
	title               string
	thumb               interface{}
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	file
	videos []fileInfo
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendVideo) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Video               interface{} `json:"video,omitempty"`
		Caption             string      `json:"caption,omitempty"`
		ParseMode           string      `json:"parse_mode,omitempty"`
		Duration            int64       `json:"duration,omitempty"`
		Width               int64       `json:"width,omitempty"`
		Height              int64       `json:"height,omitempty"`
		SupportsStreaming   bool        `json:"supports_streaming,omitempty"`
		Performer           string      `json:"performer,omitempty"`
		Title               string      `json:"title,omitempty"`
		Thumb               interface{} `json:"thumb,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sv.chatId,
		Video:               sv.video,
		Caption:             sv.caption,
		ParseMode:           sv.parseMode,
		Duration:            sv.duration,
		Width:               sv.width,
		Height:              sv.height,
		SupportsStreaming:   sv.supportsStreaming,
		Performer:           sv.performer,
		Title:               sv.title,
		Thumb:               sv.thumb,
		DisableNotification: sv.disableNotification,
		ReplyToMessageId:    sv.replyToMessageId,
		ReplyMarkup:         sv.replyMarkup,
	})
}

func (sv *sendVideo) response() interface{} {
	return &structs.Message{}
}

func (sv *sendVideo) method() string {
	return "POST"
}

func (sv *sendVideo) endpoint() string {
	return "sendVideo"
}

func (sv *sendVideo) medias() []fileInfo {
	return sv.videos
}

func (sv *sendVideo) SetChatId(chatId int64) *sendVideo {
	sv.chatId = chatId
	return sv
}

func (sv *sendVideo) SetChatUsername(username string) *sendVideo {
	sv.chatId = username
	return sv
}

func (sv *sendVideo) SetVideoId(videoId string) *sendVideo {
	sv.video = videoId
	return sv
}

func (sv *sendVideo) SetVideoFilePath(photoFilePath string) *sendVideo {
	if sv.videos == nil {
		sv.videos = []fileInfo{}
	}
	sv.videos = append(sv.videos, fileInfo{
		Path:  photoFilePath,
		Field: "video",
	})
	sv.video = "attach://video"
	return sv
}

func (sv *sendVideo) SetVideoReader(phr io.Reader, videoName string) *sendVideo {
	if sv.videos == nil {
		sv.videos = []fileInfo{}
	}
	sv.videos = append(sv.videos, fileInfo{
		Reader: phr,
		Name:   videoName,
		Field:  "video",
	})
	sv.video = "attach://video"
	return sv
}

func (sv *sendVideo) SetVideoThumbId(videoThumbId string) *sendVideo {
	sv.thumb = videoThumbId
	return sv
}

func (sv *sendVideo) SetVideoThumbFilePath(videoThumbPath string) *sendVideo {
	if sv.videos == nil {
		sv.videos = []fileInfo{}
	}
	sv.videos = append(sv.videos, fileInfo{
		Path:  videoThumbPath,
		Field: "thumb",
	})
	sv.thumb = "attach://thumb"
	return sv
}

func (sv *sendVideo) SetVideoThumbReader(phr io.Reader, videoThumbName string) *sendVideo {
	if sv.videos == nil {
		sv.videos = []fileInfo{}
	}
	sv.videos = append(sv.videos, fileInfo{
		Reader: phr,
		Name:   videoThumbName,
		Field:  "thumb",
	})
	sv.thumb = "attach://thumb"
	return sv
}

func (sv *sendVideo) SetCaption(caption string) *sendVideo {
	sv.caption = caption
	return sv
}

func (sv *sendVideo) SetParseMode(parseMode string) *sendVideo {
	sv.parseMode = parseMode
	return sv
}

func (sv *sendVideo) SetParseModeHTML() *sendVideo {
	sv.parseMode = "HTML"
	return sv
}

func (sv *sendVideo) SetParseModeMarkdown() *sendVideo {
	sv.parseMode = "Markdown"
	return sv
}

func (sv *sendVideo) SetDisableNotification() *sendVideo {
	sv.disableNotification = true
	return sv
}

func (sv *sendVideo) SetEnableNotification() *sendVideo {
	sv.disableNotification = false
	return sv
}

func (sv *sendVideo) SetReplyToMessageId(messageId int64) *sendVideo {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendVideo) SetReplyMarkup(markup interface{}) *sendVideo {
	sv.replyMarkup = &markup
	return sv
}

/*func (sv *sendVideo) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
/*func (sv *sendVideo) SetDisableWebPagePreview() *sendVideo {
  	sv.DisableWebPagePreview = true
  	return sv
  }

  func (sv *sendVideo) SetEnableWebPagePreview() *sendVideo {
  	sv.DisableWebPagePreview = false
  	return sv
  }*/
