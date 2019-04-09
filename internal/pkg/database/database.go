package database

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"time"
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
			fmt.Println(fmt.Sprintf("%s://%s:%s", "mongodb", val.Hostname, val.MongoPort))

			session, err = mgo.DialWithInfo(&mgo.DialInfo{
				Addrs:    []string{fmt.Sprintf("%s:%s", val.Hostname, val.MongoPort)},
				Timeout:  10 * time.Second,
				Database: val.DatabaseName,
			})
			if err != nil {
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
