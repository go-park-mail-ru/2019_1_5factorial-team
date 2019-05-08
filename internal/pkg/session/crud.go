package session

import (
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
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

	// инфо по вставке (матчи, изменения не нужны)
	_, err = col.Upsert(bson.M{"user_id": us.UserId}, us)
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
