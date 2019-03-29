package user

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//type DBConfig struct {
//	MongoPort         string `json:"mongo_port"`
//	DatabaseName      string `json:"database_name"`
//	CollectionName    string `json:"collection_name"`
//	GenerateFakeUsers bool   `json:"generate_fake_users"`
//	TruncateTable     bool   `json:"truncate_table"`
//}

//var ConfigDBUser = DBConfig{}
var session *mgo.Session
var collection *mgo.Collection

// оставляю инит на базу, ибо надо
func init() {
	//err := config_reader.ReadConfigFile("db_user_config.json", &ConfigDBUser)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	//}
	//fmt.Println("DB conf", ConfigDBUser)

	session, err := mgo.Dial("mongodb://localhost:" + config.GetInstance().DBUserConfig.MongoPort)
	if err != nil {
		log.Fatal(err)
	}

	collection = session.DB(config.GetInstance().DBUserConfig.DatabaseName).C(config.GetInstance().DBUserConfig.CollectionName)

	if n, _ := collection.Count(); n != 0 && config.GetInstance().DBUserConfig.TruncateTable {
		err = collection.DropCollection()
		if err != nil {
			log.Fatal("user db truncate: ", err)
		}
	}

	if config.GetInstance().DBUserConfig.GenerateFakeUsers {
		fu := GenerateUsers()

		for i, val := range fu {
			fmt.Println(i, "| id:", val.CollectionID.Hex(), ", Nick:", val.Nickname, ", Password:", val.Nickname)

			err = collection.Insert(val)
			if err != nil {
				log.Fatal(errors.Wrap(err, "error while adding new user"))
			}

		}
	}
}

const DefaultAvatarLink = "../../../img/default.jpg"

type DatabaseUser struct {
	Id           int64         `bson:"-"`
	CollectionID bson.ObjectId `bson:"_id"`
	Email        string        `bson:"email"`
	Nickname     string        `bson:"nickname"`
	HashPassword string        `bson:"hash_password"`
	Score        int           `bson:"score"`
	AvatarLink   string        `bson:"avatar_link"`
}

func getUser(login string) (DatabaseUser, error) {
	u := DatabaseUser{}

	err := collection.Find(bson.M{"nickname": login}).One(&u)
	if err != nil {
		return DatabaseUser{}, errors.New("Invalid login")
	}

	return u, nil
}

func findUserById(id string) (DatabaseUser, error) {
	u := DatabaseUser{}

	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	if err != nil {
		return DatabaseUser{}, errors.New("user with this id not found")
	}

	return u, nil
}

func addUser(nickname string, email string, password string) (User, error) {
	hashPassword, err := GetPasswordHash(password)
	if err != nil {
		return User{}, errors.Wrap(err, "Hasher password error")
	}

	dbu := DatabaseUser{
		CollectionID: bson.NewObjectId(),
		Email:        email,
		Nickname:     nickname,
		HashPassword: hashPassword,
		Score:        0,
		AvatarLink:   "",
	}

	err = collection.Insert(dbu)
	if err != nil {
		return User{}, errors.Wrap(err, "error while adding new user")
	}

	return User{
		Id:           dbu.CollectionID.Hex(),
		Email:        dbu.Email,
		Nickname:     dbu.Nickname,
		HashPassword: dbu.HashPassword,
		Score:        dbu.Score,
		AvatarLink:   dbu.AvatarLink,
	}, nil
}

func updateDBUser(user DatabaseUser) error {

	err := collection.UpdateId(user.CollectionID, user)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

func getScores(limit int, skip int) ([]DatabaseUser, error) {
	result := make([]DatabaseUser, 0, 1)

	err := collection.Find(nil).Skip(skip).
		Sort("-score", "nickname").
		Limit(limit).All(&result)
	if err != nil {
		return nil, errors.Wrap(err, "cant query leaderboard")
	}

	return result, nil
}

func getUsersCount() (int, error) {
	lenTable, err := collection.Count()
	if err != nil {
		return -1, errors.Wrap(err, "cant count users")
	}

	return lenTable, nil
}
