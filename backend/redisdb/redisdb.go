package redisdb

import (
	"backend/model/common_models"
	"backend/utilities/appjson"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"slices"
	"strconv"
)

var client *redis.Client
var Close func() error
var ctx = context.Background()

func Init() {
	opt, err := redis.ParseURL("redis://default:boilermake2024@redis-12496.c323.us-east-1-2.ec2.cloud.redislabs.com:12496/0")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	client = redis.NewClient(opt)
	Close = client.Close
}

func CreatePerson(name, password string) (int64, error) {
	if PersonExists(name) {
		return 0, fmt.Errorf("person already exists")
	}
	cmd := client.Incr(ctx, "personIDCount")
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	personID, e2 := cmd.Result()
	if e2 != nil {
		return 0, e2
	}
	cmd = client.HSet(ctx, "persons", name, personID)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	cmd = client.HSet(ctx, fmt.Sprintf("person:%d", personID), "password", password, "name", name)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	return personID, nil
}

func GetPersonFromID(personID int64) (string, error) {
	return client.HGet(ctx, fmt.Sprintf("person:%d", personID), "name").Result()
}

func PersonExists(name string) bool {
	cmd := client.HGet(ctx, "persons", name)
	res, err := cmd.Result()
	if err != nil {
		return false
	}
	return res != ""
}

func Login(name, password string) (int64, error) {
	if !PersonExists(name) {
		return 0, fmt.Errorf("person does not exist")
	}
	personID, e1 := client.HGet(ctx, "persons", name).Int64()
	if e1 != nil {
		return 0, e1
	}
	actualPassword, e2 := client.HGet(ctx, fmt.Sprintf("person:%d", personID), "password").Result()
	if e2 != nil {
		return 0, e2
	}
	if actualPassword != password {
		return 0, fmt.Errorf("invalid password")
	}
	return personID, nil
}

func GetAllMessages(person1, person2 int64) ([]common_models.MessageDBModel, error) {
	m := make([]common_models.MessageDBModel, 0, 1024)
	var minPersonID, maxPersonID int64
	if person1 < person2 {
		minPersonID = person1
		maxPersonID = person2
	} else {
		minPersonID = person2
		maxPersonID = person1
	}
	r := client.HGetAll(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID))
	if e1 := r.Err(); e1 != nil {
		return nil, e1
	}
	for _, v := range r.Val() {
		newM := &common_models.MessageDBModel{}
		if err := appjson.Unmarshal([]byte(v), newM); err != nil {
			return nil, err
		}
		m = append(m, *newM)
	}
	slices.SortFunc(m, func(i, j common_models.MessageDBModel) int {
		return int(i.MessageID - j.MessageID)
	})
	return m, nil
}

func GetAllChats(personID int64) ([]common_models.ChatDBModel, error) {
	m := make([]common_models.ChatDBModel, 0, 1024)
	r := client.HGetAll(ctx, fmt.Sprintf("chats:%d", personID))
	if e1 := r.Err(); e1 != nil {
		return nil, e1
	}
	for _, v := range r.Val() {
		newM := &common_models.ChatDBModel{}
		if err := appjson.Unmarshal([]byte(v), newM); err != nil {
			return nil, err
		}
		m = append(m, *newM)
	}
	return m, nil
}

func GetAllUsers() ([]common_models.PersonDBModel, error) {
	cmd := client.HGetAll(ctx, "persons")
	if e1 := cmd.Err(); e1 != nil {
		return nil, e1
	}
	m := make([]common_models.PersonDBModel, 0, 1024)
	for k, v := range cmd.Val() {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		newM := common_models.PersonDBModel{
			PersonName: k,
			PersonID:   val,
		}
		m = append(m, newM)
	}
	return m, nil
}
func GetUser(personID int64) (*common_models.PersonDBModel, error) {
	cmd := client.HGet(ctx, fmt.Sprintf("person:%d", personID), "name")
	if e1 := cmd.Err(); e1 != nil {
		return nil, e1
	}

	newM := &common_models.PersonDBModel{
		PersonName: cmd.Val(),
		PersonID:   personID,
	}

	return newM, nil
}

func ChatExists(person1, person2 int64) bool {
	var minPersonID, maxPersonID int64
	if person1 < person2 {
		minPersonID = person1
		maxPersonID = person2
	} else {
		minPersonID = person2
		maxPersonID = person1
	}
	cmd := client.Exists(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID))
	res, err := cmd.Result()
	if err != nil {
		return false
	}
	return res != 0
}

func CreateChat(person1, person2, nowTimestamp int64) (chatID int64, err error) {
	var chatModel1, chatModel2 common_models.ChatDBModel
	chatModel1.PersonID = person1
	chatModel1.PersonName, err = GetPersonFromID(person1)
	if err != nil {
		return 0, err
	}
	chatModel2.PersonID = person2
	chatModel2.PersonName, err = GetPersonFromID(person2)
	if err != nil {
		return 0, err
	}

	cmd := client.Incr(ctx, "chatIDCount")
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	chatID = cmd.Val()
	chatModel1.ChatID = chatID
	chatModel2.ChatID = chatID
	cmd = client.HSet(ctx, fmt.Sprintf("chats:%d", person1), chatModel1.ChatID, chatModel1)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	cmd = client.HSet(ctx, fmt.Sprintf("chats:%d", person2), chatModel2.ChatID, chatModel2)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	return chatID, nil
}

func AddMessage(fromPersonID, toPersonID int64, message *common_models.MessageDBModel) error {
	var minPersonID, maxPersonID int64
	if fromPersonID < toPersonID {
		minPersonID = fromPersonID
		maxPersonID = toPersonID
	} else {
		minPersonID = toPersonID
		maxPersonID = fromPersonID
	}
	cmd := client.Incr(ctx, "messageIDCount")
	e1 := cmd.Err()
	if e1 != nil {
		return e1
	}
	message.MessageID = cmd.Val()
	return client.HSet(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID), message.MessageID, message).Err()
}

func AddReactionToText(fromPersonID, toPersonID, messageID int64, reaction string) error {
	var minPersonID, maxPersonID int64
	if fromPersonID < toPersonID {
		minPersonID = fromPersonID
		maxPersonID = toPersonID
	} else {
		minPersonID = toPersonID
		maxPersonID = fromPersonID
	}
	r := client.HGet(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID), fmt.Sprint(messageID))
	if e1 := r.Err(); e1 != nil {
		return e1
	}
	newM := &common_models.MessageDBModel{}
	if err := appjson.Unmarshal([]byte(r.Val()), newM); err != nil {
		return err
	}
	newM.TextReaction = reaction
	return client.HSet(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID), messageID, newM).Err()
}

func AddReactionToAudio(fromPersonID, toPersonID, messageID int64, reactions []common_models.ReactionDBModel) error {
	var minPersonID, maxPersonID int64
	if fromPersonID < toPersonID {
		minPersonID = fromPersonID
		maxPersonID = toPersonID
	} else {
		minPersonID = toPersonID
		maxPersonID = fromPersonID
	}
	r := client.HGet(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID), fmt.Sprint(messageID))
	if e1 := r.Err(); e1 != nil {
		return e1
	}
	newM := &common_models.MessageDBModel{}
	if err := appjson.Unmarshal([]byte(r.Val()), newM); err != nil {
		return err
	}
	newM.AudioReaction = reactions
	return client.HSet(ctx, fmt.Sprintf("messages:%d_%d", minPersonID, maxPersonID), messageID, newM).Err()
}
