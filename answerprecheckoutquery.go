package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type answerPreCheckoutQuery struct {
	parent *TelegramBot

	preCheckoutQueryId string
	ok                 bool
	errorMessage       *string
}

func (apcq *answerPreCheckoutQuery) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PreCheckoutQueryId string  `json:"pre_checkout_query_id"`
		Ok                 bool    `json:"ok"`
		ErrorMessage       *string `json:"error_message,omitempty"`
	}{
		PreCheckoutQueryId: apcq.preCheckoutQueryId,
		Ok:                 apcq.ok,
		ErrorMessage:       apcq.errorMessage,
	})
}

func (apcq *answerPreCheckoutQuery) response() interface{} {
	return structs.ResponseTypeBool()
}

func (apcq *answerPreCheckoutQuery) method() string {
	return "POST"
}

func (apcq *answerPreCheckoutQuery) endpoint() string {
	return "answerPreCheckoutQuery"
}

func (apcq *answerPreCheckoutQuery) SetPreCheckoutQueryId(preCheckoutQueryId string) *answerPreCheckoutQuery {
	apcq.preCheckoutQueryId = preCheckoutQueryId
	return apcq
}

func (apcq *answerPreCheckoutQuery) SetOk(ok bool) *answerPreCheckoutQuery {
	apcq.ok = ok
	return apcq
}

func (apcq *answerPreCheckoutQuery) SetErrorMessage(errorMessage string) *answerPreCheckoutQuery {
	apcq.errorMessage = &errorMessage
	return apcq
}
