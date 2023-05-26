package structs

type Sticker struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int64     `json:"width"`
	Height       int64     `json:"height"`
	Duration     int64     `json:"duration"`
	Thumb        PhotoSize `json:"thumb"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
	SetName      string    `json:"set_name"`
	IsAnimated   bool      `json:"is_animated"`
	IsVideo      bool      `json:"is_video"`
}
