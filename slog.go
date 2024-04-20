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

type Slog struct {
	slog.Handler

	levels []slog.Level

	bot    *TelegramBot
	chatID int64

	l *time.Location
}

func NewSlog(handler slog.Handler, bot *TelegramBot, chatID int64, levels ...slog.Level) slog.Handler {
	return &Slog{
		Handler: handler,
		levels:  levels,

		bot:    bot,
		chatID: chatID,

		l: time.Local,
	}
}

func NewSlogWithLocation(
	handler slog.Handler,
	bot *TelegramBot,
	chatID int64,
	l *time.Location,
	levels ...slog.Level,
) slog.Handler {
	return &Slog{
		Handler: handler,
		levels:  levels,

		bot:    bot,
		chatID: chatID,

		l: l,
	}
}

func (l *Slog) Enabled(ctx context.Context, level slog.Level) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	return l.Handler.Enabled(ctx, level)
}

func (l *Slog) Handle(ctx context.Context, r slog.Record) error {
	if slices.Contains(l.levels, r.Level) {
		go func() {
			msg := r.Message

			var attrs []string

			r.Attrs(func(attr slog.Attr) bool {
				attrs = append(attrs, fmt.Sprintf("%s: %s", attr.Key, attr.Value.String()))

				return true
			})

			if len(attrs) > 0 {
				msg = fmt.Sprintf("%s - %s", msg, strings.Join(attrs, " | "))
			}

			message := fmt.Sprintf("%s - %s", r.Time.In(l.l).Format("2006-01-02 15:04:05MST"), msg)

			_, err := l.bot.Send(l.bot.Message().SetText(message).SetChatId(l.chatID))
			if err != nil {
				r.Message = err.Error()
				if err := l.Handler.Handle(ctx, r); err != nil {
					log.Println(err)
				}
			}
		}()
	}

	return l.Handler.Handle(ctx, r)
}

func (l *Slog) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewSlog(l.Handler.WithAttrs(attrs), l.bot, l.chatID, l.levels...)
}

func (l *Slog) WithGroup(name string) slog.Handler {
	return NewSlog(l.Handler.WithGroup(name), l.bot, l.chatID, l.levels...)
}
