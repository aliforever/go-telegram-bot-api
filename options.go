package tgbotapi

import "log/slog"

type Options struct {
	apiURL       *string
	logger       *slog.Logger
	logResponses bool
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) SetApiURL(apiURL string) *Options {
	o.apiURL = &apiURL
	return o
}

func (o *Options) SetLogger(logger *slog.Logger) *Options {
	o.logger = logger
	return o
}

func (o *Options) SetLogResponses(logResponses bool) *Options {
	o.logResponses = logResponses
	return o
}
