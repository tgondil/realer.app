package auth_token_data

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Model struct {
	PersonID int64  `json:"personId"`
	Token    string `json:"token"`
}

func AuthTokenDataFromJWTPayload(payload map[string]any, token jwt.Token) (m Model, e error) {
	var personIDAny any
	var personIDFloat64 float64
	var ok bool
	if personIDAny, ok = payload["personId"]; !ok {
		e = errors.New("missing personId")
		return
	}
	if personIDFloat64, ok = personIDAny.(float64); !ok {
		e = errors.New("invalid personId")
		return
	}
	m = Model{
		PersonID: int64(personIDFloat64),
		Token:    token.Raw,
	}
	e = nil
	return
}
