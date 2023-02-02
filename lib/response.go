package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}) error {

	var res Response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	res.Status = "success"
	res.Code = code
	res.Message = getStatus(code)
	res.Data = payload

	return json.NewEncoder(w).Encode(res)

}

func ResponseError(w http.ResponseWriter, code int, message string) error {

	var res Response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	res.Status = "error"
	res.Code = code
	res.Message = message

	return json.NewEncoder(w).Encode(res)
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
