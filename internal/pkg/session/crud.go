package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserSession struct {
	ID                bson.ObjectId `bson:"_id,omitempty"`
	Token             string        `bson:"token"`
	UserId            string        `bson:"user_id"`
	CookieExpiredTime time.Time     `bson:"cookie_expired_time"`
}

func (us *UserSession) Insert() error {
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	info, err := col.Upsert(bson.M{"user_id": us.UserId}, us)
	fmt.Println(info)
	fmt.Println(err)

	if err != nil {
		return errors.Wrap(err, "error while adding user session")
	}

	return nil
}

func (us *UserSession) UpdateTime() error {
	us.CookieExpiredTime = time.Now().Add(config.Get().CookieConfig.CookieTimeHours.Duration)

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	err = col.UpdateId(us.ID, us)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

// UpdateToken обновляет токен текущей сессии юзера и сеттит ей новое время
func (us *UserSession) UpdateToken(newToken string) error {
	us.Token = newToken
	us.CookieExpiredTime = time.Now().Add(config.Get().CookieConfig.CookieTimeHours.Duration)

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	err = col.UpdateId(us.ID, us)
	if err != nil {
		return errors.Wrap(err, "error while updating value in DB")
	}

	return nil
}

func (us *UserSession) Delete() error {
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	err = col.RemoveId(us.ID)
	if err != nil {
		return errors.Wrap(err, "cant delete token")
	}

	return nil
}

func (us *UserSession) CheckExpireTime() time.Duration {
	return us.CookieExpiredTime.Sub(time.Now())
}

func GetSessionByToken(token string) (UserSession, error) {
	us := UserSession{}

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "collection not found")
	}

	err = col.Find(bson.M{"token": token}).One(&us)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "token not found")
	}

	return us, nil
}
