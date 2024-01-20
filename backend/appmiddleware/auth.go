package appmiddleware

import (
	"backend/model/auth_token_data"
	"backend/model/login"
	"backend/utilities/timeutils"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

func GenerateAuthToken(l *login.Model) (string, error) {
	claims := jwt.MapClaims{}

	claims["exp"] = timeutils.EndOfDayTime(time.Now()).Unix() //expiry

	claims["pld"] = map[string]any{ //payload
		"personId": l.PersonID,
	}
	claims["iss"] = "Boilermake" //issuer
	claims["aud"] = "*"          //audience

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	if signedToken, err := token.SignedString([]byte("boilermake")); err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	const authRejectionMessage = "Authorization rejected"
	const statusUnauthorized = 401
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}

		t, _ := strings.CutPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(
			strings.TrimSpace(t),
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				}
				return []byte("boilermake"), nil
			},
		)

		if err != nil {
			log.Println("auth middleware err", err)
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}

		if issuer, err := claims.GetIssuer(); err != nil || issuer != "Boilermake" {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}

		payload, isPayloadMap := claims["pld"].(map[string]any)
		if !isPayloadMap {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}
		_, containsPersonId := payload["personId"]
		_, containsSubsId := payload["subsId"]
		isActive, containsIsActive := payload["isActive"]
		isDoctorStaff, containsIsDoctorStaff := payload["isDoctorStaff"]

		ok = ok && containsPersonId && containsSubsId && containsIsActive && containsIsDoctorStaff

		if !(ok && token.Valid) {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(authRejectionMessage))
			return
		}

		isActiveBool, ok := isActive.(bool)
		if !ok || !isActiveBool {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte("Invalid login: inactive user"))
			return
		}

		isDoctorStaffBool, ok := isDoctorStaff.(bool)
		if !ok || !isDoctorStaffBool {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte("Invalid login: Not a user"))
			return
		}

		if expiry, err := claims.GetExpirationTime(); err != nil {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte("Invalid expiration"))
			return
		} else if time.Now().After(expiry.Time) {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte("Invalid expiration"))
			return

		}

		var authToken auth_token_data.Model
		authToken, err = auth_token_data.AuthTokenDataFromJWTPayload(payload, *token)
		if err != nil {
			w.WriteHeader(statusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		newServerContext := context.WithValue(r.Context(), "user", authToken)
		r = r.WithContext(newServerContext)
		next.ServeHTTP(w, r)
	})
}
