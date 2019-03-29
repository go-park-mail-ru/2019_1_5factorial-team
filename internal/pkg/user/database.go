package user

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func getUser(login string) (User, error) {
	u := User{}

	err := database.GetUserCollection().Find(bson.M{"nickname": login}).One(&u)
	if err != nil {
		return User{}, errors.New("Invalid login")
	}

	return u, nil
}

func findUserById(id string) (User, error) {
	u := User{}

	err := database.GetUserCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	if err != nil {
		return User{}, errors.New("user with this id not found")
	}

	return u, nil
}

func addUser(nickname string, email string, password string) (User, error) {
	hashPassword, err := GetPasswordHash(password)
	if err != nil {
		return User{}, errors.Wrap(err, "Hasher password error")
	}

	dbu := User{
		ID:           bson.NewObjectId(),
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

	return dbu, nil
}

func updateDBUser(user User) error {

	err := database.GetUserCollection().UpdateId(user.ID, user)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

func getScores(limit int, skip int) ([]User, error) {
	result := make([]User, 0, 1)

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
