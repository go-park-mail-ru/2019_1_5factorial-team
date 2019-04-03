package database

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"sync"
)

var session *mgo.Session

var collections map[string]*mgo.Collection

var once sync.Once
var mu *sync.Mutex

func InitConnection() {
	once.Do(func() {
		mu = &sync.Mutex{}

		collections = make(map[string]*mgo.Collection)
		var err error

		for _, val := range config.Get().DBConfig {

			//mongodb://mongo-user:27031,
			log.WithFields(log.Fields{
				"dial": fmt.Sprintf("%s://%s:%s", "mongodb", val.Hostname, val.MongoPort),
			}).Info("database.InitConnection")

			session, err = mgo.Dial(fmt.Sprintf("%s://%s:%s", "mongodb", val.Hostname, val.MongoPort))
			if err != nil {
				//session.Close()
				log.Fatal(errors.Wrap(err, "cant connect to db"))
			}

			collection := session.DB(val.DatabaseName).C(val.CollectionName)

			// очистка коллекции по конфигу
			if n, _ := collection.Count(); n != 0 && val.TruncateTable {
				err = collection.DropCollection()
				if err != nil {
					session.Close()
					log.Fatal("db truncate: ", err, val)
				}
			}

			collections[val.CollectionName] = collection
		}
	})
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
