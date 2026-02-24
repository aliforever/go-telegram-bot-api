package structs

type Chat struct {
	Id                    int64            `json:"id"`
	Type                  string           `json:"type"`
	Title                 string           `json:"title,omitempty"`
	Username              string           `json:"username,omitempty"`
	FirstName             string           `json:"first_name,omitempty"`
	LastName              string           `json:"last_name,omitempty"`
	Photo                 *ChatPhoto       `json:"photo,omitempty"`
	Description           string           `json:"description,omitempty"`
	InviteLink            string           `json:"invite_link,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
	Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         int64            `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int64            `json:"message_auto_delete_time,omitempty"`
	StickerSetName        string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatId          bool             `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation    `json:"location,omitempty"`
}
