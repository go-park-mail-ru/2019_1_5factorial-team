package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorMessage struct {
	Error string `json:"error"`
}

func addOkHeader(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
}

func addBody(res http.ResponseWriter, bodyMessage interface{}) {
	marshalBody, err := json.Marshal(bodyMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	res.Write(marshalBody)
}

func OkResponse(res http.ResponseWriter, bodyMessage interface{}) {
	addOkHeader(res)
	addBody(res, bodyMessage)
}

func addErrHeader(res http.ResponseWriter, errCode int) {
	res.WriteHeader(errCode)
}

func addErrBody(res http.ResponseWriter, errMsg string) {
	addBody(res, errorMessage{Error: errMsg})
}

func ErrResponse(res http.ResponseWriter, errCode int, errMsg string) {
	addErrHeader(res, errCode)
	addErrBody(res, errMsg)
}
