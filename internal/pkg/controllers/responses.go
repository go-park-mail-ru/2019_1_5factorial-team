package controllers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func addOkHeader(res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
}

func addBody(res http.ResponseWriter, bodyMessage interface{}) {
	marshalBody, err := json.Marshal(bodyMessage)
	if err != nil {
		log.Error(err)

		return
	}

	res.Write(marshalBody)
}

func addEasyJSONBody(res http.ResponseWriter, bodyMessage interface{ MarshalJSON() ([]byte, error) }) {
	//marshalBody, err := json.Marshal(bodyMessage)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//res.Write(marshalBody)
	blob, err := bodyMessage.MarshalJSON()
	if err != nil {
		log.Error(err)

		return
	}

	res.Write(blob)
}

func OkResponse(res http.ResponseWriter, bodyMessage interface{}) {
	addOkHeader(res)
	addBody(res, bodyMessage)
}

func addErrHeader(res http.ResponseWriter, errCode int) {
	res.WriteHeader(errCode)
}

func addErrBody(res http.ResponseWriter, errMsg string) {
	addBody(res, errorResponse{Error: errMsg})
}

func ErrResponse(res http.ResponseWriter, errCode int, errMsg string) {
	addErrHeader(res, errCode)
	addErrBody(res, errMsg)
}
