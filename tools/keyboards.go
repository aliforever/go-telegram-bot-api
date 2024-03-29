package tools

import (
	"github.com/aliforever/go-telegram-bot-api/structs"
	"slices"
)

type Keyboards struct {
}

func (k Keyboards) NewInlineKeyboard() *structs.InlineKeyboardMarkup {
	return &structs.InlineKeyboardMarkup{}
}

func (k Keyboards) NewInlineKeyboardFromSlicesOfMaps(slicesOfMaps [][]map[string]string) *structs.InlineKeyboardMarkup {
	return (&structs.InlineKeyboardMarkup{}).FromSlicesOfMaps(slicesOfMaps)
}

func (k Keyboards) NewInlineKeyboardFromSlicesOfMapWithFormation(
	slicesOfMaps []map[string]string,
	maxPerRow int,
	formation []int,
	reverseButtonOrderInRows bool,
) *structs.InlineKeyboardMarkup {

	var rows [][]map[string]string
	var row []map[string]string

	for _, s := range slicesOfMaps {
		row = append(row, s)

		if k.shouldBreakRow(len(row), len(rows), maxPerRow, formation) || s == nil {
			if reverseButtonOrderInRows {
				slices.Reverse(row)
			}

			rows = append(rows, row)
			row = []map[string]string{}
		}
	}

	if len(row) > 0 {
		if reverseButtonOrderInRows {
			slices.Reverse(row)
		}

		rows = append(rows, row)
	}

	return (&structs.InlineKeyboardMarkup{}).FromSlicesOfMaps(rows)
}

func (k Keyboards) NewReplyKeyboardFromSlicesOfStrings(slicesOfStrings [][]string) *structs.ReplyKeyboardMarkup {
	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(slicesOfStrings).SetResizeKeyboard(true)
}

// NewReplyKeyboardFromSliceOfStrings creates a new ReplyKeyboardMarkup from a slice of strings.
// The slice of strings will be divided into slices of strings with the length of bpr (buttons per row).
func (k Keyboards) NewReplyKeyboardFromSliceOfStrings(
	sliceOfStrings []string,
	bpr int,
	reverseButtonOrderInRows bool,
) *structs.ReplyKeyboardMarkup {

	var rows [][]string
	var row []string

	for _, s := range sliceOfStrings {
		row = append(row, s)

		if len(row) >= bpr {
			if reverseButtonOrderInRows {
				slices.Reverse(row)
			}

			rows = append(rows, row)
			row = []string{}
		}
	}

	if len(row) > 0 {
		if reverseButtonOrderInRows {
			slices.Reverse(row)
		}

		rows = append(rows, row)
	}

	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(rows).SetResizeKeyboard(true)
}

// NewReplyKeyboardFromSliceOfStringsWithFormation creates a new ReplyKeyboardMarkup from a slice of strings.
// The slice of strings will be divided into slices of strings with the length of bpr (buttons per row).
// The formation of the keyboard is defined by the slice of integers.
func (k Keyboards) NewReplyKeyboardFromSliceOfStringsWithFormation(
	sliceOfStrings []string,
	maxBtnPerRow int,
	buttonFormation []int,
	reverseButtonOrderInRows bool,
) *structs.ReplyKeyboardMarkup {

	var rows [][]string
	var row []string

	for _, s := range sliceOfStrings {
		row = append(row, s)

		if k.shouldBreakRow(len(row), len(rows), maxBtnPerRow, buttonFormation) || s == "" {
			if reverseButtonOrderInRows {
				slices.Reverse(row)
			}

			rows = append(rows, row)
			row = []string{}
		}
	}

	if len(row) > 0 {
		if reverseButtonOrderInRows {
			slices.Reverse(row)
		}

		rows = append(rows, row)
	}

	return (&structs.ReplyKeyboardMarkup{}).FromSlicesOfStrings(rows).SetResizeKeyboard(true)
}

func (k Keyboards) shouldBreakRow(
	rowLength int, rowsLength int, maxButtonPerRow int, buttonFormation []int) bool {

	defaultCond := rowLength >= maxButtonPerRow

	if len(buttonFormation) > rowsLength {
		defaultCond = rowLength >= buttonFormation[rowsLength]
	}

	return defaultCond
}

func ParseReplyMarkup(markup interface{}) interface{} {
	switch val := markup.(type) {
	case *structs.ReplyKeyboardMarkup:
		if val == nil {
			return nil
		}
	case *structs.InlineKeyboardMarkup:
		if val == nil {
			return nil
		}
	default:
		return nil
	}

	return markup
}
