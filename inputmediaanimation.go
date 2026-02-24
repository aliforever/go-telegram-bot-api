package tgbotapi

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type inputMediaAnimation struct {
	media     interface{}
	thumb     interface{}
	caption   string
	parseMode string
	width     int64
	height    int64
	duration  int64
	inputMedia
	files []fileInfo
}

func (i inputMediaAnimation) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type      string      `json:"type"`
		Media     interface{} `json:"media"`
		Thumb     interface{} `json:"thumb,omitempty"`
		Caption   string      `json:"caption,omitempty"`
		ParseMode string      `json:"parse_mode,omitempty"`
		Width     int64       `json:"width,omitempty"`
		Height    int64       `json:"height,omitempty"`
		Duration  int64       `json:"duration,omitempty"`
	}{
		Type:      "animation",
		Media:     i.media,
		Thumb:     i.thumb,
		Caption:   i.caption,
		ParseMode: i.parseMode,
		Width:     i.width,
		Height:    i.height,
		Duration:  i.duration,
	})
}

func (i *inputMediaAnimation) medias() []fileInfo {
	return i.files
}

func (i *inputMediaAnimation) SetAnimationId(animationId string) *inputMediaAnimation {
	i.media = animationId
	return i
}

func (i *inputMediaAnimation) SetAnimationPath(animationPath string) *inputMediaAnimation {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	if strings.HasPrefix(animationPath, "attach://") {
		i.media = animationPath
	} else {
		i.media = "attach://animation"
	}
	animation := &fileInfo{
		Field: "animation",
		Path:  animationPath,
	}
	i.files = append(i.files, *animation)
	return i
}

func (i *inputMediaAnimation) SetAnimationFileReader(animationFileReader io.Reader, fileName string) *inputMediaAnimation {
	if fileName == "" {
		fileName = "animation_" + time.Now().Format("2006_01_02_15_04_05")
	}
	if strings.HasPrefix(fileName, "attach://") {
		i.media = fileName
	} else {
		i.media = "attach://" + fileName
	}
	animation := &fileInfo{
		Field:  "animation",
		Reader: animationFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *animation)
	return i
}

func (i *inputMediaAnimation) SetThumbId(animationId string) *inputMediaAnimation {
	i.thumb = animationId
	return i
}

func (i *inputMediaAnimation) SetThumbPath(animationPath string) *inputMediaAnimation {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field: "thumb",
		Path:  animationPath,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaAnimation) SetThumbFileReader(thumbFileReader io.Reader, fileName string) *inputMediaAnimation {
	if fileName == "" {
		fileName = "thumb_" + time.Now().Format("2006_01_02_15_04_05")
	}
	i.thumb = "attach://" + fileName
	thumb := &fileInfo{
		Field:  "thumb",
		Reader: thumbFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaAnimation) SetCaption(caption string) *inputMediaAnimation {
	i.caption = caption
	return i
}

func (i *inputMediaAnimation) SetParseModeHTML() *inputMediaAnimation {
	i.parseMode = "HTML"
	return i
}

func (i *inputMediaAnimation) SetParseModeMarkdown() *inputMediaAnimation {
	i.parseMode = "Markdown"
	return i
}

func (i *inputMediaAnimation) SetDuration(duration int64) *inputMediaAnimation {
	i.duration = duration
	return i
}
