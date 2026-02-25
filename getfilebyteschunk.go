package tgbotapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type getFileBytesChunk struct {
	parent *TelegramBot
	fileId string
	offset int64
	limit  int64
}

func (sv *getFileBytesChunk) marshalJSON() ([]byte, error) {
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

func (sv *getFileBytesChunk) Send() (io.ReadCloser, error) {
	resp, err := sv.parent.SendRaw(sv)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		defer resp.RawResponse.Body.Close()
		b, _ := io.ReadAll(resp.RawBody())
		return nil, fmt.Errorf("HTTP request failed with status code %d: %s", resp.StatusCode(), string(b))
	}

	return resp.RawBody(), nil
}

func (sv *getFileBytesChunk) method() string {
	return "POST"
}

func (sv *getFileBytesChunk) response() interface{} {
	return nil
}

// By using getFileBytes endpoint, we implicitly send offset and limit
// as our C++ custom API handles them gracefully and returns raw bytes.
func (sv *getFileBytesChunk) endpoint() string {
	return "getFileBytes"
}

func (sv *getFileBytesChunk) SetFileId(fileId string) *getFileBytesChunk {
	sv.fileId = fileId
	return sv
}

func (sv *getFileBytesChunk) SetOffset(offset int64) *getFileBytesChunk {
	sv.offset = offset
	return sv
}

func (sv *getFileBytesChunk) SetLimit(limit int64) *getFileBytesChunk {
	sv.limit = limit
	return sv
}
