package tgbotapi

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type inputMediaAudio struct {
	media     interface{}
	thumb     interface{}
	caption   string
	parseMode string
	duration  int64
	performer string
	title     string
	inputMedia
	files []fileInfo
}

func (i inputMediaAudio) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type              string      `json:"type"`
		Media             interface{} `json:"media"`
		Thumb             interface{} `json:"thumb,omitempty"`
		Caption           string      `json:"caption,omitempty"`
		ParseMode         string      `json:"parse_mode,omitempty"`
		Duration          int64       `json:"duration,omitempty"`
		Performer         string      `json:"performer,omitempty"`
		Title             string      `json:"title,omitempty"`
		SupportsStreaming bool        `json:"supports_streaming,omitempty"`
	}{
		Type:      "audio",
		Media:     i.media,
		Thumb:     i.thumb,
		Caption:   i.caption,
		ParseMode: i.parseMode,
		Duration:  i.duration,
		Performer: i.performer,
		Title:     i.title,
	})
}

func (i *inputMediaAudio) medias() []fileInfo {
	return i.files
}

func (i *inputMediaAudio) SetAudioId(audioId string) *inputMediaAudio {
	i.media = audioId
	return i
}

func (i *inputMediaAudio) SetAudioPath(audioPath string) *inputMediaAudio {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	if strings.HasPrefix(audioPath, "attach://") {
		i.media = audioPath
	} else {
		i.media = "attach://audio"
	}
	audio := &fileInfo{
		Field: "audio",
		Path:  audioPath,
	}
	i.files = append(i.files, *audio)
	return i
}

func (i *inputMediaAudio) SetAudioFileReader(audioFileReader io.Reader, fileName string) *inputMediaAudio {
	if fileName == "" {
		fileName = "audio_" + time.Now().Format("2006_01_02_15_04_05") + ".mp3"
	}
	if strings.HasPrefix(fileName, "attach://") {
		i.media = fileName
	} else {
		i.media = "attach://audio"
	}
	audio := &fileInfo{
		Field:  "audio",
		Reader: audioFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *audio)
	return i
}

func (i *inputMediaAudio) SetThumbId(audioId string) *inputMediaAudio {
	i.thumb = audioId
	return i
}

func (i *inputMediaAudio) SetThumbPath(audioPath string) *inputMediaAudio {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	fieldName := "thumb_" + time.Now().Format("2006_01_02_15_04_05")
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field: fieldName,
		Path:  audioPath,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaAudio) SetThumbFileReader(audioFileReader io.Reader, fileName string) *inputMediaAudio {
	if fileName == "" {
		fileName = "thumb_" + time.Now().Format("2006_01_02_15_04_05") + ".png"
	}
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field:  "thumb",
		Reader: audioFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaAudio) SetCaption(caption string) *inputMediaAudio {
	i.caption = caption
	return i
}

func (i *inputMediaAudio) SetParseModeHTML() *inputMediaAudio {
	i.parseMode = "HTML"
	return i
}

func (i *inputMediaAudio) SetParseModeMarkdown() *inputMediaAudio {
	i.parseMode = "Markdown"
	return i
}

func (i *inputMediaAudio) SetPerformer(performer string) *inputMediaAudio {
	i.performer = performer
	return i
}

func (i *inputMediaAudio) SetTitle(title string) *inputMediaAudio {
	i.title = title
	return i
}

func (i *inputMediaAudio) SetDuration(duration int64) *inputMediaAudio {
	i.duration = duration
	return i
}
