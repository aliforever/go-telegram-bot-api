package tgbotapi

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type inputMediaDocument struct {
	media     interface{}
	thumb     interface{}
	caption   string
	parseMode string
	inputMedia
	files []fileInfo
}

func (i inputMediaDocument) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type      string      `json:"type"`
		Media     interface{} `json:"media"`
		Thumb     interface{} `json:"thumb,omitempty"`
		Caption   string      `json:"caption,omitempty"`
		ParseMode string      `json:"parse_mode,omitempty"`
	}{
		Type:      "document",
		Media:     i.media,
		Thumb:     i.thumb,
		Caption:   i.caption,
		ParseMode: i.parseMode,
	})
}

func (i *inputMediaDocument) medias() []fileInfo {
	return i.files
}

func (i *inputMediaDocument) SetDocumentId(documentId string) *inputMediaDocument {
	i.media = documentId
	return i
}

func (i *inputMediaDocument) SetDocumentPath(documentPath string) *inputMediaDocument {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	if strings.HasPrefix(documentPath, "attach://") {
		i.media = documentPath
	} else {
		i.media = "attach://document"
	}
	document := &fileInfo{
		Field: "document",
		Path:  documentPath,
	}
	i.files = append(i.files, *document)
	return i
}

func (i *inputMediaDocument) SetDocumentFileReader(documentFileReader io.Reader, fileName string) *inputMediaDocument {
	if fileName == "" {
		fileName = "document_" + time.Now().Format("2006_01_02_15_04_05")
	}
	if strings.HasPrefix(fileName, "attach://") {
		i.media = fileName
	} else {
		i.media = "attach://" + fileName
	}
	document := &fileInfo{
		Field:  "document",
		Reader: documentFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *document)
	return i
}

func (i *inputMediaDocument) SetThumbId(documentId string) *inputMediaDocument {
	i.thumb = documentId
	return i
}

func (i *inputMediaDocument) SetThumbPath(documentPath string) *inputMediaDocument {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field: "thumb",
		Path:  documentPath,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaDocument) SetThumbFileReader(thumbFileReader io.Reader, fileName string) *inputMediaDocument {
	if fileName == "" {
		fileName = time.Now().Format("2006_01_02_15_04_05") + ".png"
	}
	i.thumb = "attach://thumb" + fileName
	thumb := &fileInfo{
		Field:  "thumb",
		Reader: thumbFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaDocument) SetCaption(caption string) *inputMediaDocument {
	i.caption = caption
	return i
}

func (i *inputMediaDocument) SetParseModeHTML() *inputMediaDocument {
	i.parseMode = "HTML"
	return i
}

func (i *inputMediaDocument) SetParseModeMarkdown() *inputMediaDocument {
	i.parseMode = "Markdown"
	return i
}
