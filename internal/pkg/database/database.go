package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var collections map[string]*mgo.Collection

var mu *sync.Mutex

func InitConnection() {
	mu = &sync.Mutex{}

	collections = make(map[string]*mgo.Collection)
	var err error

	for _, val := range config.Get().DBConfig {

		//mongodb://mongo-user:27031,
		fmt.Println(fmt.Sprintf("%s://%s:%s", "mongodb", val.Hostname, val.MongoPort))

		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{fmt.Sprintf("%s:%s", val.Hostname, val.MongoPort)},
			Timeout:  10 * time.Second,
			Database: val.DatabaseName,
		})
		if err != nil {
			logrus.Fatal(err)
		}

		collection := session.DB(val.DatabaseName).C(val.CollectionName)

		// очистка коллекции по конфигу
		if n, err := collection.Count(); n != 0 && err == nil && val.TruncateTable {
			logrus.Warn("truncating db", val.CollectionName)
			err = collection.DropCollection()
			if err != nil {
				session.Close()
				//logrus.Fatal("db truncate: ", err, val)
				logrus.Warn("db truncate: ", err, val)
			}
		}

		collections[val.CollectionName] = collection
	}
}

func GetCollection(name string) (*mgo.Collection, error) {
	defer mu.Unlock()
	mu.Lock()

	if i, ok := collections[name]; !ok {
		return nil, errors.New("no connection found, name = " + name)
	} else {
		return i, nil
	}
}
