package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"time"
)

var collectionName = "user_session"

const NoTokenFound string = "token not found"

func SetToken(id string) (string, time.Time, error) {
	now := time.Now()

	dt := UserSession{
		Token:             GenerateToken(),
		UserId:            id,
		CookieExpiredTime: now.Add(config.Get().CookieConfig.CookieTimeHours.Duration),
	}

	err := dt.Insert()
	if err != nil {
		return "", time.Time{}, errors.Wrap(err, "cant create session")
	}

	return dt.Token, dt.CookieExpiredTime, nil
}

func UpdateToken(token string) (UserSession, error) {
	us, err := GetSessionByToken(token)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "cant find token")
	}

	exp := us.CheckExpireTime()
	if exp >= 0 && exp <= 10*time.Minute {
		err = us.UpdateTime()
		if err != nil {
			return UserSession{}, errors.Wrap(err, "UpdateToken error, cant update time")
		}
	} else if exp < 0 {
		err = us.Delete()
		if err != nil {
			return UserSession{}, errors.Wrap(err, "UpdateToken error, can't delete expired token")
		}
	}

	return us, nil
}

// при взятии токена, проверяет его на время
func GetId(token string) (string, error) {
	us, err := GetSessionByToken(token)
	if err != nil {
		return "", errors.Wrap(err, "GetId error")
	}

	exp := us.CheckExpireTime()
	if exp < 0 {
		err = us.Delete()
		if err != nil {
			return "", errors.Wrap(err, "cant delete expired token")
		}

		return "", errors.New("your token expired")
	}

	return us.UserId, nil
}

func DeleteToken(token string) error {
	us, err := GetSessionByToken(token)
	if err != nil {
		return errors.Wrap(err, "DeleteToken error get token")
	}

	err = us.Delete()
	if err != nil {
		return errors.Wrap(err, "DeleteToken error delete token")
	}

	return nil
}
