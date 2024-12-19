package tgbotapi

import (
	"encoding/json"

	"github.com/aliforever/go-telegram-bot-api/tools"

	"github.com/aliforever/go-telegram-bot-api/structs"
)

type sendInvoice struct {
	parent                    *TelegramBot
	chatId                    interface{}
	messageThreadId           int64
	title                     string
	description               string
	payload                   string
	providerToken             string
	currency                  string
	prices                    []structs.LabeledPrice
	maxTipAmount              int
	suggestedTipAmounts       []int
	startParameter            string
	providerData              string
	photoURL                  string
	photoSize                 int64
	photoWidth                int
	photoHeight               int
	needName                  *bool
	needPhoneNumber           *bool
	needEmail                 *bool
	needShippingAddress       *bool
	sendPhoneNumberToProvider *bool
	sendEmailToProvider       *bool
	isFlexible                *bool
	disableNotification       *bool
	protectContent            *bool
	allowPaidBroadcast        *bool
	messageEffectId           string
	replyParameters           *struct {
		messageID                int64
		chatID                   any
		allowSendingWithoutReply bool
		quote                    string
		quoteParseMode           string
		quoteEntities            []struct {
			entityType string
			offset     int64
			length     int64
			url        string
			user       *struct {
				id                      int64
				isBot                   bool
				firstName               string
				lastName                string
				username                string
				languageCode            string
				canJoinGroups           bool
				canReadAllGroupMessages bool
				supportsInlineQueries   bool
			}
			language string
		}
		quotePosition int64
	}
	replyMarkup interface{}
}

func (m *sendInvoice) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChatId                    interface{}            `json:"chat_id"`
		MessageThreadId           int64                  `json:"message_thread_id,omitempty"`
		Title                     string                 `json:"title"`
		Description               string                 `json:"description"`
		Payload                   string                 `json:"payload"`
		ProviderToken             string                 `json:"provider_token"`
		Currency                  string                 `json:"currency"`
		Prices                    []structs.LabeledPrice `json:"prices"`
		MaxTipAmount              int                    `json:"max_tip_amount,omitempty"`
		SuggestedTipAmounts       []int                  `json:"suggested_tip_amounts,omitempty"`
		StartParameter            string                 `json:"start_parameter,omitempty"`
		ProviderData              string                 `json:"provider_data,omitempty"`
		PhotoURL                  string                 `json:"photo_url,omitempty"`
		PhotoSize                 int64                  `json:"photo_size,omitempty"`
		PhotoWidth                int                    `json:"photo_width,omitempty"`
		PhotoHeight               int                    `json:"photo_height,omitempty"`
		NeedName                  *bool                  `json:"need_name,omitempty"`
		NeedPhoneNumber           *bool                  `json:"need_phone_number,omitempty"`
		NeedEmail                 *bool                  `json:"need_email,omitempty"`
		NeedShippingAddress       *bool                  `json:"need_shipping_address,omitempty"`
		SendPhoneNumberToProvider *bool                  `json:"send_phone_number_to_provider,omitempty"`
		SendEmailToProvider       *bool                  `json:"send_email_to_provider,omitempty"`
		IsFlexible                *bool                  `json:"is_flexible,omitempty"`
		DisableNotification       *bool                  `json:"disable_notification,omitempty"`
		ProtectContent            *bool                  `json:"protect_content,omitempty"`
		AllowPaidBroadcast        *bool                  `json:"allow_paid_broadcast,omitempty"`
		MessageEffectId           string                 `json:"message_effect_id,omitempty"`
		ReplyMarkup               interface{}            `json:"reply_markup,omitempty"`
	}{
		ChatId:                    m.chatId,
		MessageThreadId:           m.messageThreadId,
		Title:                     m.title,
		Description:               m.description,
		Payload:                   m.payload,
		ProviderToken:             m.providerToken,
		Currency:                  m.currency,
		Prices:                    m.prices,
		MaxTipAmount:              m.maxTipAmount,
		SuggestedTipAmounts:       m.suggestedTipAmounts,
		StartParameter:            m.startParameter,
		ProviderData:              m.providerData,
		PhotoURL:                  m.photoURL,
		PhotoSize:                 m.photoSize,
		PhotoWidth:                m.photoWidth,
		PhotoHeight:               m.photoHeight,
		NeedName:                  m.needName,
		NeedPhoneNumber:           m.needPhoneNumber,
		NeedEmail:                 m.needEmail,
		NeedShippingAddress:       m.needShippingAddress,
		SendPhoneNumberToProvider: m.sendPhoneNumberToProvider,
		SendEmailToProvider:       m.sendEmailToProvider,
		IsFlexible:                m.isFlexible,
		DisableNotification:       m.disableNotification,
		ProtectContent:            m.protectContent,
		AllowPaidBroadcast:        m.allowPaidBroadcast,
		MessageEffectId:           m.messageEffectId,
		ReplyMarkup:               m.replyMarkup,
	})
}

func (m *sendInvoice) response() interface{} {
	return &structs.Message{}
}

func (m *sendInvoice) method() string {
	return "POST"
}

func (m *sendInvoice) endpoint() string {
	return "sendInvoice"
}

func (m *sendInvoice) SetChatId(chatId int64) *sendInvoice {
	m.chatId = chatId
	return m
}

func (m *sendInvoice) SetChatUsername(username string) *sendInvoice {
	m.chatId = username
	return m
}

func (m *sendInvoice) SetMessageThreadId(messageThreadId int64) *sendInvoice {
	m.messageThreadId = messageThreadId
	return m
}

func (m *sendInvoice) SetTitle(title string) *sendInvoice {
	m.title = title
	return m
}

func (m *sendInvoice) SetDescription(description string) *sendInvoice {
	m.description = description
	return m
}

func (m *sendInvoice) SetPayload(payload string) *sendInvoice {
	m.payload = payload
	return m
}

func (m *sendInvoice) SetProviderToken(providerToken string) *sendInvoice {
	m.providerToken = providerToken
	return m
}

func (m *sendInvoice) SetCurrency(currency string) *sendInvoice {
	m.currency = currency
	return m
}

func (m *sendInvoice) AddLabeledPrice(label string, amount int) *sendInvoice {
	m.prices = append(m.prices, structs.LabeledPrice{
		Label:  label,
		Amount: amount,
	})
	return m
}

func (m *sendInvoice) SetMaxTipAmount(maxTipAmount int) *sendInvoice {
	m.maxTipAmount = maxTipAmount
	return m
}

func (m *sendInvoice) AddSuggestedTipAmount(amount int) *sendInvoice {
	m.suggestedTipAmounts = append(m.suggestedTipAmounts, amount)
	return m
}

func (m *sendInvoice) SetStartParameter(startParameter string) *sendInvoice {
	m.startParameter = startParameter
	return m
}

func (m *sendInvoice) SetProviderData(providerData string) *sendInvoice {
	m.providerData = providerData
	return m
}

func (m *sendInvoice) SetPhotoURL(photoURL string) *sendInvoice {
	m.photoURL = photoURL
	return m
}

func (m *sendInvoice) SetPhotoSize(photoSize int64) *sendInvoice {
	m.photoSize = photoSize
	return m
}

func (m *sendInvoice) SetPhotoWidth(photoWidth int) *sendInvoice {
	m.photoWidth = photoWidth
	return m
}

func (m *sendInvoice) SetPhotoHeight(photoHeight int) *sendInvoice {
	m.photoHeight = photoHeight
	return m
}

func (m *sendInvoice) SetNeedName(needName bool) *sendInvoice {
	m.needName = &needName
	return m
}

func (m *sendInvoice) SetNeedPhoneNumber(needPhoneNumber bool) *sendInvoice {
	m.needPhoneNumber = &needPhoneNumber
	return m
}

func (m *sendInvoice) SetNeedEmail(needEmail bool) *sendInvoice {
	m.needEmail = &needEmail
	return m
}

func (m *sendInvoice) SetNeedShippingAddress(needShippingAddress bool) *sendInvoice {
	m.needShippingAddress = &needShippingAddress
	return m
}

func (m *sendInvoice) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) *sendInvoice {
	m.sendPhoneNumberToProvider = &sendPhoneNumberToProvider
	return m
}

func (m *sendInvoice) SetSendEmailToProvider(sendEmailToProvider bool) *sendInvoice {
	m.sendEmailToProvider = &sendEmailToProvider
	return m
}

func (m *sendInvoice) SetIsFlexible(isFlexible bool) *sendInvoice {
	m.isFlexible = &isFlexible
	return m
}

func (m *sendInvoice) SetDisableNotification(disableNotification bool) *sendInvoice {
	m.disableNotification = &disableNotification
	return m
}

func (m *sendInvoice) SetProtectContent(protectContent bool) *sendInvoice {
	m.protectContent = &protectContent
	return m
}

func (m *sendInvoice) SetAllowPaidBroadcast(allowPaidBroadcast bool) *sendInvoice {
	m.allowPaidBroadcast = &allowPaidBroadcast
	return m
}

func (m *sendInvoice) SetMessageEffectId(messageEffectId string) *sendInvoice {
	m.messageEffectId = messageEffectId
	return m
}

/*func (m *sendInvoice) SetReplyParameters(replyParameters interface{}) *sendInvoice {
	m.replyParameters = tools.ParseReplyParameters(replyParameters)
	return m
}*/

func (m *sendInvoice) SetReplyMarkup(markup interface{}) *sendInvoice {
	m.replyMarkup = tools.ParseReplyMarkup(markup)

	return m
}
