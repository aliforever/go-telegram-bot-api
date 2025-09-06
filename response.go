package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type genericResponse struct {
	Ok          bool            `json:"ok"`
	ErrorCode   int64           `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
	Result      json.RawMessage `json:"result"`
}

type Response struct {
	Bool              *bool
	Int               *int64
	Message           *structs.Message
	Messages          []structs.Message
	UserProfilePhotos *structs.UserProfilePhotos
	Updates           []Update
	ChatMembers       []structs.ChatMember
	ChatMember        *structs.ChatMember
	Chat              *structs.Chat
	User              *structs.User
	File              *structs.File
	StickerSet        *structs.StickerSet
	MessageId         *structs.MessageId
	WebAppData        *structs.WebAppData
	SentWebAppMessage *structs.SentWebAppMessage
	Raw               []byte
}
