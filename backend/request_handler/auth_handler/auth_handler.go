package auth_handler

import (
	auth_common "backend/appmiddleware"
	login_master "backend/model/login"
	db "backend/redisdb"
	"backend/utilities/appjson"
	"errors"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) (error, int) {
	const prefix = "Login"

	body := make(map[string]any, 5)

	err := appjson.UnmarshalRequestBody(r.Body, &body)
	if err != nil {
		log.Println(prefix, err)
		return err, 400
	}
	un, containsUserName := body["username"]
	pwd, containsPassword := body["password"]
	if !containsUserName || !containsPassword {
		return errors.New("Invalid"), 400
	}
	userName, ok := un.(string)
	if !ok {
		return errors.New("Invalid username"), 400
	}
	password, ok := pwd.(string)
	if !ok {
		return errors.New("Invalid password"), 400
	}

	var personID int64
	personID, err = db.Login(userName, password)
	if err != nil {
		return (err), 400
	}

	var token string
	loginMaster := login_master.Model{PersonID: personID}
	token, err = auth_common.GenerateAuthToken(&loginMaster)
	if err != nil {
		log.Println(prefix, err)
		return (err), 400
	}
	var jsonBytes []byte
	jsonBytes, err = appjson.Marshal(loginMaster.ResponseWithToken(token))
	if err != nil {
		log.Println(prefix, err)
		return (err), 400
	}
	_, err = w.Write(jsonBytes)
	return (err), 400
}

func Signup(w http.ResponseWriter, r *http.Request) (error, int) {
	const prefix = "Signup"

	body := make(map[string]any, 5)

	err := appjson.UnmarshalRequestBody(r.Body, &body)
	if err != nil {
		log.Println(prefix, err)
		return err, 400
	}
	un, containsUserName := body["username"]
	pwd, containsPassword := body["password"]
	if !containsUserName || !containsPassword {
		return errors.New("Invalid"), 400
	}
	userName, ok := un.(string)
	if !ok {
		return errors.New("Invalid username"), 400
	}
	password, ok := pwd.(string)
	if !ok {
		return errors.New("Invalid password"), 400
	}

	var personID int64
	personID, err = db.CreatePerson(userName, password)
	if err != nil {
		return (err), 400
	}

	var token string
	loginMaster := login_master.Model{PersonID: personID}
	token, err = auth_common.GenerateAuthToken(&loginMaster)
	if err != nil {
		log.Println(prefix, err)
		return (err), 400
	}
	var jsonBytes []byte
	jsonBytes, err = appjson.Marshal(loginMaster.ResponseWithToken(token))
	if err != nil {
		log.Println(prefix, err)
		return (err), 400
	}
	_, err = w.Write(jsonBytes)
	return (err), 400
}
