package appjson

import (
	j "github.com/json-iterator/go"
	"io"
	"log"
	"reflect"
)

type RawMessage []byte

var encDec = j.Config{
	//config Default with DisallowUnknownFields as true
	EscapeHTML:            true,
	DisallowUnknownFields: true,
}.Froze()

// Closes body
func UnmarshalRequestBody(r io.ReadCloser, v any) error {
	defer func(r io.ReadCloser) {
		err := r.Close()
		if err != nil {
			log.Println("Error in closing request body: ", err)
		}
	}(r)
	decoder := encDec.NewDecoder(r)
	e := decoder.Decode(v)
	if e != nil {
		log.Println("Error in unmarshalling request body: ", e)
		log.Println("Type: ", reflect.TypeOf(v))
	}
	return e
}

func NewDecoder(r io.Reader) *j.Decoder {
	return encDec.NewDecoder(r)
}

func NewEncoder(w io.Writer) *j.Encoder {
	return encDec.NewEncoder(w)
}

func Unmarshal(data []byte, v any) error {
	e := encDec.Unmarshal(data, v)
	if e != nil {
		log.Println("Error in unmarshalling: ", e)
		log.Println("Data: ", string(data))
		log.Println("Type: ", reflect.TypeOf(v))
	}
	return e
}

//func UnmarshalString(data string, v interface{}) error {
//	e := encDec.UnmarshalFromString(data, v)
//	if e != nil {
//		log.Println("Error in unmarshalling from string: ", e)
//		log.Println("Data: ", string(data))
//		log.Println("Type: ", reflect.TypeOf(v))
//	}
//	return e
//}

func Marshal(v any) ([]byte, error) {
	return encDec.Marshal(v)
}
