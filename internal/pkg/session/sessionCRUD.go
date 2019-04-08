package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserSession struct {
	ID                bson.ObjectId `bson:"_id"`
	Token             string        `bson:"token"`
	UserId            string        `bson:"user_id"`
	CookieExpiredTime time.Time     `bson:"cookie_expired_time"`
}

func (us *UserSession) Insert() error {
	us.ID = bson.NewObjectId()

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	// если юзер залогинин уже залогинин с другого устройства, обновляю ему сессию (не создаю новую)
	usConflict, err := getCollisionSession(us.UserId)
	// ошибки с отсутсвтем коллекции не может быть, тк проверено выше
	if err == nil {
		err = usConflict.UpdateToken(us.Token)
		if err != nil {
			return errors.Wrap(err, "error while updating exist session token")
		}

		return nil
	}


	err = col.Insert(us)
	if err != nil {
		return errors.Wrap(err, "error while adding user session")
	}

	return nil
}

func (us *UserSession) UpdateTime() error  {
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

func getCollisionSession(id string) (UserSession, error) {
	us := UserSession{}

	col, err := database.GetCollection(collectionName)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "collection not found")
	}

	err = col.Find(bson.M{"user_id": id}).One(&us)
	if err != nil {
		return UserSession{}, errors.Wrap(err, "user not found")
	}

	return us, nil
}