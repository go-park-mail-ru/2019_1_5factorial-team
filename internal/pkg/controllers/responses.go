package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
