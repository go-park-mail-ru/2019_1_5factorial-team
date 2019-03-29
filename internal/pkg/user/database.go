package user

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

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

	err := database.GetUserCollection().Find(bson.M{"nickname": login}).One(&u)
	if err != nil {
		return DatabaseUser{}, errors.New("Invalid login")
	}

	return u, nil
}

func findUserById(id string) (DatabaseUser, error) {
	u := DatabaseUser{}

	err := database.GetUserCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
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

	err = database.GetUserCollection().Insert(dbu)
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

	err := database.GetUserCollection().UpdateId(user.CollectionID, user)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

func getScores(limit int, skip int) ([]DatabaseUser, error) {
	result := make([]DatabaseUser, 0, 1)

	err := database.GetUserCollection().Find(nil).Skip(skip).
		Sort("-score", "nickname").
		Limit(limit).All(&result)
	if err != nil {
		return nil, errors.Wrap(err, "cant query leaderboard")
	}

	return result, nil
}

func getUsersCount() (int, error) {
	lenTable, err := database.GetUserCollection().Count()
	if err != nil {
		return -1, errors.Wrap(err, "cant count users")
	}

	return lenTable, nil
}
