package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorMessage struct {
	Error string `json:"error"`
}

func AddOkHeader(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
}

func AddBody(res http.ResponseWriter, bodyMessage interface{}) {
	marshalBody, err := json.Marshal(bodyMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	res.Write(marshalBody)
}

func AddErrHeader(res http.ResponseWriter, errCode int) {
	res.WriteHeader(errCode)
}

func AddErrBody(res http.ResponseWriter, errMsg string) {
	AddBody(res, errorMessage{errMsg})
}