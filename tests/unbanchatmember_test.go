package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestUnbanChatMember(t *testing.T) {
	bot, err := tgbotapi.New("796493295:AAE3EGLAnba_XAsp_ts3sbPTHpW3nitBc4s")
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.UnbanChatMember()).SetUserId(207260097).SetChatId(Tests{}.Defaults().ChatId)
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Bool != nil && *resp.Bool {
		fmt.Println("unbanned")
		return
	}
	// fmt.Println(resp.Ok, resp.Message)
}
