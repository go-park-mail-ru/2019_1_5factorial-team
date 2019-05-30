package user

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"strconv"
	"testing"
)

func InitDB() {
	configPath := "/etc/5factorial/"
	err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	config.Get().DBConfig[0].Hostname = "localhost"
	config.Get().DBConfig[0].MongoPort = "27081"
	config.Get().DBConfig[0].TruncateTable = true

	config.Get().DBConfig[1].Hostname = "localhost"
	config.Get().DBConfig[1].MongoPort = "27082"
	config.Get().DBConfig[1].TruncateTable = true

	config.Get().DBConfig[2].Hostname = "localhost"
	config.Get().DBConfig[2].MongoPort = "27063"
	config.Get().DBConfig[2].TruncateTable = true

	config.Get().AuthGRPCConfig.Hostname = "localhost"

	log.InitLogs()

	database.InitConnection()

	// indexes user
	col, _ := database.GetCollection(config.Get().DBConfig[0].CollectionName)
	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"nickname"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// indexes session
	col, _ = database.GetCollection(config.Get().DBConfig[1].CollectionName)
	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"user_id"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = col.EnsureIndex(mgo.Index{
		Key:    []string{"token"},
		Unique: true,
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

var casesCreateUser = []struct {
	password string
	want     User
	err      string
}{
	{
		password: "",
		want: User{
			Email:      "",
			Nickname:   "",
			Score:      0,
			AvatarLink: "",
		},
		err: "empty field",
	},
	{
		password: "qwerty",
		want: User{
			Email:      "qwerty@qwe.rty",
			Nickname:   "qwerty",
			Score:      0,
			AvatarLink: "000-default-avatar",
		},
		err: "",
	},
}

func TestCreateUser(t *testing.T) {
	InitDB()
	for _, val := range casesCreateUser {
		res, err := CreateUser(val.want.Nickname, val.want.Email, val.password, "")

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("ERROR expected:", val.err, "have:", err)
			}
		}
		if res.Nickname != val.want.Nickname {
			t.Error("NICKNAME expected:", val.want.Nickname, "have:", res.Nickname)
		}
		if res.Email != val.want.Email {
			t.Error("EMAIL expected:", val.want.Email, "have:", res.Email)
		}
		if res.Score != val.want.Score {
			t.Error("SCORE expected:", val.want.Score, "have:", res.Score)
		}
		if res.AvatarLink != val.want.AvatarLink {
			t.Error("AVATAR expected:", val.want.AvatarLink, "have:", res.AvatarLink)
		}
	}
}

var casesIdentifyUser = []struct {
	loginOrEmail string
	password     string
	want         User
	err          string
}{
	{
		loginOrEmail: "",
		password:     "",
		want: User{
			Email:      "",
			Nickname:   "",
			Score:      0,
			AvatarLink: "",
		},
		err: "Invalid login",
	},
	{
		loginOrEmail: "qwerty",
		password:     "qwerty",
		want: User{
			Email:      "",
			Nickname:   "",
			Score:      0,
			AvatarLink: "",
		},
		err: "Invalid login",
	},
	{
		loginOrEmail: "Lula",
		password:     "qwerty",
		want: User{
			Email:      "",
			Nickname:   "",
			Score:      0,
			AvatarLink: "",
		},
		err: "crypto/bcrypt: hashedPassword is not the hash of the given password",
	},
	{
		loginOrEmail: "Lula",
		password:     "lula",
		want: User{
			Email:      "lula@gmail.com",
			Nickname:   "Lula",
			Score:      0,
			AvatarLink: "000-default-avatar",
		},
		err: "",
	},
	{
		loginOrEmail: "lula@gmail.com",
		password:     "lula",
		want: User{
			Email:      "lula@gmail.com",
			Nickname:   "Lula",
			Score:      0,
			AvatarLink: "000-default-avatar",
		},
		err: "",
	},
	{
		loginOrEmail: "lula1@gmail.com",
		password:     "lula",
		want: User{
			Email:      "",
			Nickname:   "",
			Score:      0,
			AvatarLink: "",
		},
		err: "Invalid email",
	},
}

func TestIdentifyUser(t *testing.T) {
	InitDB()
	_, err := CreateUser("Lula", "lula@gmail.com", "lula", "")
	if err != nil {
		log.Fatal(err.Error())
	}

	for i, val := range casesIdentifyUser {
		res, err := IdentifyUser(val.loginOrEmail, val.password)

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("#", i, "ERROR expected:", val.err, "have:", errors.Cause(err).Error())
				continue
			}
		}
		if res.Nickname != val.want.Nickname {
			t.Error("#", i, "NICKNAME expected:", val.want.Nickname, "have:", res.Nickname)
			continue
		}
		if res.Email != val.want.Email {
			t.Error("#", i, "EMAIL expected:", val.want.Email, "have:", res.Email)
			continue
		}
		if res.Score != val.want.Score {
			t.Error("#", i, "SCORE expected:", val.want.Score, "have:", res.Score)
			continue
		}
		if res.AvatarLink != val.want.AvatarLink {
			t.Error("#", i, "AVATAR expected:", val.want.AvatarLink, "have:", res.AvatarLink)
			continue
		}
	}
}

var casesUpdateUserEmpty = []struct {
	id          string
	newAvatar   string
	oldPassword string
	newPassword string
	err         string
}{
	{
		id:          "",
		newAvatar:   "",
		oldPassword: "",
		newPassword: "",
		err:         "nothing to update",
	},
	{
		id:          "5556c0d9b49cd4582aaad41c",
		newAvatar:   "1",
		oldPassword: "2",
		newPassword: "3",
		err:         "user with this id not found",
	},
}

var casesUpdateUser = []struct {
	id          string
	newAvatar   string
	oldPassword string
	newPassword string
	err         string
}{
	{
		newAvatar:   "",
		oldPassword: "",
		newPassword: "",
		err:         "nothing to update",
	},
	{
		newAvatar:   "1",
		oldPassword: "",
		newPassword: "",
		err:         "",
	},
	{
		newAvatar:   "",
		oldPassword: "",
		newPassword: "1",
		err:         "please input old password",
	},
	{
		newAvatar:   "",
		oldPassword: "1",
		newPassword: "1",
		err:         "invalid new password",
	},
	{
		newAvatar:   "",
		oldPassword: "1",
		newPassword: "2123123123",
		err:         "crypto/bcrypt: hashedPassword is not the hash of the given password",
	},
	{
		newAvatar:   "",
		oldPassword: "lula",
		newPassword: "lula",
		err:         "old and new password are same",
	},
	{
		newAvatar:   "",
		oldPassword: "lula",
		newPassword: "lula123",
		err:         "old and new password are same",
	},
}

func TestUpdateUser(t *testing.T) {
	InitDB()
	u, err := addUser("Lula", "lula@gmail.com", "lula", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	for i, val := range casesUpdateUserEmpty {
		err := UpdateUser(val.id, val.newAvatar, val.oldPassword, val.newPassword)

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("#", i, "(Empty)ERROR expected:", val.err, "have:", errors.Cause(err).Error())
				continue
			}
		}
	}

	for i, val := range casesUpdateUser {
		err := UpdateUser(u.ID.Hex(), val.newAvatar, val.oldPassword, val.newPassword)

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("#", i, "(Not Empty)ERROR expected:", val.err, "have:", errors.Cause(err))
				continue
			}
		}
	}
}

var casesGetUsersScores = []struct {
	limit  int
	offset int
	res    []Scores
	err    string
}{
	{
		limit:  0,
		offset: 0,
		res: []Scores{
			{Nickname: "Lula", Score: 0},
			{Nickname: "Lula1", Score: 0},
			{Nickname: "Lula2", Score: 0},
		},
		err: "",
	},
	{
		limit:  100000,
		offset: 400,
		res:    []Scores{},
		err:    "limit * (offset - 1) > users count",
	},
	{
		limit:  -1,
		offset: -1,
		res:    []Scores{},
		err:    "invalid limit value",
	},
	{
		limit:  0,
		offset: -1,
		res:    []Scores{},
		err:    "invalid offset value",
	},
}

func TestGetUsersScores(t *testing.T) {
	InitDB()
	_, _ = addUser("Lula", "lula@gmail.com", "lula", "")
	_, _ = addUser("Lula1", "lula1@gmail.com", "lula", "")
	_, _ = addUser("Lula2", "lula2@gmail.com", "lula", "")

	for i, val := range casesGetUsersScores {
		res, err := GetUsersScores(val.limit, val.offset)

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("#", i, "ERROR expected:", val.err, "have:", err)
				continue
			}
		}

		if len(res) != len(val.res) {
			t.Error("#", i, "RES expected size:", len(val.res), "have size:", len(res))
			continue
		}

		for j := 0; j < len(res)-1; j++ {
			if val.res[j].Score != res[j].Score || val.res[j].Nickname != res[j].Nickname {
				t.Error("#", i, "RES expected(", j, "):", val.res[j], "have size:", res[j])
			}
		}
	}
}

var casesGetUsersCount = []struct {
	insert int
	res    int
	err    string
}{
	{
		insert: 0,
		res:    0,
		err:    "",
	},
	{
		insert: 1,
		res:    1,
		err:    "",
	},
}

func TestGetUsersCount(t *testing.T) {
	InitDB()

	for i, val := range casesGetUsersCount {
		for i := 0; i < val.insert; i++ {
			_, _ = addUser("Lula"+strconv.Itoa(i), "lula"+strconv.Itoa(i)+"@gmail.com", "lula", "")
		}

		res, err := GetUsersCount()

		if err != nil {
			if errors.Cause(err).Error() != val.err {
				t.Error("#", i, "ERROR expected:", val.err, "have:", err)
				continue
			}
		}

		if res != val.res {
			t.Error("#", i, "RES expected:", val.res, "have:", res)
			continue
		}
	}
}

func TestUpdateScore(t *testing.T) {
	InitDB()
	lula, _ := addUser("Lula", "lula@gmail.com", "lula", "")

	// #0
	err := UpdateScore("dsgsdfsd", 0)
	if err == nil {
		t.Error("#", 0, "ERROR expected", "have none")
	}

	// #1
	err = UpdateScore(lula.ID.Hex(), 0)
	if err != nil {
		t.Error("#", 1, "no ERROR expected", "have:", err)
	}

	u, err := GetUserById(lula.ID.Hex())
	if err != nil {
		t.Error("#", 1, "no ERROR expected", "have:", err)
	}

	if u.Score != lula.Score {
		t.Error("#", 1, "RES expected:", lula.Score, "have:", u.Score)
	}

	// #2
	err = UpdateScore(lula.ID.Hex(), 50)
	if err != nil {
		t.Error("#", 1, "no ERROR expected", "have:", err)
	}

	u, err = GetUserById(lula.ID.Hex())
	if err != nil {
		t.Error("#", 1, "no ERROR expected", "have:", err)
	}

	if u.Score != lula.Score+50 {
		t.Error("#", 1, "RES expected:", lula.Score, "have:", u.Score)
	}
}

var casesMarshalJSON = []struct {
	nick  string
	score int
	res   []uint8
	err   string
}{
	{
		nick:  "",
		score: 0,
		res:   []uint8{123, 34, 110, 105, 99, 107, 110, 97, 109, 101, 34, 58, 34, 34, 44, 34, 115, 99, 111, 114, 101, 34, 58, 48, 125},
		err:   "",
	},
}

func TestScores_MarshalJSON(t *testing.T) {
	for i, val := range casesMarshalJSON {
		sc := Scores{
			Nickname: val.nick,
			Score:    val.score,
		}
		res, err := sc.MarshalJSON()
		if err != nil {
			if err.Error() != val.err {
				t.Error("#", i, "ERROR expected:", val.err, "have:", err)
				continue
			}
		}

		if len(res) != len(val.res) {
			t.Error("#", i, "RES expected:", val.res, "have:", res)
			continue
		}

		for j, it := range res {
			if it != val.res[j] {
				t.Error("#", i, "RES expected:", val.res, "have:", res)
				continue
			}
		}
		//if res != val.res {
		//	t.Error("#", i, "RES expected:", val.res, "have:", res)
		//	continue
		//}
	}
}
