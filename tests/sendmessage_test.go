package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendMessage(t *testing.T) {
	/*var text tgbotapi.Text
	tex := text.AddWithNewLine("Hello").AddWithNewLine("Boy")*/
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	ik := [][]map[string]string{
		{
			{
				"text":          "Back",
				"callback_data": "data17",
			},
		},
		{
			{
				"text":          "Reset",
				"callback_data": "data19",
			},
		},
		{
			{
				"text":          "📌 ↔️ Width: 200px",
				"callback_data": "data1",
			},
			{
				"text":          "➡️ X: 10px",
				"callback_data": "data2",
			},
		},
		{
			{
				"text":          "↕️ Height: 100px",
				"callback_data": "data3",
			},
			{
				"text":          "⬇️ Y: 10px",
				"callback_data": "data4",
			},
		},
		{
			{
				"text":          "➕ 1",
				"callback_data": "data5",
			},
			{
				"text":          "➕ 5",
				"callback_data": "data6",
			},
			{
				"text":          "➕ 10",
				"callback_data": "data7",
			},
			{
				"text":          "➕ 20",
				"callback_data": "data8",
			},
		},
		{
			{
				"text":          "➕ 50",
				"callback_data": "data50",
			},
			{
				"text":          "➕ 100",
				"callback_data": "data60",
			},
			{
				"text":          "➕ 300",
				"callback_data": "data70",
			},
			{
				"text":          "➕ 500",
				"callback_data": "data80",
			},
		},
		{
			{
				"text":          "➖ 1",
				"callback_data": "data9",
			},
			{
				"text":          "➖ 5",
				"callback_data": "data10",
			},
			{
				"text":          "➖ 10",
				"callback_data": "data11",
			},
			{
				"text":          "➖ 20",
				"callback_data": "data12",
			},
		},
		{
			{
				"text":          "➖ 50",
				"callback_data": "data13",
			},
			{
				"text":          "➖ 100",
				"callback_data": "data14",
			},
			{
				"text":          "➖ 300",
				"callback_data": "data15",
			},
			{
				"text":          "➖ 500",
				"callback_data": "data16",
			},
		},
		{
			{
				"text":          "Apply",
				"callback_data": "data18",
			},
		},
	}

	_, err = bot.Send(bot.Message().
		SetChatId(Tests{}.Defaults().ChatId).
		SetText("testtttttttttttttttttttttttttttt").
		SetReplyMarkup(bot.Tools.Keyboards.NewInlineKeyboardFromSlicesOfMaps(ik)))
	if err != nil {
		fmt.Println(err, "err")
		return
	}

	/*message := (bot.Message()).SetChatId(Tests{}.Defaults().UserId).SetText("good idea" + bot.Tools.Strings.EmptyNewLine())
	m, err := bot.Send(message)
	if err != nil {
		fmt.Println(err, "err")
		return
	}
	fmt.Println(m)*/
}
