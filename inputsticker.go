package tgbotapi

import "github.com/aliforever/go-telegram-bot-api/structs"

type inputSticker struct {
	Sticker      interface{}           `json:"sticker"`
	Format       string                `json:"format"`
	EmojiList    []string              `json:"emoji_list"`
	MaskPosition *structs.MaskPosition `json:"mask_position"`
	Keywords     []string              `json:"keywords"`
}
