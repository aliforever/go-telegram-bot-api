package tgbotapi

import (
	"encoding/json"
	"io"
	"time"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type uploadStickerFile struct {
	parent *TelegramBot

	userId        int64
	sticker       interface{}
	stickerFormat string

	fileInfo *fileInfo
}

func (sv *uploadStickerFile) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId        int64       `json:"user_id"`
		Sticker       interface{} `json:"sticker"`
		StickerFormat string      `json:"sticker_format"`
	}{
		UserId:        sv.userId,
		Sticker:       sv.sticker,
		StickerFormat: sv.stickerFormat,
	})
}

func (sv *uploadStickerFile) response() interface{} {
	return &structs.File{}
}

func (sv *uploadStickerFile) method() string {
	return "POST"
}

func (sv *uploadStickerFile) endpoint() string {
	return "uploadStickerFile"
}

func (sv *uploadStickerFile) medias() []fileInfo {
	if sv.fileInfo != nil {
		return []fileInfo{*sv.fileInfo}
	}

	return nil
}

func (sv *uploadStickerFile) SetUserId(userId int64) *uploadStickerFile {
	sv.userId = userId
	return sv
}

func (sv *uploadStickerFile) SetStickerId(stickerID string) *uploadStickerFile {
	sv.sticker = stickerID
	return sv
}

func (sv *uploadStickerFile) SetStickerFilePath(stickerFilePath string) *uploadStickerFile {
	sv.fileInfo = &fileInfo{
		Field: "sticker",
		Path:  stickerFilePath,
	}

	sv.sticker = "attach://sticker"

	return sv
}

func (sv *uploadStickerFile) SetStickerFileReader(stickerFileReader io.Reader, fileName string) *uploadStickerFile {
	if fileName == "" {
		fileName = time.Now().Format("2006_01_02_15_04_05")
	}

	sv.fileInfo = &fileInfo{
		Field:  "sticker",
		Reader: stickerFileReader,
		Name:   fileName,
	}

	sv.sticker = "attach://sticker"

	return sv
}

func (sv *uploadStickerFile) SetStickerFormat(stickerFormat string) *uploadStickerFile {
	sv.stickerFormat = stickerFormat
	return sv
}

// SetStickerFormatStatic sets the sticker format to "static"
func (sv *uploadStickerFile) SetStickerFormatStatic() *uploadStickerFile {
	sv.stickerFormat = "static"
	return sv
}

// SetStickerFormatAnimated sets the sticker format to "animated"
func (sv *uploadStickerFile) SetStickerFormatAnimated() *uploadStickerFile {
	sv.stickerFormat = "animated"
	return sv
}

// SetStickerFormatVideo sets the sticker format to "video"
func (sv *uploadStickerFile) SetStickerFormatVideo() *uploadStickerFile {
	sv.stickerFormat = "video"
	return sv
}

/*func (sv *uploadStickerFile) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
