package structs

type InlineKeyboardButtons []*InlineKeyboardButton

func (ikb *InlineKeyboardButtons) AddUrlButton(buttonText, url string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, url: url}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddLoginUrlButton(buttonText string, loginUrl LoginUrl) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, loginUrl: loginUrl}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddCallbackButton(buttonText, callback string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, callbackData: callback}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddSwitchInlineQueryButton(buttonText, switchInlineQuery string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, callbackData: switchInlineQuery}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddSwitchInlineQueryCurrentChatButton(buttonText, switchInlineQueryCurrentChat string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, callbackData: switchInlineQueryCurrentChat}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddCallbackGameButton(buttonText string, callbackGame interface{}) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, callbackGame: callbackGame}
	ikb.addToFirst(button)
	return
}

func (ikb *InlineKeyboardButtons) AddPayButton(buttonText string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, pay: true}
	ikb.addToFirst(button)
	return
}

func (ikb *InlineKeyboardButtons) AddWebAppButton(buttonText string, webApp *WebApp) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, webApp: webApp}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) AddWebAppUrlButton(buttonText, url string) (button *InlineKeyboardButton) {
	button = &InlineKeyboardButton{text: buttonText, webApp: &WebApp{URL: url}}
	*ikb = append(*ikb, button)
	return
}

func (ikb *InlineKeyboardButtons) addToFirst(button *InlineKeyboardButton) {
	if len(*ikb) == 0 {
		*ikb = append(*ikb, button)
	} else {
		nIkb := &InlineKeyboardButtons{}
		*nIkb = append(*nIkb, button)
		for _, keyboardButton := range *ikb {
			*nIkb = append(*nIkb, keyboardButton)
		}
		*ikb = *nIkb
	}
}
