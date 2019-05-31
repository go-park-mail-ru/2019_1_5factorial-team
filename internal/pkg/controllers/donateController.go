package controllers

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"io/ioutil"
	"net/http"
)

//func ConnectDonateNotification(res http.ResponseWriter, req *http.Request) {
//	ctxLogger := req.Context().Value("logger").(*logrus.Entry)
//
//	upgrader := websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}
//
//	conn, err := upgrader.Upgrade(res, req, nil)
//	if err != nil {
//		ctxLogger.Printf("error while connecting: %s", err)
//		res.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//}

func GetNotification(res http.ResponseWriter, req *http.Request)  {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("cant get body request")
	}

	log.Println("body:", body)

	OkResponse(res, nil)
}