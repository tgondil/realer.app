package request_handler

import "net/http"

func AllUsers(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	r.URL.Query().Get("name")
	return nil, 400
}
