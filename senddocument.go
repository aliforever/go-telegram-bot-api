package tgbotapi

import (
	"encoding/json"
	"github.com/aliforever/go-telegram-bot-api/tools"
	"io"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendDocument struct {
	parent              *TelegramBot
	chatId              interface{}
	document            interface{}
	thumb               interface{}
	caption             string
	parseMode           string
	disableNotification bool
	replyToMessageId    int64
	replyMarkup         interface{}
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	file
	documents []fileInfo
}

func (sd *sendDocument) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{} `json:"chat_id"`
		Document            interface{} `json:"document,omitempty"`
		Thumb               interface{} `json:"thumb,omitempty"`
		Caption             string      `json:"caption,omitempty"`
		ParseMode           string      `json:"parse_mode,omitempty"`
		DisableNotification bool        `json:"disable_notification,omitempty"`
		ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
		ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	}{
		ChatId:              sd.chatId,
		Document:            sd.document,
		Thumb:               sd.thumb,
		Caption:             sd.caption,
		ParseMode:           sd.parseMode,
		DisableNotification: sd.disableNotification,
		ReplyToMessageId:    sd.replyToMessageId,
		ReplyMarkup:         sd.replyMarkup,
	})
}

func (sd *sendDocument) response() interface{} {
	return &structs.Message{}
}

func (sd *sendDocument) method() string {
	return "POST"
}

func (sd *sendDocument) endpoint() string {
	return "sendDocument"
}

func (sd *sendDocument) medias() []fileInfo {
	for i := range sd.documents {
		sd.documents[i].Field = "document"
	}
	return sd.documents
}

func (sd *sendDocument) SetChatId(chatId int64) *sendDocument {
	sd.chatId = chatId
	return sd
}

func (sd *sendDocument) SetChatUsername(username string) *sendDocument {
	sd.chatId = username
	return sd
}

func (sd *sendDocument) SetDocumentId(documentId string) *sendDocument {
	sd.document = documentId
	return sd
}

func (sd *sendDocument) SetDocumentFilePath(photoFilePath string) *sendDocument {
	if sd.documents == nil {
		sd.documents = []fileInfo{}
	}
	sd.documents = append(sd.documents, fileInfo{
		Path: photoFilePath,
	})
	return sd
}

func (sd *sendDocument) SetDocumentReader(phr io.Reader, documentName string) *sendDocument {
	if sd.documents == nil {
		sd.documents = []fileInfo{}
	}
	sd.documents = append(sd.documents, fileInfo{
		Reader: phr,
		Name:   documentName,
	})
	return sd
}

func (sd *sendDocument) SetDocumentThumbId(documentThumbId string) *sendDocument {
	sd.thumb = documentThumbId
	return sd
}

func (sd *sendDocument) SetDocumentThumbFilePath(documentThumbPath string) *sendDocument {
	if sd.documents == nil {
		sd.documents = []fileInfo{}
	}
	sd.documents = append(sd.documents, fileInfo{
		Path: documentThumbPath,
	})
	return sd
}

func (sd *sendDocument) SetDocumentThumbReader(phr io.Reader, documentThumbName string) *sendDocument {
	if sd.documents == nil {
		sd.documents = []fileInfo{}
	}
	sd.documents = append(sd.documents, fileInfo{
		Reader: phr,
		Name:   documentThumbName,
	})
	return sd
}

func (sd *sendDocument) SetCaption(caption string) *sendDocument {
	sd.caption = caption
	return sd
}

func (sd *sendDocument) SetParseMode(parseMode string) *sendDocument {
	sd.parseMode = parseMode
	return sd
}

func (sd *sendDocument) SetParseModeHTML() *sendDocument {
	sd.parseMode = "HTML"
	return sd
}

func (sd *sendDocument) SetParseModeMarkdown() *sendDocument {
	sd.parseMode = "Markdown"
	return sd
}

/*func (sd *sendDocument) SetDisableWebPagePreview() *sendDocument {
	sd.disableWebPagePreview = true
	return sd
}

func (sd *sendDocument) SetEnableWebPagePreview() *sendDocument {
	sd.disableWebPagePreview = false
	return sd
}*/

func (sd *sendDocument) SetDisableNotification() *sendDocument {
	sd.disableNotification = true
	return sd
}

func (sd *sendDocument) SetEnableNotification() *sendDocument {
	sd.disableNotification = false
	return sd
}

func (sd *sendDocument) SetReplyToMessageId(messageId int64) *sendDocument {
	sd.replyToMessageId = messageId
	return sd
}

func (sd *sendDocument) SetReplyMarkup(markup interface{}) *sendDocument {
	sd.replyMarkup = tools.ParseReplyMarkup(markup)

	return sd
}

/*func (sd *sendDocument) Send() (message *structs.Message, err error) {
	message, err = sd.parent.Send(sd)
	return
}
*/
