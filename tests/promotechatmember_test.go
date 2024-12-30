package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestPromoteChatMember(t *testing.T) {
	bot, err := tgbotapi.New("796493295:AAE3EGLAnba_XAsp_ts3sbPTHpW3nitBc4s")
	if err != nil {
		fmt.Println(err)
		return
	}
	up := (bot.PromoteChatMember()).SetUserId(207260097).SetChatId(Tests{}.Defaults().ChatId).EnablePinMessages()
	resp, err := bot.Send(up)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Bool != nil && *resp.Bool {
		fmt.Println("user can now pin messages")
		return
	}
	// fmt.Println(resp.Ok, resp.Message)
}
