package request_handler

import (
	"backend/model/auth_token_data"
	db "backend/redisdb"
	"backend/utilities/appjson"
	r2 "backend/utilities/cloudflareR2utils"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetFiles(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	const prefix = "GetFiles"

	var (
		b        []byte
		filePath string
	)
	filePathRaw := chi.URLParam(r, "*")
	var err error
	if filePath, err = url.QueryUnescape(filePathRaw); err != nil {
		filePath = filePathRaw
	}
	filePathLen := len(filePath)
	if len(filePath) == 0 || filePath[filePathLen-1] == "/"[0] {
		return errors.New("Invalid"), 400
	}
	const bucketPrefix = "Chats/AudioFiles/"

	builder := new(strings.Builder)
	builder.Grow(len(bucketPrefix) + len(filePath))
	_, _ = builder.WriteString(bucketPrefix)
	_, _ = builder.WriteString(filePath)
	b, err = r2.GetFile(builder.String())
	if err != nil {
		log.Println(prefix, err)
		return err, 400
	}
	_, err = w.Write(b)
	return err, 400
}

func GetSingleChatMessages(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	authToken, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	var (
		otherPersonIDStr string
	)
	chatIDRaw := chi.URLParam(r, "otherPersonID")
	var err error
	if otherPersonIDStr, err = url.QueryUnescape(chatIDRaw); err != nil {
		otherPersonIDStr = chatIDRaw
	}
	otherPersonID, err := strconv.ParseInt(otherPersonIDStr, 10, 64)
	if err != nil {
		return err, 400
	}
	res, err := db.GetAllMessages(authToken.PersonID, otherPersonID)
	if err != nil {
		return err, 400
	}
	var b []byte
	if b, err = appjson.Marshal(res); err != nil {
		log.Println(err)
		return err, 400
	}
	_, err = w.Write(b)
	return err, 400

}

func GetChats(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	authToken, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	res, err := db.GetAllChats(authToken.PersonID)

	if err != nil {
		log.Println(err)
		return err, 400
	}

	var b []byte
	if b, err = appjson.Marshal(res); err != nil {
		log.Println(err)
		return err, 400
	}
	_, err = w.Write(b)
	return err, 400
}
