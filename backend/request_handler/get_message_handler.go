package request_handler

import (
	"backend/model/auth_token_data"
	db "backend/redisdb"
	"backend/utilities/appjson"
	r2 "backend/utilities/s3utils"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetUsers(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	_, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	res, err := db.GetAllUsers()
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

func GetUser(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	_, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	var (
		personIDStr string
	)
	personIDRaw := chi.URLParam(r, "userID")
	var err error
	if personIDStr, err = url.QueryUnescape(personIDRaw); err != nil {
		personIDStr = personIDRaw
	}
	personID, err := strconv.ParseInt(personIDStr, 10, 64)
	if err != nil {
		return err, 400
	}

	res, err := db.GetUser(personID)
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
	otherPersonIDRaw := chi.URLParam(r, "otherPersonID")
	var err error
	if otherPersonIDStr, err = url.QueryUnescape(otherPersonIDRaw); err != nil {
		otherPersonIDStr = otherPersonIDRaw
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
