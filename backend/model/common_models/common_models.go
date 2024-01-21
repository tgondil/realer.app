package common_models

type SocketAndResponseModel struct {
	ToPersonID             int64             `json:"toPerson"`
	FromPersonID           int64             `json:"fromPerson"`
	MessageID              int64             `json:"messageID"`
	Message                string            `json:"message"`
	MessageFile            string            `json:"messageFile,omitempty"`
	MessageFileSizeInBytes int64             `json:"messageFileSizeInBytes,omitempty"`
	MessageTime            int64             `json:"messageTime,omitempty"`
	TextReaction           string            `json:"reaction,omitempty"`
	AudioReaction          []ReactionDBModel `json:"audioReaction,omitempty"`
}

type SendMessageRequestModel struct {
	ToPersonID int64  `json:"id"`
	AudioBytes []byte `json:"audio"`
	Message    string `json:"content"`
}

func NewSendMessageRequestModel() *SendMessageRequestModel {
	return new(SendMessageRequestModel)
}

type MessageDBModel struct {
	MessageID     int64             `json:"messageID"`
	Timestamp     string            `json:"timestamp"`
	MessageText   string            `json:"messageText"`
	TextReaction  string            `json:"textReaction"`
	MessageAudio  string            `json:"audioPath"`
	AudioReaction []ReactionDBModel `json:"audioReaction"`
}

type ChatDBModel struct {
	ChatID      int64 `json:"chatID"`
	ForPersonID int64 `json:"forPersonID"`
}

type ReactionDBModel struct {
	Timestamp int64  `json:"timestamp"`
	Reaction  string `json:"reaction"`
}
