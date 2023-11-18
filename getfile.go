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

type getFile struct {
	parent *TelegramBot
	fileId string
}

func (sv *getFile) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FileId string `json:"file_id"`
	}{
		FileId: sv.fileId,
	})
}

func (sv *getFile) response() interface{} {
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

func (sv *getFile) getBytesReader(tgFilePath string) (io.ReadCloser, error) {
	fileAddress := fmt.Sprintf(`https://api.telegram.org/file/bot%s/%s`, sv.parent.apiToken, tgFilePath)

	resp, err := sv.parent.client.SetDebug(true).R().Get(fileAddress)
	if err != nil {
		return nil, err
	}

	return resp.RawBody(), nil
}

func (sv *getFile) method() string {
	return "POST"
}

func (sv *getFile) endpoint() string {
	return "getFile"
}

func (sv *getFile) SetFileId(fileId string) *getFile {
	sv.fileId = fileId
	return sv
}

/*func (sv *getFile) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
