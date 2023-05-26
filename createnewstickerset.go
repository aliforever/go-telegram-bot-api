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

	userId          int64
	name            string
	title           string
	stickers        []inputSticker
	format          string
	kind            string
	needsRepainting bool

	fileInfo []fileInfo
}

func (sv *createNewStickerSet) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId          int64          `json:"user_id"`
		Name            string         `json:"name"`
		Title           string         `json:"title"`
		Stickers        []inputSticker `json:"stickers"`
		StickerFormat   string         `json:"sticker_format"`
		StickerKind     string         `json:"sticker_type"`
		NeedsRepainting bool           `json:"needs_repainting"`
	}{
		UserId:          sv.userId,
		Name:            sv.name,
		Title:           sv.title,
		Stickers:        sv.stickers,
		StickerFormat:   sv.format,
		StickerKind:     sv.kind,
		NeedsRepainting: sv.needsRepainting,
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

func (sv *createNewStickerSet) SetFormat(format string) *createNewStickerSet {
	sv.format = format
	return sv
}

func (sv *createNewStickerSet) SetFormatStatic() *createNewStickerSet {
	sv.format = "static"
	return sv
}

func (sv *createNewStickerSet) SetFormatAnimated() *createNewStickerSet {
	sv.format = "animated"
	return sv
}

func (sv *createNewStickerSet) SetFormatVideo() *createNewStickerSet {
	sv.format = "video"
	return sv
}

func (sv *createNewStickerSet) SetType(kind string) *createNewStickerSet {
	sv.kind = kind
	return sv
}

func (sv *createNewStickerSet) SetTypeRegular() *createNewStickerSet {
	sv.kind = "regular"
	return sv
}

func (sv *createNewStickerSet) SetTypeMask() *createNewStickerSet {
	sv.kind = "mask"
	return sv
}

func (sv *createNewStickerSet) SetTypeCustomEmoji() *createNewStickerSet {
	sv.kind = "custom_emoji"
	return sv
}

func (sv *createNewStickerSet) SetNeedsRepainting(needsRepainting bool) *createNewStickerSet {
	sv.needsRepainting = needsRepainting
	return sv
}

func (sv *createNewStickerSet) NeedsRepainting() *createNewStickerSet {
	sv.needsRepainting = true
	return sv
}

func (sv *createNewStickerSet) DoesntNeedRepainting() *createNewStickerSet {
	sv.needsRepainting = false
	return sv
}
