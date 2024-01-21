package request_handler

import (
	socket "backend/appsocket"
	"backend/model/auth_token_data"
	"backend/model/common_models"
	"backend/redisdb"
	"backend/utilities/appjson"
	s3 "backend/utilities/s3utils"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func SendMessageWithFile(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	ctx := r.Context()
	authTokenData, containsAuthToken := ctx.Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}

	body := common_models.NewSendMessageRequestModel()
	err := appjson.UnmarshalRequestBody(r.Body, body)
	if err != nil {
		return err, 400
	}
	if body.AudioBytes == nil {
		return errors.New("Invalid audio"), 400
	}

	pathBuilder := strings.Builder{}
	pathBuilder.Grow(35)
	pathBuilder.WriteString("Chats/AudioFiles/")
	pathBuilder.WriteString(fmt.Sprint(body.ToPersonID))
	nowTime := time.Now()
	pathBuilder.WriteString(fmt.Sprint(nowTime.UnixMilli()))
	pathBuilder.WriteString(".webm")
	fileNameString := pathBuilder.String()
	contentType := "application/octet-stream"

	if err = s3.UploadBytes(body.AudioBytes, fileNameString, contentType); err != nil {
		return err, 400
	}

	redisM := common_models.MessageDBModel{
		FromPersonID:         authTokenData.PersonID,
		MessageID:            0,
		Timestamp:            nowTime.Format(time.DateTime),
		MessageAudio:         fileNameString,
		AudioLengthInSeconds: *body.AudioLength,
		TextReaction:         "",
	}
	if !redisdb.ChatExists(authTokenData.PersonID, body.ToPersonID) {
		if _, err = redisdb.CreateChat(authTokenData.PersonID, body.ToPersonID, nowTime.UnixMilli()); err != nil {
			return err, 400
		}
	}
	err = redisdb.AddMessage(authTokenData.PersonID, body.ToPersonID, &redisM)
	if err != nil {
		return err, 400
	}
	m := common_models.SocketAndResponseModel{
		FromPersonID:           authTokenData.PersonID,
		ToPersonID:             body.ToPersonID,
		MessageID:              redisM.MessageID,
		MessageFile:            fileNameString,
		MessageFileSizeInBytes: int64(len(body.AudioBytes) >> 10),
		MessageTime:            nowTime.Format(time.DateTime),
	}
	socket.Broadcast([]int64{body.ToPersonID}, "new_message", m)
	b, err := appjson.Marshal(m)
	if err != nil {
		return err, 400
	}
	_, err = w.Write(b)
	return err, 400
}

func SendMessage(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	ctx := r.Context()
	authTokenData, containsAuthToken := ctx.Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}

	body := common_models.NewSendMessageRequestModel()
	err := appjson.UnmarshalRequestBody(r.Body, body)
	if err != nil {
		return err, 400
	}

	var nowTime = time.Now()

	redisM := common_models.MessageDBModel{
		FromPersonID: authTokenData.PersonID,
		MessageID:    0,
		Timestamp:    nowTime.Format(time.DateTime),
		MessageText:  body.Message,
		TextReaction: "",
	}
	if !redisdb.ChatExists(authTokenData.PersonID, body.ToPersonID) {
		if _, err = redisdb.CreateChat(authTokenData.PersonID, body.ToPersonID, nowTime.UnixMilli()); err != nil {
			return err, 400
		}
	}
	err = redisdb.AddMessage(authTokenData.PersonID, body.ToPersonID, &redisM)
	if err != nil {
		return err, 400
	}
	m := common_models.SocketAndResponseModel{
		FromPersonID: authTokenData.PersonID,
		ToPersonID:   body.ToPersonID,
		MessageID:    redisM.MessageID,
		Message:      body.Message,
		MessageTime:  nowTime.Format(time.DateTime),
	}
	socket.Broadcast([]int64{body.ToPersonID}, "new_message", m)
	b, err := appjson.Marshal(m)
	if err != nil {
		return err, 400
	}
	_, err = w.Write(b)
	return err, 400
}

func AddReactionToAudio(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	authTokenData, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}

	body := new(struct {
		MessageID  int64                           `json:"messageID"`
		ToPersonID int64                           `json:"toPersonID"`
		Reactions  []common_models.ReactionDBModel `json:"reactions"`
	})

	err := appjson.UnmarshalRequestBody(r.Body, body)
	if err != nil {
		return err, 400
	}

	if err = redisdb.AddReactionToAudio(authTokenData.PersonID, body.ToPersonID, body.MessageID, body.Reactions); err != nil {
		return err, 400
	}
	socket.Broadcast([]int64{body.ToPersonID}, "audio_reaction", common_models.SocketAndResponseModel{
		MessageID:     body.MessageID,
		AudioReaction: body.Reactions,
	})
	_, err = w.Write([]byte("success"))
	return err, 400
}

func AddReactionToText(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	messageIDRaw := chi.URLParam(r, "messageID")
	queryParams := r.URL.Query()

	reaction := queryParams.Get("reaction")
	if reaction == "" {
		return errors.New("Invalid reaction"), 400
	}
	toPersonIDStr := queryParams.Get("toPersonID")
	if toPersonIDStr == "" {
		return errors.New("Invalid toPersonID"), 400
	}
	var err error
	var toPersonID int64
	if toPersonID, err = strconv.ParseInt(toPersonIDStr, 10, 64); err != nil {
		return err, 400
	}

	var messageIDString string
	if messageIDString, err = url.QueryUnescape(messageIDRaw); err != nil {
		messageIDString = messageIDRaw
	}
	var messageID int64
	if messageID, err = strconv.ParseInt(messageIDString, 10, 64); err != nil {
		return err, 400
	}
	authTokenData, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	if err = redisdb.AddReactionToText(authTokenData.PersonID, toPersonID, messageID, reaction); err != nil {
		return err, 400
	}
	socket.Broadcast([]int64{toPersonID}, "text_reaction", common_models.SocketAndResponseModel{
		MessageID:    messageID,
		TextReaction: reaction,
	})
	_, err = w.Write([]byte("success"))
	return err, 400

}
