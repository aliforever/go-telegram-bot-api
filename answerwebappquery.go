package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type answerWebAppQuery struct {
	parent        *TelegramBot
	webAppQueryId string
	result        interface{}
}

func (a *answerWebAppQuery) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		WebAppQueryId string      `json:"web_app_query_id"`
		Result        interface{} `json:"result"`
	}{
		WebAppQueryId: a.webAppQueryId,
		Result:        a.result,
	})
}

func (a *answerWebAppQuery) response() interface{} {
	return &structs.SentWebAppMessage{}
}

func (a *answerWebAppQuery) method() string {
	return "POST"
}

func (a *answerWebAppQuery) endpoint() string {
	return "answerWebAppQuery"
}

func (a *answerWebAppQuery) SetWebAppQueryId(webAppQueryId string) *answerWebAppQuery {
	a.webAppQueryId = webAppQueryId
	return a
}

func (a *answerWebAppQuery) SetResult(result interface{}) *answerWebAppQuery {
	a.result = result
	return a
}
