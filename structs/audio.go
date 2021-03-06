package structs

type Audio struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int64     `json:"duration"`
	Performer    string    `json:"performer"`
	Title        string    `json:"title"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
	Thumb        PhotoSize `json:"thumb"`
	FileName     string    `json:"file_name"`
}
