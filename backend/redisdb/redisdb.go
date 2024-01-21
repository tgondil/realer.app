package redisdb

import (
	"backend/model/common_models"
	"backend/utilities/appjson"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var client *redis.Client
var ctx = context.Background()

func Init() {
	opt, err := redis.ParseURL("redis://default:boilermake2024@redis-12496.c323.us-east-1-2.ec2.cloud.redislabs.com:12496/0")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	client = redis.NewClient(opt)
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
	cmd = client.HSet(ctx, fmt.Sprintf("person:%d", personID), "password", password)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	return personID, nil
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

func CreateChat(person1, person2 int64) (chatID int64, err error) {
	cmd := client.Incr(ctx, "chatIDCount")
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	chatID, e2 := cmd.Result()
	if e2 != nil {
		return 0, e2
	}
	var chatModel1, chatModel2 common_models.ChatDBModel
	chatModel1.ChatID = chatID
	chatModel1.ForPersonID = person1
	chatModel2.ChatID = chatID
	chatModel2.ForPersonID = person2

	cmd = client.HSet(ctx, fmt.Sprintf("chats:%d", person1), chatModel1)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	cmd = client.HSet(ctx, fmt.Sprintf("chats:%d", person2), chatModel2)
	if e1 := cmd.Err(); e1 != nil {
		return 0, e1
	}
	return chatID, nil
}

func AddMessage(fromPersonID, toPersonID int64, message *common_models.MessageDBModel) error {
	// auto increment messageID
	cmd := client.Incr(ctx, "messageIDCount")
	e1 := cmd.Err()
	if e1 != nil {
		return e1
	}
	message.MessageID, e1 = cmd.Result()
	if e1 != nil {
		return e1
	}
	return client.HSet(ctx, fmt.Sprintf("messages:%d_%d", fromPersonID, toPersonID), message.MessageID, message).Err()
}

func AddReaction(fromPersonID, toPersonID, messageID int64, reaction string) error {
	return client.HSet(ctx, fmt.Sprintf("messages:%d_%d", fromPersonID, toPersonID), messageID, reaction).Err()
}

func Close() error {
	return client.Close()
}
