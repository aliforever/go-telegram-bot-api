package tgbotapi

import (
	"log"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestNewSlog(t *testing.T) {
	opt := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, &opt)

	type args struct {
		handler slog.Handler
		bot     *TelegramBot
		chatID  int64
		levels  []slog.Level
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success 1",
			args: args{
				handler: handler,
				bot:     newTelegramBot(),
				chatID:  -1002122182579,
				levels: []slog.Level{
					slog.LevelDebug,
					slog.LevelInfo,
					slog.LevelWarn,
					slog.LevelError,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSlog(tt.args.handler, tt.args.bot, tt.args.chatID, tt.args.levels...)

			l := slog.New(got)
			l.Info("Info")
			time.Sleep(time.Second * 1)
			l.Debug("Debug")
			time.Sleep(time.Second * 1)
			l.Warn("Warn")
			time.Sleep(time.Second * 2)
			l.Error("Error")
			time.Sleep(time.Second * 2)
		})
	}
}

func newTelegramBot() *TelegramBot {
	token := ""

	bot, err := New(token)
	if err != nil {
		log.Panic(err)
	}

	return bot
}
