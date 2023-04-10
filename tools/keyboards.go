package tools

import "github.com/aliforever/go-telegram-bot-api/structs"

type Keyboards struct {
}

func (k Keyboards) NewInlineKeyboard() *structs.InlineKeyboardMarkup {
	return &structs.InlineKeyboardMarkup{}
}

func (k Keyboards) NewInlineKeyboardFromSlicesOfMaps(slicesOfMaps [][]map[string]string) *structs.InlineKeyboardMarkup {
	return (&structs.InlineKeyboardMarkup{}).FromSlicesOfMaps(slicesOfMaps)
}

func (k Keyboards) NewReplyKeyboardFromSlicesOfStrings(slicesOfStrings [][]string) *structs.ReplyKeyboardMarkup {
	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(slicesOfStrings).SetResizeKeyboard(true)
}

// NewReplyKeyboardFromSliceOfStrings creates a new ReplyKeyboardMarkup from a slice of strings.
// The slice of strings will be divided into slices of strings with the length of bpr (buttons per row).
func (k Keyboards) NewReplyKeyboardFromSliceOfStrings(sliceOfStrings []string, bpr int) *structs.ReplyKeyboardMarkup {
	var rows [][]string
	var row []string

	for _, s := range sliceOfStrings {
		row = append(row, s)

		if len(row) >= bpr {
			rows = append(rows, row)
			row = []string{}
		}
	}

	if len(row) > 0 {
		rows = append(rows, row)
	}

	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(rows).SetResizeKeyboard(true)
}
