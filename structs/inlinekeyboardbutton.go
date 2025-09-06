package structs

import "encoding/json"

type InlineKeyboardButton struct {
	text                         string
	url                          string
	loginUrl                     LoginUrl
	webApp                       *WebApp
	callbackData                 string
	switchInlineQuery            string
	switchInlineQueryCurrentChat string
	callbackGame                 interface{}
	pay                          bool
}

func (i InlineKeyboardButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Text                         string      `json:"text"`
		Url                          string      `json:"url,omitempty"`
		LoginUrl                     LoginUrl    `json:"login_url,omitempty"`
		CallbackData                 string      `json:"callback_data,omitempty"`
		SwitchInlineQuery            string      `json:"switch_inline_query,omitempty"`
		SwitchInlineQueryCurrentChat string      `json:"switch_inline_query_current_chat,omitempty"`
		CallbackGame                 interface{} `json:"callback_game,omitempty"`
		Pay                          bool        `json:"pay,omitempty"`
		WebApp                       *WebApp     `json:"web_app,omitempty"`
	}{
		Text:                         i.text,
		Url:                          i.url,
		LoginUrl:                     i.loginUrl,
		CallbackData:                 i.callbackData,
		SwitchInlineQuery:            i.switchInlineQuery,
		SwitchInlineQueryCurrentChat: i.switchInlineQueryCurrentChat,
		CallbackGame:                 i.callbackGame,
		Pay:                          i.pay,
		WebApp:                       i.webApp,
	})
}

func (i *InlineKeyboardButton) UnmarshalJSON(data []byte) (err error) {
	object := struct {
		Text                         string      `json:"text"`
		Url                          string      `json:"url"`
		LoginUrl                     LoginUrl    `json:"login_url"`
		WebApp                       *WebApp     `json:"web_app"`
		CallbackData                 string      `json:"callback_data"`
		SwitchInlineQuery            string      `json:"switch_inline_query"`
		SwitchInlineQueryCurrentChat string      `json:"switch_inline_query_current_chat"`
		CallbackGame                 interface{} `json:"callback_game"`
		Pay                          bool        `json:"pay"`
	}{}
	err = json.Unmarshal(data, &object)
	if err == nil {
		button := InlineKeyboardButton{
			text:                         object.Text,
			url:                          object.Url,
			loginUrl:                     object.LoginUrl,
			webApp:                       object.WebApp,
			callbackData:                 object.CallbackData,
			switchInlineQuery:            object.SwitchInlineQuery,
			switchInlineQueryCurrentChat: object.SwitchInlineQueryCurrentChat,
			callbackGame:                 object.CallbackGame,
			pay:                          object.Pay,
		}
		*i = button
	}
	return
}

func (i *InlineKeyboardButton) SetText(text string) {
	(*i).text = text
}

func (i *InlineKeyboardButton) SetWebApp(webApp *WebApp) {
	(*i).webApp = webApp
}

func (i *InlineKeyboardButton) SetWebAppURL(url string) {
	(*i).webApp = &WebApp{URL: url}
}

//
// func (i *InlineKeyboardButton) SetUrl(url string) {
// 	i.url = url
// }
//
// func (i *InlineKeyboardButton) SetLoginUrl(loginUrl LoginUrl) {
// 	i.loginUrl = loginUrl
// }
//
// func (i *InlineKeyboardButton) SetCallbackData(callbackData string) {
// 	i.callbackData = callbackData
// }
//
// func (i *InlineKeyboardButton) SetSwitchInlineQuery(switchInlineQuery string) {
// 	i.switchInlineQuery = switchInlineQuery
// }
//
// func (i *InlineKeyboardButton) SwitchInlineQueryCurrentChat() string {
// 	return i.switchInlineQueryCurrentChat
// }
//
// func (i *InlineKeyboardButton) SetCallbackGame(callbackGame interface{}) {
// 	i.callbackGame = callbackGame
// }
//
// func (i *InlineKeyboardButton) EnablePay() {
// 	i.pay = true
// }
//
// func (i *InlineKeyboardButton) DisablePay() {
// 	i.pay = false
// }
