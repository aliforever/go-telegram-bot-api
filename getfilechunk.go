package tgbotapi

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type getFileChunk struct {
	parent *TelegramBot
	fileId string
	offset int64
	limit  int64
}

func (sv *getFileChunk) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FileId string `json:"file_id"`
		Offset int64  `json:"offset,omitempty"`
		Limit  int64  `json:"limit,omitempty"`
	}{
		FileId: sv.fileId,
		Offset: sv.offset,
		Limit:  sv.limit,
	})
}

func (sv *getFileChunk) response() interface{} {
	f := &structs.File{}

	f.DownloadBytes = func() ([]byte, error) {
		resp, err := sv.getBytesReader(f.FilePath)
		if err != nil {
			return nil, err
		}

		defer resp.Close()

		b, err := io.ReadAll(resp)
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	f.Download = func(path string) error {
		resp, err := sv.getBytesReader(f.FilePath)
		if err != nil {
			return err
		}

		defer resp.Close()

		if path == "" {
			path = f.FilePath

			fPath := filepath.Base(path)

			err = os.MkdirAll(strings.ReplaceAll(path, fPath, ""), os.ModePerm)
			if err != nil {
				return err
			}
		}

		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}

		defer f.Close()

		_, err = io.Copy(f, resp)
		if err != nil {
			return err
		}

		return nil
	}

	return f
}

func (sv *getFileChunk) getBytesReader(tgFilePath string) (io.ReadCloser, error) {
	fileAddress := fmt.Sprintf(`%sfile/bot%s/%s`, sv.parent.apiUrl, sv.parent.apiToken, tgFilePath)

	resp, err := sv.parent.client.SetDebug(true).R().Get(fileAddress)
	if err != nil {
		return nil, err
	}

	return resp.RawBody(), nil
}

func (sv *getFileChunk) method() string {
	return "POST"
}

// By using getFile endpoint, we implicitly send offset and limit
// as our C++ custom API handles them gracefully.
func (sv *getFileChunk) endpoint() string {
	return "getFile"
}

func (sv *getFileChunk) SetFileId(fileId string) *getFileChunk {
	sv.fileId = fileId
	return sv
}

func (sv *getFileChunk) SetOffset(offset int64) *getFileChunk {
	sv.offset = offset
	return sv
}

func (sv *getFileChunk) SetLimit(limit int64) *getFileChunk {
	sv.limit = limit
	return sv
}
