package tgbotapi

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"slices"
	"strings"
	"time"
)

type SlogPeriodic struct {
	slog.Handler

	bot      *TelegramBot
	chatID   int64
	interval time.Duration
	title    string
	levels   []slog.Level
	l        *time.Location

	logs chan string
}

func NewSlogPeriodic(
	handler slog.Handler,
	bot *TelegramBot,
	chatID int64,
	interval time.Duration,
	title string,
	levels ...slog.Level,
) slog.Handler {

	tp := SlogPeriodic{
		Handler: handler,

		bot:      bot,
		chatID:   chatID,
		interval: interval,
		logs:     make(chan string),
		title:    title,
		levels:   levels,

		l: time.Local,
	}

	go tp.periodicSender()

	return &tp
}

func (l *SlogPeriodic) Enabled(ctx context.Context, level slog.Level) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	return l.Handler.Enabled(ctx, level)
}

func (l *SlogPeriodic) Handle(ctx context.Context, r slog.Record) error {
	if slices.Contains(l.levels, r.Level) {
		go func() {
			l.logs <- fmt.Sprintf("%s - %s", r.Time.In(l.l).Format("2006-01-02 15:04:05MST"), r.Message)
		}()
	}

	return l.Handler.Handle(ctx, r)
}

func (l *SlogPeriodic) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewSlog(l.Handler.WithAttrs(attrs), l.bot, l.chatID, l.levels...)
}

func (l *SlogPeriodic) WithGroup(name string) slog.Handler {
	return NewSlog(l.Handler.WithGroup(name), l.bot, l.chatID, l.levels...)
}

func (l *SlogPeriodic) periodicSender() {
	ticker := time.NewTicker(l.interval)

	data := []string{}

	for {
		select {
		case <-ticker.C:
			if len(data) > 0 {
				_, err := l.bot.Send(l.bot.Document().
					SetDocumentReader(strings.NewReader(strings.Join(data, "\n")),
						fmt.Sprintf("logs_%s.txt", l.title)).
					SetChatId(l.chatID))
				if err != nil {
					log.Println(err)
				}
				data = []string{}
			}
			break
		case txt := <-l.logs:
			data = append(data, txt)
			break
		}
	}
}
