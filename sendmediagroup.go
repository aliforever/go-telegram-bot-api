package tgbotapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendMediaGroup struct {
	parent              *TelegramBot
	chatId              interface{}
	media               []interface{}
	disableNotification bool
	replyToMessageId    int64
	// DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	file
	files []fileInfo
}

func (sd *sendMediaGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId              interface{}   `json:"chat_id"`
		Media               []interface{} `json:"media"`
		DisableNotification bool          `json:"disable_notification"`
		ReplyToMessageId    int64         `json:"reply_to_message_id"`
	}{
		ChatId:              sd.chatId,
		Media:               sd.media,
		DisableNotification: sd.disableNotification,
		ReplyToMessageId:    sd.replyToMessageId,
	})
}

func (sd *sendMediaGroup) response() interface{} {
	return &[]structs.Message{}
}

func (sd *sendMediaGroup) method() string {
	return "POST"
}

func (sd *sendMediaGroup) endpoint() string {
	return "sendMediaGroup"
}

func (sd *sendMediaGroup) medias() []fileInfo {
	return sd.files
}

func (sd *sendMediaGroup) SetChatId(chatId int64) *sendMediaGroup {
	sd.chatId = chatId
	return sd
}

func (sd *sendMediaGroup) SetChatUsername(username string) *sendMediaGroup {
	sd.chatId = username
	return sd
}

func (sd *sendMediaGroup) DisableNotification() *sendMediaGroup {
	sd.disableNotification = true
	return sd
}

func (sd *sendMediaGroup) EnableNotification() *sendMediaGroup {
	sd.disableNotification = false
	return sd
}

/*func (sd *sendMediaGroup) SetMedia(media []interface{}) *sendMediaGroup {
	sd.media = media
	return sd
}*/

func (sd *sendMediaGroup) AddPhotoId(id, caption string) *sendMediaGroup {
	i := (&inputMediaPhoto{}).SetCaption(caption).SetPhotoId(id)

	if sd.media == nil {
		sd.media = []interface{}{}
	}
	sd.media = append(sd.media, i)
	return sd
}

func (sd *sendMediaGroup) AddPhotoFilePath(path, caption string) *sendMediaGroup {
	i := &inputMediaPhoto{}
	i.SetCaption(caption)
	if sd.media == nil {
		sd.files = []fileInfo{}
	}
	fieldName := "photo_" + time.Now().Format("2006_01_02_15_04_05") + randomString()
	i.SetPhotoPath("attach://" + fieldName)
	sd.files = append(sd.files, fileInfo{
		Field: fieldName,
		Path:  path,
	})

	if sd.media == nil {
		sd.media = []interface{}{}
	}
	sd.media = append(sd.media, i)
	return sd
}

// func (sd *sendMediaGroup) AddPhotoFileReader(photoFileReader io.Reader, fileName, caption string) *sendMediaGroup {
// 	i := structs.inputMediaPhoto{}
// 	i.SetCaption(caption)
// 	if sd.pngFile == nil {
// 		sd.pngFile = []fileInfo{}
// 	}
// 	fieldName := "photo_" + time.Now().Format("2006_01_02_15_04_05") + randomString()
// 	i.SetPhotoFileReader("attach://" + fieldName)
// 	sd.pngFile = append(sd.pngFile, fileInfo{
// 		Field:  fieldName,
// 		Reader: photoFileReader,
// 		Name:   fileName,
// 	})
//
// 	if sd.media == nil {
// 		sd.media = []interface{}{}
// 	}
// 	sd.media = append(sd.media, i)
// 	return sd
// }

/*func (sd *sendMediaGroup) AddVideoId(id, caption string) *sendMediaGroup {
	i := (&inputMediaVideo{}).SetCaption(caption).SetMedia(id)

	if sd.media == nil {
		sd.media = []interface{}{}
	}
	sd.media = append(sd.media, i)
	return sd
}
*/
/*func (sd *sendMediaGroup) AddVideoFilePath(path, caption string) *sendMediaGroup {
	i := inputMediaVideo{}
	i.SetCaption(caption)
	if sd.media == nil {
		sd.pngFile = []fileInfo{}
	}
	fieldName := "video_" + time.Now().Format("2006_01_02_15_04_05") + randomString()
	i.SetMedia("attach://" + fieldName)
	sd.pngFile = append(sd.pngFile, fileInfo{
		Field: fieldName,
		Path:  path,
	})

	if sd.media == nil {
		sd.media = []interface{}{}
	}
	sd.media = append(sd.media, i)
	return sd
}

func (sd *sendMediaGroup) AddVideoFileReader(videoFileReader io.Reader, fileName, caption string) *sendMediaGroup {
	i := inputMediaVideo{}
	i.SetCaption(caption)
	if sd.pngFile == nil {
		sd.pngFile = []fileInfo{}
	}
	fieldName := "video_" + time.Now().Format("2006_01_02_15_04_05") + randomString()
	i.SetMedia("attach://" + fieldName)
	sd.pngFile = append(sd.pngFile, fileInfo{
		Field:  fieldName,
		Reader: videoFileReader,
		Name:   fileName,
	})

	if sd.media == nil {
		sd.media = []interface{}{}
	}
	sd.media = append(sd.media, i)
	return sd
}
*/
func (sd *sendMediaGroup) SetReplyToMessageId(replyToMessageId int64) {
	sd.replyToMessageId = replyToMessageId
}

func (sd *sendMediaGroup) Send() (messages []structs.Message, err error) {
	request := sd.parent.client.R()
	if len(sd.files) != 0 {
		for _, info := range sd.files {
			if info.Reader != nil {
				request.SetFileReader(info.Field, info.Name, info.Reader)
			} else if info.Path != "" {
				request.SetFile(info.Field, info.Path)
			}
		}
		data := map[string]string{}
		if chatId, isStr := sd.chatId.(string); isStr {
			data["chat_id"] = chatId
		} else if chatId, isInt := sd.chatId.(int64); isInt {
			data["chat_id"] = fmt.Sprintf("%d", chatId)
		}
		mediaByte, _ := json.Marshal(sd.media)
		data["media"] = string(mediaByte)
		data["disable_notification"] = strconv.FormatBool(sd.disableNotification)
		data["reply_to_message_id"] = fmt.Sprintf("%d", sd.replyToMessageId)
		request.SetFormData(data)
	} else {
		request.SetHeader("Content-Type", "application/json")
		body, _ := json.Marshal(struct {
			ChatId              interface{}   `json:"chat_id"`
			Media               []interface{} `json:"media"`
			DisableNotification bool          `json:"disable_notification"`
			ReplyToMessageId    int64         `json:"reply_to_message_id"`
		}{
			ChatId:              sd.chatId,
			Media:               sd.media,
			DisableNotification: sd.disableNotification,
			ReplyToMessageId:    sd.replyToMessageId,
		})
		request = request.SetBody(string(body))
	}
	resp, err := request.Post(sd.endpoint())
	if err != nil {
		return nil, err
	}
	result, _, err := sd.parent.getMessageResponse(resp, sd)
	if err != nil {
		return
	}
	messages = result.Messages
	return
}
