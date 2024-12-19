package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendInvoice(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := bot.Send(bot.Invoice().
		SetChatId(Tests{}.Defaults().ChatId).
		SetTitle("Test Payment").
		SetDescription("Test Description").
		SetCurrency("XTR").
		SetPayload("payload_1234").
		// AddSuggestedTipAmount(1).
		// SetMaxTipAmount(2).
		// SetStartParameter("start_param_unique_v1").
		AddLabeledPrice("XTR", 1))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
}
