package tgbotapi

import (
	"github.com/sirupsen/logrus"
	"log"
)

type LogrusHook struct {
	bot    *TelegramBot
	chatID int64
}

func NewLogrusHook(bot *TelegramBot, chatID int64) logrus.Hook {
	return LogrusHook{bot: bot, chatID: chatID}
}

func (t LogrusHook) Levels() []logrus.Level {
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

func (t LogrusHook) Fire(entry *logrus.Entry) error {
	go func() {
		_, err := t.bot.Send(t.bot.Message().SetText(entry.Message).SetChatId(t.chatID))
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
