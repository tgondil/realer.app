package common_models

type SendMessageRequestModel struct {
	ToPersonID int64  `json:"id"`
	AudioBytes []byte `json:"audio"`
	Message    string `json:"content"`
}

func NewSendMessageRequestModel() *SendMessageRequestModel {
	return new(SendMessageRequestModel)
}

type MessageDBModel struct {
	MessageID    int64  `json:"messageID"`
	Timestamp    string `json:"timestamp"`
	MessageAudio string `json:"audioString"`
	MessageText  string `json:"message"`
	Reaction     string `json:"reaction"`
}

//func (m *MessageDBModel) MarshalBinary() (data []byte, err error) {
//	return appjson.Marshal(struct {
//		MessageID    int64  `json:"messageID"`
//		Timestamp    string `json:"timestamp"`
//		MessageAudio string `json:"audioString"`
//		MessageText  string `json:"message"`
//		Reaction     string `json:"reaction"`
//	}{
//		MessageID:    m.MessageID,
//		Timestamp:    m.Timestamp,
//		MessageAudio: m.MessageAudio,
//		MessageText:  m.MessageText,
//		Reaction:     m.Reaction,
//	})
//
//}

type ChatDBModel struct {
	ChatID      int64 `redis:"chatID"`
	ForPersonID int64 `redis:"forPersonID"`
}
