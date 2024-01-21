package request_handler

import (
	socket "backend/appsocket"
	"backend/model/auth_token_data"
	"backend/model/common_models"
	"backend/redisdb"
	"backend/utilities/appjson"
	s3 "backend/utilities/s3utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var filePool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func SendMessageWithFile(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	ctx := r.Context()
	authTokenData, containsAuthToken := ctx.Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}

	err := r.ParseMultipartForm(55 << 20)
	if err != nil {
		return err, 400
	}

	var (
		form = r.MultipartForm
	)

	if len(form.File) == 0 {
		return errors.New("No files"), 400
	}

	var (
		files                       []*multipart.FileHeader
		tempSlice                   []string
		toPersonID                  int64
		fileNameString, contentType string
		file                        multipart.File
		ok                          bool
		fileLengthInSeconds         int64
		fileSizeInBytes             int64
		nowTime                     time.Time
	)

	if files, ok = form.File["files"]; !ok || len(files) == 0 {
		return errors.New("No files"), 400
	}
	err = nil

	pathBuilder := strings.Builder{}
	pathBuilder.Grow(35)
	pathBuilder.WriteString("Chats/AudioFiles/")
	if tempSlice, ok = form.Value["toPersonID"]; ok && len(tempSlice) > 0 {
		if tmpInt, parseErr := strconv.ParseInt(tempSlice[0], 10, 64); err != nil {
			return parseErr, 400
		} else {
			toPersonID = tmpInt
		}
	}
	if tempSlice, ok = form.Value["audioLength"]; ok && len(tempSlice) > 0 {
		if tmpInt, parseErr := strconv.ParseInt(tempSlice[0], 10, 64); err != nil {
			return parseErr, 400
		} else {
			fileLengthInSeconds = tmpInt
		}
	}

	fileBytes := filePool.Get().(*bytes.Buffer)
	defer func() {
		fileBytes.Reset()
		filePool.Put(fileBytes)
	}()
	formFile := files[0]

	fileBytes.Reset()
	fileSizeInBytes = formFile.Size >> 10
	nowTime = time.Now()
	pathBuilder.WriteString(fmt.Sprint(nowTime.UnixMilli()))
	pathBuilder.WriteString("_")
	pathBuilder.WriteString(formFile.Filename)
	fileNameString = pathBuilder.String()
	contentType = s3.GetContentType(formFile.Filename)

	if file, err = formFile.Open(); err != nil {
		_ = file.Close()
		return err, 400
	} else if _, err = io.Copy(fileBytes, file); err != nil {
		_ = file.Close()
		return err, 400
	}
	_ = file.Close()
	if err = s3.UploadBytes(fileBytes.Bytes(), fileNameString, contentType); err != nil {
		return err, 400
	}

	redisM := common_models.MessageDBModel{
		FromPersonID:         authTokenData.PersonID,
		MessageID:            0,
		Timestamp:            nowTime.Format(time.DateTime),
		MessageAudio:         fileNameString,
		AudioLengthInSeconds: fileLengthInSeconds,
		TextReaction:         "",
	}
	if !redisdb.ChatExists(authTokenData.PersonID, toPersonID) {
		if _, err = redisdb.CreateChat(authTokenData.PersonID, toPersonID, nowTime.UnixMilli()); err != nil {
			return err, 400
		}
	}
	err = redisdb.AddMessage(authTokenData.PersonID, toPersonID, &redisM)
	if err != nil {
		return err, 400
	}
	m := common_models.SocketAndResponseModel{
		FromPersonID:           authTokenData.PersonID,
		ToPersonID:             toPersonID,
		MessageID:              redisM.MessageID,
		MessageFile:            fileNameString,
		MessageFileSizeInBytes: fileSizeInBytes,
		MessageTime:            nowTime.Format(time.DateTime),
	}
	socket.Broadcast([]int64{toPersonID}, "new_message", m)
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
		MessageID  int64 `json:"messageID"`
		ToPersonID int64 `json:"toPersonID"`
		Reaction   []common_models.ReactionDBModel
	})

	err := appjson.UnmarshalRequestBody(r.Body, body)
	if err != nil {
		return err, 400
	}

	if err = redisdb.AddReactionToAudio(authTokenData.PersonID, body.ToPersonID, body.MessageID, body.Reaction); err != nil {
		return err, 400
	}
	socket.Broadcast([]int64{body.ToPersonID}, "audio_reaction", common_models.SocketAndResponseModel{
		MessageID:     body.MessageID,
		AudioReaction: body.Reaction,
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
