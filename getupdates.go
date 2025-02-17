package tgbotapi

import (
	"encoding/json"
	"log/slog"
)

type getUpdates struct {
	parent         *TelegramBot
	offset         int64
	limit          int64
	timeout        int64
	allowedUpdates []string
	onError        func(error) bool
	logger         *slog.Logger
	logUpdates     bool
}

func (gu *getUpdates) Offset() int64 {
	return gu.offset
}

func (gu *getUpdates) Limit() int64 {
	return gu.limit
}

func (gu *getUpdates) Timeout() int64 {
	return gu.timeout
}

func (gu *getUpdates) AllowedUpdates() []string {
	return gu.allowedUpdates
}

func (gu *getUpdates) SetOnError(onError func(error) bool) *getUpdates {
	gu.onError = onError

	return gu
}

func (gu *getUpdates) SetLogger(logger *slog.Logger) *getUpdates {
	gu.logger = logger
	return gu
}

func (gu *getUpdates) LogUpdates() *getUpdates {
	gu.logUpdates = true
	return gu
}

func (gu *getUpdates) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Offset         int64    `json:"offset,omitempty"`
		Limit          int64    `json:"limit,omitempty"`
		Timeout        int64    `json:"timeout,omitempty"`
		AllowedUpdates []string `json:"allowed_updates,omitempty"`
	}{
		Offset:         gu.offset,
		Limit:          gu.limit,
		Timeout:        gu.timeout,
		AllowedUpdates: gu.allowedUpdates,
	})
}

func (gu *getUpdates) response() interface{} {
	return &[]Update{}
}

func (gu *getUpdates) method() string {
	return "POST"
}

func (gu *getUpdates) endpoint() string {
	return "getUpdates"
}

func (gu *getUpdates) SetOffset(offset int64) *getUpdates {
	gu.offset = offset
	return gu
}

func (gu *getUpdates) SetLimit(limit int64) *getUpdates {
	gu.limit = limit
	return gu
}

func (gu *getUpdates) SetTimeout(timeout int64) *getUpdates {
	gu.timeout = timeout
	return gu
}

func (gu *getUpdates) SetAllowedUpdates(allowed []string) *getUpdates {
	gu.allowedUpdates = allowed
	return gu
}

func (gu *getUpdates) Get() (result *Response, err error) {
	result, err = gu.parent.Send(gu)
	return
}

func (gu *getUpdates) LongPoll() error {
	return gu.pollUpdates()
}

func (gu *getUpdates) handleErrorAndReturn(err error) *bool {
	if err == nil {
		return nil
	}

	if gu.onError != nil {
		result := gu.onError(err)
		return &result
	}

	shouldReturn := true

	return &shouldReturn
}

func (gu *getUpdates) pollUpdates() error {
	var update Update

	for {
		resp, err := gu.parent.Send(gu)
		if shouldReturn := gu.handleErrorAndReturn(err); shouldReturn != nil {
			if *shouldReturn {
				gu.parent.updates <- newUpdateError(err)
				return err
			}

			continue
		}

		for _, update = range resp.Updates {
			gu.parent.updates <- update
		}

		gu.SetOffset(update.UpdateId + 1)
	}
}
