package structs

type AlternativeVideo struct {
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	Codec        string `json:"codec"`
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
}

type Video struct {
	FileId            string             `json:"file_id"`
	FileUniqueId      string             `json:"file_unique_id"`
	Width             int64              `json:"width"`
	Height            int64              `json:"height"`
	Duration          int64              `json:"duration"`
	Thumb             PhotoSize          `json:"thumb"`
	MimeType          string             `json:"mime_type"`
	FileSize          int64              `json:"file_size"`
	FileName          string             `json:"file_name"`
	AlternativeVideos []AlternativeVideo `json:"alternative_videos"`
}
