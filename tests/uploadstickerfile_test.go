package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-telegram-bot-api"
)

func TestUploadStickerFile(t *testing.T) {
	bot, err := tgbotapi.New(Tests{}.Defaults().BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	// AgACAgQAAxkDAAECEudeYF9HGJ7VZzj6XzLdKNrNiy4IoAACDrIxGyWdCVNLOHwy3uUB1V0pthsABAEAAwIAA20AA1LzAwABGAQ
	p := bot.UploadStickerFile()
	p.SetUserId(Tests{}.Defaults().UserId).SetStickerFilePath("photo.png")
	m, err := bot.Send(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*m.File)

}
