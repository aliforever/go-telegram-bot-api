package structs

type StickerSet struct {
	Name          string      `json:"name"`
	Title         string      `json:"title"`
	StickerType   string      `json:"sticker_type"`
	IsVideo       bool        `json:"is_video"`
	IsAnimated    bool        `json:"is_animated"`
	ContainsMasks bool        `json:"contains_masks"`
	Stickers      []Sticker   `json:"stickers"`
	Thumbnail     []PhotoSize `json:"thumbnail"`
}
