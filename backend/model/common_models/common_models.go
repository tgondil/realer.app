package common_models

import "backend/utilities/appjson"

type SocketAndResponseModel struct {
	ToPersonID             int64             `json:"toPerson"`
	FromPersonID           int64             `json:"fromPerson"`
	MessageID              int64             `json:"messageID"`
	Message                string            `json:"message"`
	MessageFile            string            `json:"messageFile,omitempty"`
	MessageFileSizeInBytes int64             `json:"messageFileSizeInBytes,omitempty"`
	MessageTime            string            `json:"messageTime,omitempty"`
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
	MessageID            int64             `json:"messageID"`
	FromPersonID         int64             `json:"fromPersonID"`
	Timestamp            string            `json:"timestamp,omitempty"`
	MessageText          string            `json:"content,omitempty"`
	TextReaction         string            `json:"textReaction,omitempty"`
	MessageAudio         string            `json:"audioPath,omitempty"`
	AudioLengthInSeconds int64             `json:"audioLengthInSeconds,omitempty"`
	AudioReaction        []ReactionDBModel `json:"audioReaction,omitempty"`
}

type ChatDBModel struct {
	ChatID     int64  `redis:"chatID"`
	PersonID   int64  `redis:"personID"`
	PersonName string `redis:"personName"`
}

func (c *ChatDBModel) MarshalBinary() (data []byte, err error) {
	data, err = appjson.Marshal(struct {
		ChatID     int64  `json:"chatID"`
		PersonID   int64  `json:"personID"`
		PersonName string `json:"personName"`
	}{
		ChatID:     c.ChatID,
		PersonID:   c.PersonID,
		PersonName: c.PersonName,
	})
	return data, err
}

type PersonDBModel struct {
	PersonID   int64  `json:"personID"`
	PersonName string `json:"personName"`
}

type ReactionDBModel struct {
	Timestamp int64  `json:"timestamp"`
	Reaction  string `json:"reaction"`
}
