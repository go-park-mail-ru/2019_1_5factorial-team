package donation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/panicWorker"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type Server struct {
	users          UsersMap
	notificationCh chan Notification
	registerUserCh chan User

	GetOperationInfo func(*Notification) error
}

func NewServer() *Server {
	s := &Server{
		users:          NewUsersMap(),
		notificationCh: make(chan Notification, 100),
		registerUserCh: make(chan User, 100),
	}

	return s
}

func (s *Server) Run() error {
	if err := s.initYaClient(); err != nil {
		return errors.Wrap(err, "failed to init Ya client")
	}

	go s.NotificationProcessor()
	go s.UsersProcessor()
	return nil
}

func (s *Server) initYaClient() error {
	yaClient := &http.Client{
		Timeout: time.Second * 2,
	}

	s.GetOperationInfo = func(notification *Notification) error {
		form := url.Values{}
		form.Set("operation_id", notification.OperationID)

		req, err := http.NewRequest("POST", "https://money.yandex.ru/api/operation-details", strings.NewReader(form.Encode()))
		if err != nil {
			return errors.Wrap(err, "failed to create request struct")
		}

		req.Header.Add("Authorization", "Bearer "+config.Get().DonationConfig.AccessToken)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

		res, err := yaClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "failed to create request struct")
		}

		log.Printf("done http  request for operation data %v, \n got response %v\n", req, res)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Wrap(err, "body read error")
		}

		err = json.Unmarshal(body, &notification.NotificationExtendedInfo)
		if err != nil {
			return errors.Wrapf(err, "failed to create parse response %q", string(body))
		}

		if notification.Status != "success" {
			return errors.Wrapf(err, "wrong operation status %s", notification.Status)
		}

		return nil
	}

	return nil
}

func (s *Server) NotificationProcessor() {
	for {
		notification := <-s.notificationCh

		err := s.GetOperationInfo(&notification)
		if err != nil {
			continue
		}

		log.Printf("got notification %v \n, SENDING...\n", notification)

		//TODO: send notification to user map
		s.users.SendMessages(notification.Message)

		// do not spam notifications
		time.Sleep(config.Get().DonationConfig.SleepBetweenNotifications.Duration)
	}
}

func (s *Server) UsersProcessor() {
	for {
		user := <-s.registerUserCh
		s.users.Add(user)
		log.Println("new user to notify ", user)
	}
}

func (s *Server) GetWssHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO: check this function
		ctxLogger := req.Context().Value("logger").(*logrus.Entry)

		// TODO(smet1): ping-pong or ws will be close on timeout
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(res, req, nil)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)

			ctxLogger.Errorf("error while connecting: %s", err)
			return
		}

		user, err := NewUser(conn)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)

			ctxLogger.Errorf("error while connecting: %s", err)
			return
		}

		// TODO: user listen
		go panicWorker.PanicWorker(user.Listen)
		s.registerUserCh <- user
	}
}

func (s *Server) GetNotificationHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if !checkNotification(req) {
			res.WriteHeader(http.StatusBadRequest)

			log.Errorf("wrong checksum for notification request %v")
			return
		}

		amountString := req.PostForm.Get("amount")
		amountF, err := strconv.ParseFloat(amountString, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)

			log.Errorf("cant parse notification amount from %q, error: %v", amountString, err)
			return
		}
		amount := int(amountF)
		log.Printf("got notification: %v", req)

		if amount > config.Get().DonationConfig.MinimumAmountForDonation {
			notification := Notification{
				Amount:      amount,
				OperationID: req.PostForm.Get("operation_id"),
			}
			s.notificationCh <- notification
		}

		res.WriteHeader(http.StatusOK)
	}
}
