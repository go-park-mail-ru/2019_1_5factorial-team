package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var collectionName = "user_session"

const NoTokenFound string = "token not found"

type UserSession struct {
	ID                bson.ObjectId `bson:"_id"`
	Token             string        `bson:"token"`
	UserId            string        `bson:"user_id"`
	CookieExpiredTime time.Time     `bson:"cookie_expired_time"`
}

func SetToken(id string) (string, time.Time, error) {
	now := time.Now()

	dt := UserSession{
		ID:                bson.NewObjectId(),
		Token:             GenerateToken(),
		UserId:            id,
		CookieExpiredTime: now.Add(config.Get().CookieConfig.CookieTimeHours.Duration),
	}

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return "", time.Time{}, errors.Wrap(err, "collection not found")
	}

	err = col.Insert(dt)
	if err != nil {
		return "", time.Time{}, errors.Wrap(err, "error while adding new user")
	}

	return dt.Token, dt.CookieExpiredTime, nil
}

func UpdateToken(token string) (UserSession, error) {
	dt := UserSession{}
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "collection not found")
	}

	err = col.Find(bson.M{"token": token}).One(&dt)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "token not found")
	}

	// если через 10 минут кука умрет, добавим ему времени
	if dt.CookieExpiredTime.Sub(time.Now()) < 10*time.Minute && dt.CookieExpiredTime.Sub(time.Now()) > 0 {
		dt.CookieExpiredTime = time.Now().Add(config.Get().CookieConfig.CookieTimeHours.Duration)

		err = col.UpdateId(dt.ID, dt)
		if err != nil {
			return UserSession{}, errors.Wrap(err, "error while updating value in DB")
		}
	} else if dt.CookieExpiredTime.Sub(time.Now()) < 0 {
		// убиваем истекшую куку, вряд ли такое случится (хз)
		err = DeleteToken(dt.Token)
		if err != nil {
			return UserSession{}, errors.Wrap(err, "cant delete expired token")
		}

		return UserSession{}, errors.New("your token expired")
	}

	return dt, nil
}

// при взятии токена, проверяет его на время
func GetId(token string) (string, error) {
	dt := UserSession{}
	col, err := database.GetCollection(collectionName)
	if err != nil {
		log.Println(errors.Wrap(err, "collection not found"))

		return "", errors.Wrap(err, "collection not found")
	}

	err = col.Find(bson.M{"token": token}).One(&dt)
	if err != nil {
		log.Println(errors.Wrap(err, "token not found"))

		return "", errors.Wrap(err, "token not found")
	}

	fmt.Println("--== get id ==--")
	fmt.Println(dt.CookieExpiredTime.Sub(time.Now()))

	if dt.CookieExpiredTime.Sub(time.Now()) < 0 {
		err = DeleteToken(dt.Token)
		if err != nil {
			log.Println(errors.Wrap(err, "cant delete expired token"))

			return "", errors.Wrap(err, "cant delete expired token")
		}

		log.Println(errors.New("your token expired"))

		return "", errors.New("your token expired")
	}

	return dt.UserId, nil
}

func DeleteToken(token string) error {
	dt := UserSession{}
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	err = col.Find(bson.M{"token": token}).One(&dt)
	if err != nil {
		return errors.Wrap(err, "token not found")
	}

	err = col.RemoveId(dt.ID)
	return err
}
