package session

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WorkerSession struct {
	SleepTime string `json:"sleep_time_minute"`
}

func RemoveSession(port *string) {

	fmt.Println("---=== CLEANING NONALIDE SESSIONS ===---")
	session, err := mgo.Dial("mongodb://localhost:" + *port)
	if err != nil {

		log.Fatal(err)
	}

	collection := session.DB("user_session").C("user_session")

	_, err = collection.RemoveAll(bson.M{"cookie_expired_time": bson.M{"$lt": time.Now()}})
	if err != nil {
		log.Fatal(errors.Wrap(err, "error remove session"))
	}

	session.Close()

}
