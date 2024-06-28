package tgbotapi

import (
	"fmt"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logrus struct {
	bot    *TelegramBot
	chatID int64
	levels []logrus.Level
}

func NewLogrus(bot *TelegramBot, chatID int64, levels ...logrus.Level) logrus.Hook {
	return Logrus{bot: bot, chatID: chatID, levels: levels}
}

func (t Logrus) Levels() []logrus.Level {
	if len(t.levels) > 0 {
		return t.levels
	}

	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}

func (t Logrus) Fire(entry *logrus.Entry) error {
	go func() {
		var data []string

		for key, val := range entry.Data {
			data = append(data, fmt.Sprintf("%s: %v", key, val))
		}

		text := entry.Message

		if len(data) > 0 {
			text += "\n" + strings.Join(data, "\n")
		}

		_, err := t.bot.Send(t.bot.Message().SetText(text).SetChatId(t.chatID))
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
