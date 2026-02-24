package tgbotapi

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type inputMediaVideo struct {
	media             interface{}
	thumb             interface{}
	caption           string
	parseMode         string
	width             int64
	height            int64
	duration          int64
	supportsStreaming bool
	inputMedia
	files []fileInfo
}

func (i *inputMediaVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type              string      `json:"type"`
		Media             interface{} `json:"media"`
		Thumb             interface{} `json:"thumb,omitempty"`
		Caption           string      `json:"caption,omitempty"`
		ParseMode         string      `json:"parse_mode,omitempty"`
		Width             int64       `json:"width,omitempty"`
		Height            int64       `json:"height,omitempty"`
		Duration          int64       `json:"duration,omitempty"`
		SupportsStreaming bool        `json:"supports_streaming,omitempty"`
	}{
		Type:              "video",
		Media:             i.media,
		Thumb:             i.thumb,
		Caption:           i.caption,
		ParseMode:         i.parseMode,
		Width:             i.width,
		Height:            i.height,
		Duration:          i.duration,
		SupportsStreaming: i.supportsStreaming,
	})
}

func (i *inputMediaVideo) medias() []fileInfo {
	return i.files
}

func (i *inputMediaVideo) SetVideoId(videoId string) *inputMediaVideo {
	i.media = videoId
	return i
}

func (i *inputMediaVideo) SetVideoPath(videoPath string) *inputMediaVideo {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	if strings.HasPrefix(videoPath, "attach://") {
		i.media = videoPath
	} else {
		i.media = "attach://video"
	}
	video := &fileInfo{
		Field: "video",
		Path:  videoPath,
	}
	i.files = append(i.files, *video)
	return i
}

func (i *inputMediaVideo) SetVideoFileReader(videoFileReader io.Reader, fileName string) *inputMediaVideo {
	if fileName == "" {
		fileName = "video_" + time.Now().Format("2006_01_02_15_04_05") + ".mp4"
	}
	if strings.HasPrefix(fileName, "attach://") {
		i.media = fileName
	} else {
		i.media = "attach://video"
	}
	video := &fileInfo{
		Field:  "video",
		Reader: videoFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *video)
	return i
}

func (i *inputMediaVideo) SetThumbId(videoId string) *inputMediaVideo {
	i.thumb = videoId
	return i
}

func (i *inputMediaVideo) SetThumbPath(videoPath string) *inputMediaVideo {
	if i.files == nil {
		i.files = []fileInfo{}
	}
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field: "thumb",
		Path:  videoPath,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaVideo) SetThumbFileReader(videoFileReader io.Reader, fileName string) *inputMediaVideo {
	if fileName == "" {
		fileName = "thumb_" + time.Now().Format("2006_01_02_15_04_05") + ".png"
	}
	i.thumb = "attach://thumb"
	thumb := &fileInfo{
		Field:  "thumb",
		Reader: videoFileReader,
		Name:   fileName,
	}
	i.files = append(i.files, *thumb)
	return i
}

func (i *inputMediaVideo) SetCaption(caption string) *inputMediaVideo {
	i.caption = caption
	return i
}

func (i *inputMediaVideo) SetParseModeHTML() *inputMediaVideo {
	i.parseMode = "HTML"
	return i
}

func (i *inputMediaVideo) SetParseModeMarkdown() *inputMediaVideo {
	i.parseMode = "Markdown"
	return i
}

func (i *inputMediaVideo) SetWidth(width int64) *inputMediaVideo {
	i.width = width
	return i
}

func (i *inputMediaVideo) SetHeight(height int64) *inputMediaVideo {
	i.height = height
	return i
}

func (i *inputMediaVideo) SetDuration(duration int64) *inputMediaVideo {
	i.duration = duration
	return i
}

func (i *inputMediaVideo) SetSupportsStreaming(supportsStreaming bool) *inputMediaVideo {
	i.supportsStreaming = supportsStreaming
	return i
}
