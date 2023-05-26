package tgbotapi

import (
	"encoding/json"
	"fmt"
	"github.com/aliforever/go-telegram-bot-api/structs"
	"io"
	"time"
)

type createNewStickerSet struct {
	parent *TelegramBot

	userId   int64
	name     string
	title    string
	stickers []inputSticker

	fileInfo []fileInfo
}

func (sv *createNewStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId   int64          `json:"user_id"`
		Name     string         `json:"name"`
		Title    string         `json:"title"`
		Stickers []inputSticker `json:"stickers"`
	}{
		UserId:   sv.userId,
		Name:     sv.name,
		Title:    sv.title,
		Stickers: sv.stickers,
	})
}

func (sv *createNewStickerSet) response() interface{} {
	var resp bool
	return &resp
}

func (sv *createNewStickerSet) method() string {
	return "POST"
}

func (sv *createNewStickerSet) endpoint() string {
	return "createNewStickerSet"
}

func (sv *createNewStickerSet) medias() []fileInfo {
	return sv.fileInfo
}

func (sv *createNewStickerSet) SetUserId(userId int64) *createNewStickerSet {
	sv.userId = userId
	return sv
}

func (sv *createNewStickerSet) SetName(name string) *createNewStickerSet {
	sv.name = name
	return sv
}

func (sv *createNewStickerSet) SetTitle(title string) *createNewStickerSet {
	sv.title = title
	return sv
}

func (sv *createNewStickerSet) AddStickerWithFileId(
	fileID string,
	emojies []string,
	maskPosition *structs.MaskPosition,
	keywords []string,
) *createNewStickerSet {

	sv.stickers = append(sv.stickers, inputSticker{
		Sticker:      fileID,
		EmojiList:    emojies,
		MaskPosition: maskPosition,
		Keywords:     keywords,
	})

	return sv
}

func (sv *createNewStickerSet) AddStickerWithFilePath(
	filePath string,
	emojies []string,
	maskPosition *structs.MaskPosition,
	keywords []string,
) *createNewStickerSet {

	fieldName := "photo_" + time.Now().Format("2006_01_02_15_04_05") + randomString()

	sv.stickers = append(sv.stickers, inputSticker{
		Sticker:      fmt.Sprintf("attach://%s", fieldName),
		EmojiList:    emojies,
		MaskPosition: maskPosition,
		Keywords:     keywords,
	})

	sv.fileInfo = append(sv.fileInfo, fileInfo{
		Field: fieldName,
		Path:  filePath,
	})

	return sv
}

func (sv *createNewStickerSet) AddStickerWithFileReader(
	filename string,
	fileReader io.Reader,
	emojies []string,
	maskPosition *structs.MaskPosition,
	keywords []string,
) *createNewStickerSet {

	fieldName := "photo_" + time.Now().Format("2006_01_02_15_04_05") + randomString()

	sv.stickers = append(sv.stickers, inputSticker{
		Sticker:      fmt.Sprintf("attach://%s", fieldName),
		EmojiList:    emojies,
		MaskPosition: maskPosition,
		Keywords:     keywords,
	})

	if filename == "" {
		filename = time.Now().Format("2006_01_02_15_04_05")
	}

	sv.fileInfo = append(sv.fileInfo, fileInfo{
		Field:  fieldName,
		Reader: fileReader,
		Name:   filename,
	})

	return sv
}
