package structs

import "fmt"

type File struct {
	FileId       string             `json:"file_id"`
	FileUniqueId string             `json:"file_unique_id"`
	FileSize     int64              `json:"file_size"`
	FilePath     string             `json:"file_path"`
	Download     func(string) error `json:"-"`
}

// DownloadPath returns the path to download the file to.
func (f *File) DownloadPath(botToken string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", botToken, f.FilePath)
}
