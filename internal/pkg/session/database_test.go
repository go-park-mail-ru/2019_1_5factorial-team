package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func InitDB() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	config.Get().DBConfig[0].Hostname = "localhost"
	config.Get().DBConfig[0].MongoPort = "27061"
	config.Get().DBConfig[0].TruncateTable = true

	config.Get().DBConfig[1].Hostname = "localhost"
	config.Get().DBConfig[1].MongoPort = "27062"
	config.Get().DBConfig[1].TruncateTable = true

	config.Get().DBConfig[2].Hostname = "localhost"
	config.Get().DBConfig[2].MongoPort = "27063"
	config.Get().DBConfig[2].TruncateTable = true

	config.Get().AuthGRPCConfig.Hostname = "localhost"

	log.InitLogs()

	database.InitConnection()

	//// indexes user
	//col, _ := database.GetCollection(config.Get().DBConfig[0].CollectionName)
	//err = col.EnsureIndex(mgo.Index{
	//	Key:    []string{"email"},
	//	Unique: true,
	//})
	//if err != nil {
	//	logrus.Fatal(err.Error())
	//}
	//
	//err = col.EnsureIndex(mgo.Index{
	//	Key:    []string{"nickname"},
	//	Unique: true,
	//})
	//if err != nil {
	//	logrus.Fatal(err.Error())
	//}
	//
	//// indexes session
	//col, _ = database.GetCollection(config.Get().DBConfig[1].CollectionName)
	//err = col.EnsureIndex(mgo.Index{
	//	Key:    []string{"user_id"},
	//	Unique: true,
	//})
	//if err != nil {
	//	logrus.Fatal(err.Error())
	//}
	//
	//err = col.EnsureIndex(mgo.Index{
	//	Key:    []string{"token"},
	//	Unique: true,
	//})
	//if err != nil {
	//	logrus.Fatal(err.Error())
	//}
}

func TestSetToken(t *testing.T) {
	InitDB()
	lula, _ := user.CreateUser("Lula", "lula@gmail.com", "lula")
	// #0
	now := time.Now()
	token, expiration, err := SetToken(lula.ID.Hex())
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	if expiration.Hour() != now.Hour()+int(config.Get().CookieConfig.CookieTimeHours.Hours()) {
		t.Error("#", 0, "wrong expiration time expected:", now.Hour()+int(config.Get().CookieConfig.CookieTimeHours.Hours()), "have:", expiration.Hour())
	}

	if token == "" {
		t.Error("#", 0, "no token found")
	}
}

func TestUpdateToken(t *testing.T) {
	InitDB()
	config.Get().CookieConfig.CookieTimeHours = config.Duration{Duration: 2 * time.Second}

	lula, _ := user.CreateUser("Lula", "lula@gmail.com", "lula")
	// #0
	token, exp, err := SetToken(lula.ID.Hex())
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	us, err := UpdateToken(token)
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	if us.CookieExpiredTime == exp {
		t.Error("#", 0, "time should be changed, old", exp, "have:", us.CookieExpiredTime)
	}

	// #1
	config.Get().CookieConfig.CookieTimeHours = config.Duration{Duration: 1 * time.Second}
	time.Sleep(4 * time.Second)
	us, err = UpdateToken(us.Token)
	if err != nil {
		t.Error("#", 1, "no ERROR expected", "have:", err)
	}

	if us.CookieExpiredTime == exp {
		t.Error("#", 1, "time should be changed, old", exp, "have:", us.CookieExpiredTime)
	}

	// #2
	_, err = UpdateToken("32324")
	if err == nil {
		t.Error("#", 1, "ERROR expected", "have:", err)
	}
}

func TestDeleteToken(t *testing.T) {
	InitDB()
	lula, _ := user.CreateUser("Lula", "lula@gmail.com", "lula")

	// #0
	token, _, err := SetToken(lula.ID.Hex())
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	err = DeleteToken(token)
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	// #1
	err = DeleteToken(token)
	if err == nil {
		t.Error("#", 0, "ERROR expected", "have:", err)
	}
}

func TestGetId(t *testing.T) {
	InitDB()
	lula, _ := user.CreateUser("Lula", "lula@gmail.com", "lula")

	// #0
	token, _, err := SetToken(lula.ID.Hex())
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	id, err := GetId(token)
	if err != nil {
		t.Error("#", 0, "no ERROR expected", "have:", err)
	}

	if id != lula.ID.Hex() {
		t.Error("#", 0, "mismatch id, expected", lula.ID.Hex(), "have:", id)
	}

	// #1
	id, err = GetId("sdgsdfasdfsddsf")
	if err == nil {
		t.Error("#", 0, "ERROR expected", "have:", err)
	}
}
