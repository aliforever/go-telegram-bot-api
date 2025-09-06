package structs

import (
	"encoding/json"
)

type InlineKeyboardMarkup struct {
	inlineKeyboard []*InlineKeyboardButtons
}

func (ikm InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		InlineKeyboard []*InlineKeyboardButtons `json:"inline_keyboard"`
	}{
		InlineKeyboard: ikm.inlineKeyboard,
	})
}

func (ikm *InlineKeyboardMarkup) UnmarshalJSON(data []byte) (err error) {
	object := struct {
		InlineKeyboard []*InlineKeyboardButtons `json:"inline_keyboard,omitempty"`
	}{}
	err = json.Unmarshal(data, &object)
	if err != nil {
		return
	}
	*ikm = InlineKeyboardMarkup{inlineKeyboard: object.InlineKeyboard}
	return
}

func (ikm *InlineKeyboardMarkup) NewRow() (row *InlineKeyboardButtons) {
	row = &InlineKeyboardButtons{}
	if ikm.inlineKeyboard == nil {
		ikm.inlineKeyboard = []*InlineKeyboardButtons{}
	}
	ikm.inlineKeyboard = append(ikm.inlineKeyboard, row)
	return row
}

func (ikm InlineKeyboardMarkup) FromSlicesOfMaps(slicesOfMap [][]map[string]string) *InlineKeyboardMarkup {
	keyboard := &InlineKeyboardMarkup{}
	for _, m := range slicesOfMap {
		keyboardRow := keyboard.NewRow()
		for _, row := range m {
			text, exists := row["text"]
			if !exists {
				continue
			}
			if url, exists := row["url"]; exists {
				keyboardRow.AddUrlButton(text, url)
				continue
			}
			if callbackData, exists := row["callback_data"]; exists {
				keyboardRow.AddCallbackButton(text, callbackData)
				continue
			}
			if webappUrl, exists := row["web_app"]; exists {
				keyboardRow.AddWebAppUrlButton(text, webappUrl)
				continue
			}
		}
	}
	return keyboard
}
