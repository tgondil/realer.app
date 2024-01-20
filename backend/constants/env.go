package constants

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

type envConfig struct {
	Port           int
	SendGridApiKey string
}

var (
	LogInfo             bool
	IsDebug             bool
	TokenSigningPrivKey *rsa.PrivateKey
	TokenSigningPubKey  *rsa.PublicKey
)

var (
	DBInitialised     = false
	SocketInitialised = false
)

func InitEnv(debug, prefork, logInfo bool) {
	IsDebug = debug
	LogInfo = logInfo
	pemFile, err := os.ReadFile("./private_key.pem")
	if err != nil {
		panic(err)
	}
	if TokenSigningPrivKey, err = jwt.ParseRSAPrivateKeyFromPEM(pemFile); err != nil {
		panic(err)
	}
	pemFile, err = os.ReadFile("./public_key.pem")
	if err != nil {
		panic(err)
	}
	if TokenSigningPubKey, err = jwt.ParseRSAPublicKeyFromPEM(pemFile); err != nil {
		panic(err)
	}

	f, fOpenErr := os.Open("server.config")
	if fOpenErr != nil {
		log.Panicln("error opening server.config file:", fOpenErr)
		return
	}
	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Println("error closing cron.config file:", err)
		}
	}(f)
}
