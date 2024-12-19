package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestSendInvoice(t *testing.T) {
	// bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	bot, err := tgbotapi.New("1668694512:AAHVj8lqmudyZCEVrhKfrVCMv8oN_182F9A")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := bot.Send(bot.Invoice().
		// SetChatId(Tests{}.Defaults().ChatId).
		SetChatId(81997375).
		SetTitle("Increase Balance").
		SetDescription("Increase Balance to Use Premium Download").
		SetPayload(fmt.Sprintf("%d_%d", 81997375, 50)).
		SetCurrency("XTR").
		AddLabeledPrice("XTR", 100))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
}
