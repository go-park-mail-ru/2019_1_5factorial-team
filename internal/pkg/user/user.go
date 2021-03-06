package user

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/validator"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           bson.ObjectId `bson:"_id"`
	Email        string        `bson:"email"`
	Nickname     string        `bson:"nickname"`
	HashPassword string        `bson:"hash_password"`
	Score        int           `bson:"score"`
	AvatarLink   string        `bson:"avatar_link"`
}

func CreateUser(nickname string, email string, password string, avatar string) (User, error) {
	// TODO(smet1): добавить валидацию на повторение ника и почты
	if nickname == "" || email == "" || password == "" {
		return User{}, errors.New("empty field")
	}

	if avatar == "" {
		avatar = "000-default-avatar.png"
	}

	u, err := addUser(nickname, email, password, avatar)
	if err != nil {
		err = errors.Wrap(err, "Cannot create user")
		return User{}, err
	}

	return u, nil
}

func IdentifyUser(loginOrEmail string, password string) (User, error) {
	u, err := getUser(loginOrEmail)
	if err != nil {
		return User{}, errors.Wrap(err, "Can't find user")
	}

	err = comparePassword(password, u.HashPassword)
	if err != nil {
		return User{}, errors.Wrap(err, "Wrong password")
	}

	return u, nil
}

func GetUserById(id string) (User, error) {
	u, err := findUserById(id)
	if err != nil {
		return User{}, errors.Wrap(err, "can't find user")
	}

	return u, nil
}

func UpdateUser(id string, newAvatar string, oldPassword string, newPassword string) error {
	if newPassword == "" && newAvatar == "" {
		return errors.New("nothing to update")
	}

	u, err := findUserById(id)
	if err != nil {
		return errors.Wrap(err, "update user error")
	}

	if newAvatar != "" {
		u.AvatarLink = newAvatar
		err := updateDBUser(u)
		if err != nil {
			return errors.Wrap(err, "cant update avatar")
		}
	}

	if newPassword == "" {
		return nil
	}

	err = validateChangingPasswords(oldPassword, newPassword, u.HashPassword)
	if err != nil {
		return errors.Wrap(err, "validate passwords error")
	}

	newHashPassword, err := GetPasswordHash(newPassword)
	if err != nil {
		return errors.Wrap(err, "some password error")
	}

	u.HashPassword = newHashPassword
	err = updateDBUser(u)
	if err != nil {
		return errors.Wrap(err, "cant update avatar")
	}

	return nil
}

func validateChangingPasswords(oldPassword string, newPassword string, currentHashPassword string) error {
	if oldPassword == "" {
		return errors.New("please input old password")
	}

	flagValidNewPassword := validator.ValidUpdatePassword(newPassword)
	if !flagValidNewPassword {
		return errors.New("invalid new password")
	}

	err := comparePassword(newPassword, currentHashPassword)
	if err == nil {
		return errors.New("old and new password are same")
	}

	err = comparePassword(oldPassword, currentHashPassword)
	if err != nil {
		return errors.Wrap(err, "invalid old password")
	}

	return nil
}

//easyjson:json
type Scores struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
}

func GetUsersScores(limit int, offset int) ([]Scores, error) {
	if limit < 0 {
		return nil, errors.New("invalid limit value")
	}

	if offset < 0 {
		return nil, errors.New("invalid offset value")
	}

	begin := limit * (offset - 1)
	end := limit * offset

	usersCount, err := getUsersCount()
	if err != nil {
		return nil, errors.Wrap(err, "GetUsersScores-cant get users count")
	}

	if begin > usersCount {
		return nil, errors.New("limit * (offset - 1) > users count")
	}

	if end > usersCount {
		end = usersCount
	}

	page := make([]Scores, 0, 1)

	databaseScores, err := getScores(limit, begin)
	if err != nil {
		return nil, errors.Wrap(err, "cant get score")
	}

	for _, val := range databaseScores {
		page = append(page, Scores{
			Nickname: val.Nickname,
			Score:    val.Score,
		})
	}

	return page, nil
}

func GetUsersCount() (int, error) {
	usersCount, err := getUsersCount()
	if err != nil {
		return -1, errors.Wrap(err, "GetUsersScores-cant get users count")
	}

	return usersCount, nil
}

func UpdateScore(id string, score int) error {
	u, err := findUserById(id)
	if err != nil {
		return errors.Wrap(err, "cant update score")
	}

	if u.Score < score {
		u.Score = score
		err = updateDBUser(u)
		if err != nil {
			return errors.Wrap(err, "cant update user score")
		}
	}

	return nil
}
