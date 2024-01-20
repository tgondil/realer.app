package encryptionutils

import (
	"backend/model/auth_token_data"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"golang.org/x/text/encoding"
	"hash"
	"io"
	"unicode/utf8"
)

func DecryptAPIBody(body []byte, authTokenData *auth_token_data.Model) ([]byte, error) {
	//privateKey := GetPrivateKey(authTokenData)
	//key := []byte(privateKey)
	//
	//if base64Bytes, err := isBase64(body); err == nil {
	//	decrypted := bytes.NewBuffer(nil)
	//	decrypted.Grow(len(base64Bytes))
	//	_, err = sio.Decrypt(decrypted, bytes.NewBuffer(base64Bytes), sio.Config{
	//		Key: key,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	return decrypted.Bytes(), nil
	//}

	return body, nil
}

//func GetPrivateKey(authTokenData *auth_token_data.Model) string {
//	token := authTokenData.Token
//	key := token[:16] + token[len(token)-16:]
//	return key
//}
//
//func GetIV(authTokenData *auth_token_data.Model) io.Writer {
//	ivStr := fmt.Sprintf("%d%dBES", authTokenData.SubsID, authTokenData.PersonID)
//	hexString := strings.Builder{}
//	for _, c := range []byte(ivStr) {
//		hexString.WriteString(strings.ToUpper(hex.EncodeToString([]byte{c})))
//	}
//	decoded, err := hex.DecodeString(hexString.String())
//	if err != nil {
//		panic(err) // Handle the error appropriately
//	}
//	return bytes.NewBuffer(decoded)
//}
//
//func isBase64(s []byte) ([]byte, error) {
//	b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
//	_, err := base64.StdEncoding.Decode(s, b)
//	return b, err
//}

func EncryptPassword(password *string) (string, error) {
	if password == nil {
		return "", errors.New("Null pointer to password")
	}
	if len(*password) == 0 {
		return "", errors.New("Empty password")
	}
	if utf8.ValidString(*password) {
		hasher := sha1.New()
		hasher.Write([]byte((*password)))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	} else {
		e := encoding.Encoder{}
		strRes, err := e.String(*password)
		if err != nil {
			return "", err
		}
		return strRes, nil
	}
}

func EncryptOAEP(hash hash.Hash, random io.Reader, public *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := public.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, public, msg[start:finish], label)
		if err != nil {
			return nil, err
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return encryptedBytes, nil
}

func DecryptOAEP(hash hash.Hash, random io.Reader, private *rsa.PrivateKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := private.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, private, msg[start:finish], label)
		if err != nil {
			return nil, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, nil
}
