package common_models

type SendMessageRequestModel struct {
	ToPersonID int64  `json:"toPerson"`
	AudioBytes []byte `json:"audio"`
	//Message    string `json:"message"`
}

func NewSendMessageRequestModel() *SendMessageRequestModel {
	return new(SendMessageRequestModel)
}

type MessageDBModel struct {
	MessageID    int64  `redis:"messageID"`
	Timestamp    string `redis:"timestamp"`
	MessageAudio string `redis:"audioString"`
	Reaction     string `redis:"reaction"`
}

type ChatDBModel struct {
	ChatID      int64 `redis:"chatID"`
	ForPersonID int64 `redis:"forPersonID"`
}
