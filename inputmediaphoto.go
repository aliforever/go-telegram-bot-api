package tgbotapi

import (
	"encoding/json"
	"io"
	"strings"
)

type inputMediaPhoto struct {
	media     interface{}
	caption   string
	parseMode string
	inputMedia
	photo *fileInfo
}

func (i inputMediaPhoto) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type      string      `json:"type"`
		Media     interface{} `json:"media"`
		Caption   string      `json:"caption,omitempty"`
		ParseMode string      `json:"parse_mode,omitempty"`
	}{
		Type:      "photo",
		Media:     i.media,
		Caption:   i.caption,
		ParseMode: i.parseMode,
	})
}

func (i *inputMediaPhoto) medias() []fileInfo {
	var files []fileInfo
	if i.photo != nil {
		files = append(files, *i.photo)
	}
	return files
}

func (i *inputMediaPhoto) SetPhotoId(photoId string) *inputMediaPhoto {
	i.media = photoId
	return i
}

func (i *inputMediaPhoto) SetPhotoPath(photoPath string) *inputMediaPhoto {
	i.photo = &fileInfo{
		Field: "photo",
		Path:  photoPath,
	}
	if strings.HasPrefix(photoPath, "attach://") {
		i.media = photoPath
	} else {
		i.media = "attach://photo"
	}
	return i
}

func (i *inputMediaPhoto) SetPhotoFileReader(photoFileReader io.Reader, fileName string) *inputMediaPhoto {
	if fileName == "" {
		fileName = "photo" + ".png"
	}
	i.photo = &fileInfo{
		Field:  "photo",
		Reader: photoFileReader,
		Name:   fileName,
	}
	if strings.HasPrefix(fileName, "attach://") {
		i.media = fileName
	} else {
		i.media = "attach://photo"
	}
	return i
}

func (i *inputMediaPhoto) SetCaption(caption string) *inputMediaPhoto {
	i.caption = caption
	return i
}

func (i *inputMediaPhoto) SetParseModeHTML() *inputMediaPhoto {
	i.parseMode = "HTML"
	return i
}

func (i *inputMediaPhoto) SetParseModeMarkdown() *inputMediaPhoto {
	i.parseMode = "Markdown"
	return i
}
