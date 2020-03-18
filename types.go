package ted

import (
	"strings"
)

type Update struct {
	ID            int            `json:"update_id"`
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type Message struct {
	// Unique message identifier inside this chat
	ID int `json:"message_id"`

	// Sender, empty for messages sent to channels
	From *User `json:"from"`

	// Conversation the message belongs to
	Chat Chat `json:"chat"`

	// For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message"`

	// For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Text string `json:"text"`

	// For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`
}

// CommandAndArgs extracts and returns a Telegram bot command and the rest of
// the message excluding the command. The command will have its leading slash
// and possible bot mention removed. If a command was not present in the
// message or the message did not begin with a command, command will be an
// empty string and args will contain the entire message text.
func (m Message) CommandAndArgs() (string, string) {
	for _, e := range m.Entities {
		if e.Type == "bot_command" && e.Offset == 0 {
			command := strings.TrimPrefix(m.Text[:e.Length], "/")
			args := strings.TrimSpace(m.Text[e.Length:])
			mention := strings.Index(command, "@")
			if mention != -1 {
				return command[:mention], args
			}
			return command, args
		}
	}
	return "", m.Text
}

type User struct {
	ID        int    `json:"ID"`
	FirstName string `json:"first_name"`
}

type Chat struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	Data            string   `json:"data"`
}
