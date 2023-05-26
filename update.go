package tgbotapi

import (
	"github.com/aliforever/go-telegram-bot-api/structs"
)

type Update struct {
	UpdateId          int64                `json:"update_id"`
	Message           *structs.Message     `json:"message,omitempty"`
	EditedMessage     *structs.Message     `json:"edited_message,omitempty"`
	ChannelPost       *structs.Message     `json:"channel_post,omitempty"`
	EditedChannelPost *structs.Message     `json:"edited_channel_post,omitempty"`
	InlineQuery       *structs.InlineQuery `json:"inline_query"`
	// ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery *structs.CallbackQuery     `json:"callback_query,omitempty"`
	MyChatMember  *structs.ChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember    *structs.ChatMemberUpdated `json:"chat_member,omitempty"`
	// ShippingQuery      ShippingQuery      `json:"shipping_query"`
	// PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll       *structs.Poll       `json:"poll,omitempty"`
	PollAnswer *structs.PollAnswer `json:"poll_answer,omitempty"`
	err        error
	raw        []byte
}

func newUpdateError(err error) Update {
	return Update{err: err}
}

func (u *Update) Error() error {
	return u.err
}

func (u *Update) Raw() []byte {
	return u.raw
}

func (u *Update) Chat() (chat *structs.Chat) {
	if u.Message != nil && u.Message.Chat != nil {
		return u.Message.Chat
	}

	if u.EditedMessage != nil && u.EditedMessage.Chat != nil {
		return u.EditedMessage.Chat
	}

	if u.CallbackQuery != nil && u.CallbackQuery.Message != nil && u.CallbackQuery.Message.Chat != nil {
		return u.CallbackQuery.Message.Chat
	}

	if u.ChatMember != nil && u.ChatMember.Chat != nil {
		return u.ChatMember.Chat
	}

	if u.MyChatMember != nil && u.MyChatMember.Chat != nil {
		return u.MyChatMember.Chat
	}

	// TODO: Add More Cases
	return nil
}

func (u *Update) From() (user *structs.User) {
	if u.Message != nil && u.Message.From != nil {
		return u.Message.From
	}

	if u.EditedMessage != nil && u.EditedMessage.From != nil {
		return u.EditedMessage.From
	}

	if u.CallbackQuery != nil && u.CallbackQuery.From != nil {
		return u.CallbackQuery.From
	}

	if u.ChatMember != nil && u.ChatMember.From != nil {
		return u.ChatMember.From
	}

	if u.MyChatMember != nil && u.MyChatMember.From != nil {
		return u.MyChatMember.From
	}

	// TODO: Add More Cases
	return
}

func (u *Update) PrivateMessageText() string {
	if u.Message != nil && u.Message.IsPrivate() && u.Message.Text != "" {
		return u.Message.Text
	}

	return ""
}
