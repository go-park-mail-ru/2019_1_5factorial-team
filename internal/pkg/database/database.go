package database

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"gopkg.in/mgo.v2"
	"log"
	"sync"
)

var session *mgo.Session
var userCollection *mgo.Collection

var once sync.Once

func InitConnection() {
	once.Do(func() {
		var err error
		session, err = mgo.Dial("mongodb://localhost:" + config.GetInstance().DBUserConfig.MongoPort)
		if err != nil {
			log.Fatal(err)
		}

		userCollection = session.DB(config.GetInstance().DBUserConfig.DatabaseName).
			C(config.GetInstance().DBUserConfig.CollectionName)

		// очистка коллекции юзеров по конфигу
		if n, _ := userCollection.Count(); n != 0 && config.GetInstance().DBUserConfig.TruncateTable {
			err = userCollection.DropCollection()
			if err != nil {
				log.Fatal("user db truncate: ", err)
			}
		}

		//// заполнение коллекции юзеров по конфигу
		//if config.GetInstance().DBUserConfig.GenerateFakeUsers {
		//	fu := user.GenerateUsers()
		//
		//	for i, val := range fu {
		//		fmt.Println(i, "| id:", val.CollectionID.Hex(), ", Nick:", val.Nickname, ", Password:", val.Nickname)
		//
		//		err = userCollection.Insert(val)
		//		if err != nil {
		//			log.Fatal(errors.Wrap(err, "error while adding new user"))
		//		}
		//
		//	}
		//}

	})
}

func GetDBSesion() *mgo.Session {
	return session
}

func GetUserCollection() *mgo.Collection {
	return userCollection
}
