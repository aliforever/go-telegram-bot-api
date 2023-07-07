package tgbotapi

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
	"time"
)

type LogrusPeriodicHook struct {
	bot      *TelegramBot
	chatID   int64
	interval time.Duration
	logs     chan string
	title    string
	levels   []logrus.Level
}

func NewLogrusPeriodicHook(
	bot *TelegramBot,
	chatID int64,
	interval time.Duration,
	title string,
	levels ...logrus.Level,
) logrus.Hook {

	tp := LogrusPeriodicHook{
		bot:      bot,
		chatID:   chatID,
		interval: interval,
		logs:     make(chan string),
		title:    title,
		levels:   levels,
	}

	go tp.periodicSender()

	return tp
}

func (t LogrusPeriodicHook) Levels() []logrus.Level {
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

func (t LogrusPeriodicHook) Fire(entry *logrus.Entry) error {
	go func() {
		t.logs <- fmt.Sprintf("%s - %s", time.Now().Format("2006-01-02 15:04:05"), entry.Message)
	}()

	return nil
}

func (t LogrusPeriodicHook) periodicSender() {
	ticker := time.NewTicker(t.interval)

	data := []string{}

	for {
		select {
		case <-ticker.C:
			if len(data) > 0 {
				_, err := t.bot.Send(t.bot.Document().
					SetDocumentReader(strings.NewReader(strings.Join(data, "\n")),
						fmt.Sprintf("logs_%s.txt", t.title)).
					SetChatId(t.chatID))
				if err != nil {
					log.Println(err)
				}
				data = []string{}
			}
			break
		case txt := <-t.logs:
			data = append(data, txt)
			break
		}
	}
}
