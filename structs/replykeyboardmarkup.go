package structs

import "encoding/json"

type ReplyKeyboardMarkup struct {
	keyboard        [][]KeyboardButton
	resizeKeyboard  bool
	oneTimeKeyboard bool
	selective       bool
}

func (rkm ReplyKeyboardMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Keyboard        [][]KeyboardButton `json:"keyboard"`
		ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
		Selective       bool               `json:"selective,omitempty"`
	}{
		Keyboard:        rkm.keyboard,
		ResizeKeyboard:  rkm.resizeKeyboard,
		OneTimeKeyboard: rkm.oneTimeKeyboard,
		Selective:       rkm.selective,
	})
}

func (rkm ReplyKeyboardMarkup) WebAppButton(url string) *ReplyKeyboardMarkup {
	return &ReplyKeyboardMarkup{
		keyboard: [][]KeyboardButton{
			{
				KeyboardButton{
					Text:   "Edit Gif",
					WebApp: &WebApp{URL: url},
				},
			},
		},
	}
}

func (rkm ReplyKeyboardMarkup) FromSlicesOfStrings(rows [][]string) *ReplyKeyboardMarkup {
	keyboard := &ReplyKeyboardMarkup{}
	for _, row := range rows {
		var rowSet []KeyboardButton
		for _, column := range row {
			rowSet = append(rowSet, KeyboardButton{Text: column})
		}
		keyboard.keyboard = append(keyboard.keyboard, rowSet)
	}
	return keyboard
}

func (rkm *ReplyKeyboardMarkup) SetResizeKeyboard(resize bool) *ReplyKeyboardMarkup {
	rkm.resizeKeyboard = resize
	return rkm
}

func (rkm *ReplyKeyboardMarkup) SetOneTimeKeyboard(oneTime bool) *ReplyKeyboardMarkup {
	rkm.oneTimeKeyboard = oneTime
	return rkm
}

func (rkm *ReplyKeyboardMarkup) SetSelective(selective bool) *ReplyKeyboardMarkup {
	rkm.selective = selective
	return rkm
}
