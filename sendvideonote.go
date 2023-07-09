package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendVideoNote struct {
	parent              *TelegramBot
	chatId              interface{}
	videoNote           interface{}
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
	videoNotes []fileInfo
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
}

func (sv *sendVideoNote) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		VideoNote           interface{} `json:"video_note,omitempty"`
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
		VideoNote:           sv.videoNote,
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

func (sv *sendVideoNote) response() interface{} {
	return &structs.Message{}
}

func (sv *sendVideoNote) method() string {
	return "POST"
}

func (sv *sendVideoNote) endpoint() string {
	return "sendVideoNote"
}

func (sv *sendVideoNote) medias() []fileInfo {
	return sv.videoNotes
}

func (sv *sendVideoNote) SetChatId(chatId int64) *sendVideoNote {
	sv.chatId = chatId
	return sv
}

func (sv *sendVideoNote) SetChatUsername(username string) *sendVideoNote {
	sv.chatId = username
	return sv
}

func (sv *sendVideoNote) SetVideoNoteId(videoId string) *sendVideoNote {
	sv.videoNote = videoId
	return sv
}

func (sv *sendVideoNote) SetVideoNoteFilePath(photoFilePath string) *sendVideoNote {
	if sv.videoNotes == nil {
		sv.videoNotes = []fileInfo{}
	}
	sv.videoNotes = append(sv.videoNotes, fileInfo{
		Path:  photoFilePath,
		Field: "video_note",
	})
	return sv
}

func (sv *sendVideoNote) SetVideoNoteReader(phr io.Reader, videoName string) *sendVideoNote {
	if sv.videoNotes == nil {
		sv.videoNotes = []fileInfo{}
	}
	sv.videoNotes = append(sv.videoNotes, fileInfo{
		Reader: phr,
		Name:   videoName,
		Field:  "video_note",
	})
	return sv
}

func (sv *sendVideoNote) SetThumbId(videoThumbId string) *sendVideoNote {
	sv.thumb = videoThumbId
	return sv
}

func (sv *sendVideoNote) SetThumbFilePath(videoThumbPath string) *sendVideoNote {
	if sv.videoNotes == nil {
		sv.videoNotes = []fileInfo{}
	}
	sv.videoNotes = append(sv.videoNotes, fileInfo{
		Path:  videoThumbPath,
		Field: "thumb",
	})
	return sv
}

func (sv *sendVideoNote) SetThumbReader(phr io.Reader, videoThumbName string) *sendVideoNote {
	if sv.videoNotes == nil {
		sv.videoNotes = []fileInfo{}
	}
	sv.videoNotes = append(sv.videoNotes, fileInfo{
		Reader: phr,
		Name:   videoThumbName,
		Field:  "thumb",
	})
	return sv
}

func (sv *sendVideoNote) SetCaption(caption string) *sendVideoNote {
	sv.caption = caption
	return sv
}

func (sv *sendVideoNote) SetParseMode(parseMode string) *sendVideoNote {
	sv.parseMode = parseMode
	return sv
}

func (sv *sendVideoNote) SetParseModeHTML() *sendVideoNote {
	sv.parseMode = "HTML"
	return sv
}

func (sv *sendVideoNote) SetParseModeMarkdown() *sendVideoNote {
	sv.parseMode = "Markdown"
	return sv
}

func (sv *sendVideoNote) SetDisableNotification() *sendVideoNote {
	sv.disableNotification = true
	return sv
}

func (sv *sendVideoNote) SetEnableNotification() *sendVideoNote {
	sv.disableNotification = false
	return sv
}

func (sv *sendVideoNote) SetReplyToMessageId(messageId int64) *sendVideoNote {
	sv.replyToMessageId = messageId
	return sv
}

func (sv *sendVideoNote) SetReplyMarkup(markup interface{}) *sendVideoNote {
	sv.replyMarkup = tools.ParseReplyMarkup(markup)

	return sv
}

/*func (sv *sendVideoNote) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sv)
	return
}
*/
/*func (sv *sendVideoNote) SetDisableWebPagePreview() *sendVideoNote {
  	sv.DisableWebPagePreview = true
  	return sv
  }

  func (sv *sendVideoNote) SetEnableWebPagePreview() *sendVideoNote {
  	sv.DisableWebPagePreview = false
  	return sv
  }*/
