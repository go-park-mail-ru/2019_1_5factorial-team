package user

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/manveru/faker"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
)

type DBConfig struct {
	MongoPort      string `json:"mongo_port"`
	DatabaseName   string `json:"database_name"`
	CollectionName string `json:"collection_name"`
}

var ConfigDBUser = DBConfig{}
var session *mgo.Session
var collection *mgo.Collection

func init() {
	err := config_reader.ReadConfigFile("db_user_config.json", &ConfigDBUser)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	}
	fmt.Println("DB conf", ConfigDBUser)

	session, err = mgo.Dial("mongodb://localhost:" + ConfigDBUser.MongoPort)
	if err != nil {
		log.Fatal(err)
	}

	collection = session.DB(ConfigDBUser.DatabaseName).C(ConfigDBUser.CollectionName)

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

var once sync.Once
var mu *sync.Mutex
var users map[string]DatabaseUser
var currentId int64

func init() {
	once.Do(func() {
		fmt.Println("init users map")

		fake, _ := faker.New("en")
		fake.Rand = rand.New(rand.NewSource(42))

		users = make(map[string]DatabaseUser)

		hash, _ := getPasswordHash("password")
		users["kekkekkek"] = DatabaseUser{
			Id:           0,
			Email:        "kek.k.ek",
			Nickname:     "kekkekkek",
			HashPassword: hash,
			Score:        100500,
			AvatarLink:   DefaultAvatarLink,
		}
		mu = &sync.Mutex{}
		currentId = 0

		var id int64
		// TODO(smet1): generate fake accs in func
		for i := 0; i < 20; i++ {
			id = GetNextId()
			nick := fake.FirstName()
			hash, _ := getPasswordHash(nick)

			fmt.Println("id:", id, ", Nick:", nick, ", Password:", nick)

			users[nick] = DatabaseUser{
				Id:           id,
				Email:        fake.Email(),
				Nickname:     nick,
				HashPassword: hash,
				Score:        rand.Intn(250000),
				AvatarLink:   DefaultAvatarLink,
			}

		}

	})
}

func getUsers() map[string]DatabaseUser {
	mu.Lock()
	fmt.Println(users)
	mu.Unlock()

	return users
}

func getUser(login string) (DatabaseUser, error) {
	//defer mu.Unlock()
	//
	//mu.Lock()
	//if _, ok := users[login]; !ok {
	//	return DatabaseUser{}, errors.New("Invalid login")
	//} else {
	//	return users[login], nil
	//}

	u := DatabaseUser{}

	err := collection.Find(bson.M{"nickname": login}).One(&u)
	if err != nil {
		return DatabaseUser{}, errors.New("Invalid login")
	}

	return u, nil
}

func findUserById(id string) (DatabaseUser, error) {
	//defer mu.Unlock()
	//
	//mu.Lock()
	//for _, val := range users {
	//	if val.Id == id {
	//		return val, nil
	//	}
	//}

	u := DatabaseUser{}

	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	if err != nil {
		return DatabaseUser{}, errors.New("user with this id not found")
	}

	return u, nil
}

func addUser(nickname string, email string, password string) (User, error) {
	//defer mu.Unlock()
	//mu.Lock()
	//
	//if _, ok := users[in.Nickname]; ok {
	//	return errors.New("User with this nickname already exist")
	//}
	//
	//users[in.Nickname] = in

	hashPassword, err := getPasswordHash(password)
	if err != nil {
		return User{}, errors.Wrap(err, "Hasher password error")
	}

	dbu := DatabaseUser{
		CollectionID: bson.NewObjectId(),
		Email: email,
		Nickname: nickname,
		HashPassword: hashPassword,
		Score: 0,
		AvatarLink: DefaultAvatarLink,
	}

	//in.CollectionID = bson.NewObjectId()
	err = collection.Insert(dbu)
	if err != nil {
		return User{}, errors.Wrap(err, "error while adding new user")
	}

	return User{
		Id: dbu.CollectionID.Hex(),
		Email: dbu.Email,
		Nickname: dbu.Nickname,
		HashPassword: dbu.HashPassword,
		Score: dbu.Score,
		AvatarLink: dbu.AvatarLink,
	}, nil
}

func PrintUsers() {
	mu.Lock()

	for i, val := range users {
		fmt.Println("\t", i, val)
	}
	fmt.Println("----end----")

	mu.Unlock()
}

func GetNextId() int64 {
	atomic.AddInt64(&currentId, 1)

	return currentId
}

func updateDBUser(user DatabaseUser) error {
	//defer mu.Unlock()
	//
	//mu.Lock()
	//if _, ok := users[user.Nickname]; !ok {
	//	return errors.New("cannot find user")
	//}
	//
	//users[user.Nickname] = user
	//return nil

	err := collection.UpdateId(user.CollectionID, user)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

type ByNameScore []DatabaseUser

func (a ByNameScore) Len() int           { return len(a) }
func (a ByNameScore) Less(i, j int) bool { return a[i].Score < a[j].Score }
func (a ByNameScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getScores() []DatabaseUser {
	mu.Lock()
	result := make([]DatabaseUser, 0, 1)
	for _, val := range users {
		result = append(result, val)
	}

	sort.Sort(ByNameScore(result))

	mu.Unlock()
	return result
}

func getUsersCount() (int, error) {
	lenTable, err := collection.Count()
	if err != nil {
		return -1, errors.Wrap(err, "cant count users")
	}

	return lenTable, nil
}
