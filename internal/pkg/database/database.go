package database

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"log"
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

		for _, val := range config.Get().DBUserConfig {
			session, err = mgo.Dial("mongodb://mongo:" + val.MongoPort)
			if err != nil {
				//session.Close()
				log.Fatal(err)
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
