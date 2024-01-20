package stringutils

import (
	"regexp"
	str "strings"
)

//const (
//	SomethingWentWrong = "Something went wrong"
//	InvalidQuery       = "Invalid request"
//	InvalidPayload     = "Invalid payload"
//	MissingPayload     = "Missing payload"
//	InvalidRequest     = "Invalid request"
//	NoRowsAffected     = "No rows affected"
//)

//func CountryCodeWithoutPlusSign(s string) string {
//	return str.TrimPrefix(s, "+")
//}

func SuccessResponse(body any) map[string]any {
	return map[string]any{
		"success": true,
		"message": "Success",
		"data":    body,
	}
}

func ValidSQLQuery(query string) bool {
	r1 := regexp.MustCompile("'(.*?)'")
	r2 := regexp.MustCompile(`"(.*?)"`)
	r3 := regexp.MustCompile("`(.*?)`")
	query = r1.ReplaceAllLiteralString(query, "")
	query = r2.ReplaceAllLiteralString(query, "")
	query = r3.ReplaceAllLiteralString(query, "")
	if str.HasSuffix(query, ";") {
		return !str.Contains(str.TrimSuffix(query, ";"), ";")
	} else {
		return !str.Contains(query, ";")
	}
}

func HandleQuotesForSQLPtr(value *string) string {
	if value == nil {
		return ""
	}
	return str.ReplaceAll(*value, "'", "\\'")
}

func HandleQuotesForSQLString(value string) string {
	return str.ReplaceAll(value, "'", "\\'")
}

func HandleQuotesForSQLBytes(value []byte) []byte {
	if value == nil {
		return []byte{}
	}
	res := make([]byte, 0, int(float32(len(value))*1.1))
	c := "'"[0]
	escape := `\`[0]
	for i := 0; i < len(value); i++ {
		if value[i] != c {
			res = append(res, value[i])
		} else {
			res = append(res, escape, value[i])
		}
	}
	return res
}
