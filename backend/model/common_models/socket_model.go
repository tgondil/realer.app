package common_models

type SocketAndResponseModel struct {
	ToPersonID             int64  `json:"toPerson"`
	FromPersonID           int64  `json:"fromPerson"`
	MessageID              int64  `json:"messageID"`
	Message                string `json:"message"`
	MessageFile            string `json:"messageFile,omitempty"`
	MessageFileSizeInBytes int64  `json:"messageFileSizeInBytes,omitempty"`
	MessageTime            int64  `json:"messageTime,omitempty"`
	Reaction               string `json:"reaction,omitempty"`
}
