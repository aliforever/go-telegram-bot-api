package structs

const (
	chatTypeGroup      = "group"
	chatTypeSuperGroup = "supergroup"
	chatTypePrivate    = "private"
	chatTypeChannel    = "channel"
)

type Message struct {
	MessageId                    int64                         `json:"message_id,omitempty"`
	From                         *User                         `json:"from,omitempty"`
	Date                         int64                         `json:"date,omitempty"`
	Chat                         *Chat                         `json:"chat,omitempty"`
	SenderChat                   *Chat                         `json:"sender_chat,omitempty"`
	ForwardFrom                  *User                         `json:"forward_from,omitempty,omitempty"`
	ForwardFromChat              *Chat                         `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId         int64                         `json:"forward_from_message_id,omitempty"`
	ForwardSignature             string                        `json:"forward_signature,omitempty"`
	ForwardSenderName            string                        `json:"forward_sender_name,omitempty"`
	ForwardDate                  int64                         `json:"forward_date,omitempty"`
	ReplyToMessage               *Message                      `json:"reply_to_message,omitempty"`
	EditDate                     int64                         `json:"edit_date,omitempty"`
	MediaGroupId                 string                        `json:"media_group_id,omitempty"`
	AuthorSignature              string                        `json:"author_signature,omitempty"`
	Text                         string                        `json:"text,omitempty"`
	Entities                     []MessageEntity               `json:"entities,omitempty"`
	CaptionEntities              []MessageEntity               `json:"caption_entities,omitempty"`
	Audio                        *Audio                        `json:"audio,omitempty"`
	Document                     *Document                     `json:"document,omitempty"`
	Animation                    *Animation                    `json:"animation,omitempty"`
	Game                         *Game                         `json:"game,omitempty"`
	Photo                        []PhotoSize                   `json:"photo,omitempty"`
	Sticker                      *Sticker                      `json:"sticker,omitempty"`
	Video                        *Video                        `json:"video,omitempty"`
	Voice                        *Voice                        `json:"voice,omitempty"`
	VideoNote                    *VideoNote                    `json:"video_note,omitempty"`
	Caption                      string                        `json:"caption,omitempty"`
	Contact                      *Contact                      `json:"contact,omitempty"`
	Dice                         *Dice                         `json:"dice,omitempty"`
	Location                     *Location                     `json:"location,omitempty"`
	Venue                        *Venue                        `json:"venue,omitempty"`
	Poll                         *Poll                         `json:"poll,omitempty"`
	NewChatMembers               *[]User                       `json:"new_chat_members,omitempty"`
	LeftChatMember               *User                         `json:"left_chat_member,omitempty"`
	NewChatTitle                 string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                 []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto              bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated             bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated        bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated           bool                          `json:"channel_chat_created,omitempty"`
	MigrateToChatId              int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId            int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                *Message                      `json:"pinned_message,omitempty"`
	Invoice                      *Invoice                      `json:"invoice,omitempty"`
	SuccessfulPayment            *SuccessfulPayment            `json:"successful_payment,omitempty"`
	ConnectedWebsite             string                        `json:"connected_website,omitempty"`
	PassportData                 *PassportData                 `json:"passport_data,omitempty"`
	ReplyMarkup                  *InlineKeyboardMarkup         `json:"reply_markup,omitempty"`
	VoiceChatScheduled           *VoiceChatScheduled           `json:"voice_chat_scheduled,omitempty"`
	VoiceChatStarted             *VoiceChatStarted             `json:"voice_chat_started,omitempty"`
	VoiceChatParticipantsInvited *VoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`
}

func (m *Message) IsPrivate() bool {
	return m.Chat.Type == chatTypePrivate
}

func (m *Message) IsGroup() bool {
	return m.Chat.Type == chatTypeGroup
}

func (m *Message) IsSuperGroup() bool {
	return m.Chat.Type == chatTypeSuperGroup
}

func (m *Message) IsChannel() bool {
	return m.Chat.Type == chatTypeChannel
}

func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	entity := m.Entities[0]

	return entity.Offset == 0 && entity.IsBotCommand()
}

func (m *Message) CommandInput() string {
	if !m.IsCommand() {
		return ""
	}

	entity := m.Entities[0]

	if int(entity.Length+1) > len(m.Text) {
		return ""
	}

	return m.Text[entity.Length+1:]
}
