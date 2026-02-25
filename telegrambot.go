package tgbotapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"github.com/aliforever/go-telegram-bot-api/tools"
)

type TelegramBot struct {
	apiToken        string
	apiUrl          string
	username        string
	recipientChatId int64
	updates         chan Update
	randSource      rand.Source
	client          *resty.Client
	Tools           tools.Tools
	logger          *slog.Logger
	logEvents       bool
}

func NewSelfHosted(apiToken string, address string) (tb *TelegramBot, err error) {
	address = fmt.Sprintf(`%s/bot%s/`, address, apiToken)

	client := resty.New()
	client.SetHostURL(address)
	client.SetDoNotParseResponse(true)

	bot := &TelegramBot{
		apiToken:   apiToken,
		apiUrl:     address,
		client:     client,
		updates:    make(chan Update),
		randSource: rand.NewSource(time.Now().UnixNano()),
		Tools:      tools.Tools{},
	}

	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}

	if resp.User == nil {
		err = errors.New("empty_user_response")
	}

	tb = bot

	return
}

func New(apiToken string, options ...*Options) (tb *TelegramBot, err error) {
	address := fmt.Sprintf(`https://api.telegram.org/bot%s/`, apiToken)

	var option *Options

	if len(options) > 0 {
		option = options[0]
	}

	var (
		logger       *slog.Logger
		logResponses bool
	)

	if option != nil {
		if option.apiURL != nil {
			address = fmt.Sprintf("%s/bot%s/", *option.apiURL, apiToken)
		}

		if option.logger != nil {
			logger = option.logger
		}

		logResponses = option.logResponses
	}

	client := resty.New()
	client.SetBaseURL(address)
	client.SetDoNotParseResponse(true)

	bot := &TelegramBot{
		apiToken:   apiToken,
		apiUrl:     address,
		client:     client,
		updates:    make(chan Update),
		randSource: rand.NewSource(time.Now().UnixNano()),
		Tools:      tools.Tools{},
		logger:     logger,
		logEvents:  logResponses,
	}

	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}

	if resp.User == nil {
		err = errors.New("empty_user_response")
		return nil, err
	}

	tb = bot

	return
}

func NewTelegramBotWithLogger(apiToken string, logger *slog.Logger) (tb *TelegramBot, err error) {
	address := fmt.Sprintf(`https://api.telegram.org/bot%s/`, apiToken)
	client := resty.New()
	client.SetHostURL(address)
	client.SetDoNotParseResponse(true)
	bot := &TelegramBot{
		apiToken:   apiToken,
		apiUrl:     address,
		client:     client,
		updates:    make(chan Update),
		Tools:      tools.Tools{},
		randSource: rand.NewSource(time.Now().UnixNano()),
		logger:     logger,
	}
	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}
	if resp.User == nil {
		err = errors.New("empty_user_response")
	}
	tb = bot
	return
}

func NewTelegramBotWithLoggerEvents(apiToken string, logger *slog.Logger) (tb *TelegramBot, err error) {
	address := fmt.Sprintf(`https://api.telegram.org/bot%s/`, apiToken)
	client := resty.New()
	client.SetHostURL(address)
	client.SetDoNotParseResponse(true)
	bot := &TelegramBot{
		apiToken:   apiToken,
		apiUrl:     address,
		client:     client,
		updates:    make(chan Update),
		Tools:      tools.Tools{},
		randSource: rand.NewSource(time.Now().UnixNano()),
		logger:     logger,
		logEvents:  true,
	}
	var resp *Response
	resp, err = bot.Send(bot.GetMe())
	if err != nil {
		return
	}
	if resp.User == nil {
		err = errors.New("empty_user_response")
	}
	tb = bot
	return
}

func (tb *TelegramBot) SetAPIServerUrl(address string) {
	if address[len(address)-1] == '/' {
		address = address[:len(address)-1]
	}
	tb.client.SetBaseURL(fmt.Sprintf("%s/bot%s/", address, tb.apiToken))
}

func (tb *TelegramBot) SetRecipientChatId(chatId int64) {
	tb.recipientChatId = chatId
}

func (tb *TelegramBot) GetMe() (m *getMe) {
	m = &getMe{parent: tb}
	return
}

func (tb *TelegramBot) GetChat() (m *getChat) {
	m = &getChat{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}

	return
}

func (tb *TelegramBot) LeaveChat() (m *leaveChat) {
	m = &leaveChat{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}

	return
}

func (tb *TelegramBot) GetChatMember() (m *getChatMember) {
	m = &getChatMember{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}

	return
}

func (tb *TelegramBot) DeleteMessage() (m *deleteMessage) {
	m = &deleteMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) DeleteMessages() (m *deleteMessages) {
	m = &deleteMessages{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Message() (m *sendMessage) {
	m = &sendMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) CopyMessage() (m *copyMessage) {
	m = &copyMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageText() (m *editMessageText) {
	m = &editMessageText{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageCaption() (m *editMessageCaption) {
	m = &editMessageCaption{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageReplyMarkup() (m *editMessageReplyMarkup) {
	m = &editMessageReplyMarkup{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) ForwardMessage() (m *forwardMessage) {
	m = &forwardMessage{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Photo() (m *sendPhoto) {
	m = &sendPhoto{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Audio() (m *sendAudio) {
	m = &sendAudio{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Video() (m *sendVideo) {
	m = &sendVideo{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Voice() (m *sendVoice) {
	m = &sendVoice{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) VideoNote() (m *sendVideoNote) {
	m = &sendVideoNote{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) MediaGroup() (m *sendMediaGroup) {
	m = &sendMediaGroup{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Document() (m *sendDocument) {
	m = &sendDocument{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Dice() (m *sendDice) {
	m = &sendDice{parent: tb}

	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}

	return
}

func (tb *TelegramBot) Animation() (m *sendAnimation) {
	m = &sendAnimation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Location() (m *sendLocation) {
	m = &sendLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Contact() (m *sendContact) {
	m = &sendContact{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Venue() (m *sendVenue) {
	m = &sendVenue{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Poll() (m *sendPoll) {
	m = &sendPoll{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) ChatAction() (m *sendChatAction) {
	m = &sendChatAction{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) UserProfilePhotos() (m *getUserProfilePhotos) {
	m = &getUserProfilePhotos{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetUserId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) File() (m *getFile) {
	m = &getFile{parent: tb}
	return
}

func (tb *TelegramBot) FileBytesChunk() (m *getFileBytesChunk) {
	m = &getFileBytesChunk{parent: tb}
	return
}

func (tb *TelegramBot) KickChatMember() (m *kickChatMember) {
	m = &kickChatMember{parent: tb}
	return
}

func (tb *TelegramBot) AnswerCallbackQuery() (m *answerCallbackQuery) {
	m = &answerCallbackQuery{parent: tb}
	return
}

func (tb *TelegramBot) AnswerPreCheckoutQuery() (m *answerPreCheckoutQuery) {
	m = &answerPreCheckoutQuery{parent: tb}
	return
}

func (tb *TelegramBot) AnswerWebAppQuery() (m *answerWebAppQuery) {
	m = &answerWebAppQuery{parent: tb}
	return
}

func (tb *TelegramBot) PromoteChatMember() (m *promoteChatMember) {
	m = &promoteChatMember{parent: tb}
	return
}

func (tb *TelegramBot) RestrictChatMember() (m *restrictChatMember) {
	m = &restrictChatMember{parent: tb}
	return
}

func (tb *TelegramBot) UnbanChatMember() (m *unbanChatMember) {
	m = &unbanChatMember{parent: tb}
	return
}

func (tb *TelegramBot) SetChatAdministratorCustomTitle() (m *setChatAdministratorCustomTitle) {
	m = &setChatAdministratorCustomTitle{parent: tb}
	return
}

func (tb *TelegramBot) EditLiveLocation() (m *editMessageLiveLocation) {
	m = &editMessageLiveLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) StopLiveLocation() (m *stopMessageLiveLocation) {
	m = &stopMessageLiveLocation{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) EditMessageMedia() (m *editMessageMedia) {
	m = &editMessageMedia{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) Sticker() (m *sendSticker) {
	m = &sendSticker{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

func (tb *TelegramBot) GetStickerSet() (m *getStickerSet) {
	m = &getStickerSet{parent: tb}
	return
}

func (tb *TelegramBot) CreateNewStickerSet() (m *createNewStickerSet) {
	m = &createNewStickerSet{parent: tb}
	return
}

func (tb *TelegramBot) UploadStickerFile() (m *uploadStickerFile) {
	m = &uploadStickerFile{parent: tb}
	return
}

func (tb *TelegramBot) DeleteStickerFromSet() (m *deleteStickerFromSet) {
	m = &deleteStickerFromSet{parent: tb}
	return
}

func (tb *TelegramBot) AddStickerToSet() (m *addStickerToSet) {
	m = &addStickerToSet{parent: tb}
	return
}

func (tb *TelegramBot) GetUpdates() (m *getUpdates) {
	m = &getUpdates{parent: tb}
	return
}

func (tb *TelegramBot) SetWebhook() (m *setWebhook) {
	m = &setWebhook{parent: tb}
	return
}

func (tb *TelegramBot) DeleteWebhook() (m *deleteWebhook) {
	m = &deleteWebhook{parent: tb}
	return
}

func (tb *TelegramBot) Invoice() (m *sendInvoice) {
	m = &sendInvoice{parent: tb}
	if tb.recipientChatId != 0 {
		m.SetChatId(tb.recipientChatId)
	}
	return
}

// StopReceivingUpdates TODO: Not working yet
func (tb *TelegramBot) StopReceivingUpdates() {
	close(tb.updates)
}

func (tb *TelegramBot) ListenWebhook(address string) (err error) {
	http.HandleFunc(fmt.Sprintf("/%s", tb.apiToken), func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		var u Update

		j := json.NewDecoder(request.Body)
		err := j.Decode(&u)
		if err != nil {
			fmt.Println("error decoding update", err.Error())
			return
		}

		tb.updates <- u

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("ok"))
	})

	return http.ListenAndServe(address, nil)
}

func (tb *TelegramBot) Updates() chan Update {
	return tb.updates
}

func (tb *TelegramBot) Send(config Config) (result *Response, err error) {
	request := tb.client.R()

	err = tb.prepareRequest(config, request)
	if err != nil {
		return nil, err
	}

	res, err := request.Execute(config.method(), config.endpoint())
	if err != nil {
		return nil, err
	}

	var raw []byte
	result, raw, err = tb.getMessageResponse(res, config)
	if err != nil {
		tb.logErrBytes(raw, fmt.Sprintf("%+v", config), err)
		return nil, err
	}

	return
}

func (tb *TelegramBot) SendRaw(config Config) (*resty.Response, error) {
	request := tb.client.R()

	err := tb.prepareRequest(config, request)
	if err != nil {
		return nil, err
	}

	return request.Execute(config.method(), config.endpoint())
}

func (tb *TelegramBot) SendWithOptions(config Config, options *SendOptions) (result *Response, err error) {
	client := tb.client

	request := client.R()

	if options != nil {
		if options.timeout != 0 {
			ctx, cancel := context.WithTimeout(context.Background(), options.timeout)
			defer cancel()
			request = request.SetContext(ctx)
		}
	}

	err = tb.prepareRequest(config, request)
	if err != nil {
		return nil, err
	}

	res, err := request.Execute(config.method(), config.endpoint())
	if err != nil {
		return nil, err
	}

	var raw []byte
	result, raw, err = tb.getMessageResponse(res, config)
	if err != nil {
		tb.logErrBytes(raw, fmt.Sprintf("%+v", config), err)
		return nil, err
	}

	return
}

// LogrusPeriodicHook is a logger implementing logrus.Hook that logs periodically
func (tb *TelegramBot) LogrusPeriodicHook(
	chatID int64,
	interval time.Duration,
	title string,
	levels ...logrus.Level,
) logrus.Hook {

	return NewLogrusPeriodic(tb, chatID, interval, title, levels...)
}

// LogrusHook is a logger implementing logrus.Hook that logs immediately
func (tb *TelegramBot) LogrusHook(chatID int64, levels ...logrus.Level) logrus.Hook {
	return NewLogrus(tb, chatID, levels...)
}

// SlogHandler is a logger implementing logrus.Hook that logs immediately
func (tb *TelegramBot) SlogHandler(handler slog.Handler, chatID int64, levels ...slog.Level) slog.Handler {
	return NewSlog(handler, tb, chatID, levels...)
}

// SlogPeriodicHandler is a logger implementing logrus.Hook that logs periodically
func (tb *TelegramBot) SlogPeriodicHandler(
	handler slog.Handler,
	chatID int64,
	interval time.Duration,
	title string,
	levels ...slog.Level,
) slog.Handler {
	return NewSlogPeriodic(handler, tb, chatID, interval, title, levels...)
}

func (tb *TelegramBot) logErrBytes(rawBytes []byte, msg string, err error) {
	if tb.logger != nil {
		tb.logger.Error(
			msg,
			slog.Any("error", err),
			slog.String("response", string(rawBytes)),
			slog.String("trace", string(debug.Stack())),
		)
	}
}
