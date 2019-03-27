package user

import (
	"github.com/pkg/errors"
)

type User struct {
	Id           string
	Email        string
	Nickname     string
	HashPassword string
	Score        int
	AvatarLink   string
}

func CreateUser(nickname string, email string, password string) (User, error) {
	// TODO(smet1): добавить валидацию на повторение ника и почты

	//rid := GetNextId()
	//hashPassword, err := GetPasswordHash(password)
	//if err != nil {
	//	return User{}, errors.Wrap(err, "Hasher password error")
	//}

	u, err := addUser(nickname, email, password)
	if err != nil {
		err = errors.Wrap(err, "Cannot create user")
		return User{}, err
	}

	//u := User{
	//	Id:           rid,
	//	Email:        email,
	//	Nickname:     nickname,
	//	HashPassword: hashPassword,
	//	Score:        0,
	//	AvatarLink:   DefaultAvatarLink,
	//}

	return u, nil
}

func IdentifyUser(login string, password string) (User, error) {
	u, err := getUser(login)
	if err != nil {
		return User{}, errors.Wrap(err, "Can't find user")
	}

	err = comparePassword(password, u.HashPassword)
	if err != nil {
		return User{}, errors.Wrap(err, "Wrong password")
	}

	return User{
		Id:       u.CollectionID.Hex(),
		Email:    u.Email,
		Nickname: u.Nickname,
	}, nil
}

func GetUserById(id string) (User, error) {
	u, err := findUserById(id)
	if err != nil {
		return User{}, errors.Wrap(err, "can't find user")
	}

	return User{
		Id:         u.CollectionID.Hex(),
		Email:      u.Email,
		Nickname:   u.Nickname,
		Score:      u.Score,
		AvatarLink: u.AvatarLink,
	}, nil
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

type Scores struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
}

func GetUsersScores(limit int, offset int) ([]Scores, error) {
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
