package tgbotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"

	"github.com/aliforever/go-telegram-bot-api/responses"
	"github.com/aliforever/go-telegram-bot-api/structs"
	"github.com/go-resty/resty/v2"
)

func (tb *TelegramBot) prepareRequest(config Config, request *resty.Request) error {
	if val, isFile := config.(file); isFile && len(val.medias()) != 0 {
		for _, info := range val.medias() {
			if info.Reader != nil {
				request.SetFileReader(info.Field, info.Name, info.Reader)
			} else if info.Path != "" {
				request.SetFile(info.Field, info.Path)
			}
		}
		request.SetFormData(tb.getFormData(config))
	} else {
		request.SetHeader("Content-Type", "application/json")
		body, err := config.marshalJSON()
		if err != nil {
			return err
		}

		request = request.SetBody(string(body))
	}

	return nil
}

func (tb *TelegramBot) getFormData(config Config) (fd map[string]string) {
	j, _ := config.marshalJSON()
	var result map[string]interface{}
	_ = json.Unmarshal(j, &result)
	fd = map[string]string{}
	for k, v := range result {
		switch v.(type) {
		case string:
			fd[k] = v.(string)
		case float64:
			if k != "latitude" && k != "longitude" && k != "x_shift" && k != "y_shift" && k != "scale" {
				fd[k] = fmt.Sprintf("%d", int64(v.(float64)))
			} else {
				fd[k] = fmt.Sprintf("%f", v.(float64))
			}
		case bool:
			fd[k] = strconv.FormatBool(v.(bool))
		case map[string]interface{}:
			jm, _ := json.Marshal(v.(map[string]interface{}))
			fd[k] = string(jm)
		case []interface{}:
			jm, _ := json.Marshal(v.([]interface{}))
			fd[k] = string(jm)
		case [][]map[string]string:
			jm, _ := json.Marshal(v.([][]map[string]string))
			fd[k] = string(jm)
		default:
			fmt.Println(fmt.Sprintf("Unknown type: %T", v))
		}
	}
	return
}

func (tb *TelegramBot) getMessageResponse(resp *resty.Response, config Config) (response *Response, raw []byte, err error) {
	defer resp.RawBody().Close()

	raw, err = io.ReadAll(resp.RawBody())
	if err != nil {
		return
	}

	var genericResp *genericResponse

	err = json.NewDecoder(bytes.NewReader(raw)).Decode(&genericResp)
	if err != nil {
		return
	}

	if !genericResp.Ok {
		return nil, raw, responses.Error{
			ErrorCode:   genericResp.ErrorCode,
			Description: genericResp.Description,
		}
	}

	var responseVar = config.response()

	err = json.Unmarshal(genericResp.Result, &responseVar)
	if err != nil {
		return nil, raw, err
	}

	switch responseVar.(type) {
	case *[]structs.Message:
		messages := responseVar.(*[]structs.Message)
		response.Messages = *messages
	case *[]Update:
		updates := responseVar.(*[]Update)
		response.Updates = *updates
	case *[]structs.ChatMember:
		updates := responseVar.(*[]structs.ChatMember)
		response.ChatMembers = *updates
	case *structs.ChatMember:
		response.ChatMember = responseVar.(*structs.ChatMember)
	case *structs.Message:
		response.Message = responseVar.(*structs.Message)
	case *structs.UserProfilePhotos:
		response.UserProfilePhotos = responseVar.(*structs.UserProfilePhotos)
	case *structs.File:
		response.File = responseVar.(*structs.File)
	case *structs.Chat:
		response.Chat = responseVar.(*structs.Chat)
	case *structs.User:
		response.User = responseVar.(*structs.User)
	case *structs.StickerSet:
		response.StickerSet = responseVar.(*structs.StickerSet)
	case *structs.MessageId:
		response.MessageId = responseVar.(*structs.MessageId)
	case *bool:
		response.Bool = responseVar.(*bool)
	case *int64:
		response.Int = responseVar.(*int64)
	default:
		err = errors.New(fmt.Sprintf(
			"unknown response result %s - %T - %+v", string(genericResp.Result), responseVar, config,
		))
	}
	return
}

func randomString() string {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	return s[rand.Intn(len(s)-1)]
}
