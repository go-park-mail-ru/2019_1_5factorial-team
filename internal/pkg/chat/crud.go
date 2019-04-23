package chat

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/database"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "chat_global"

func (um *UserMessage) Insert() error {
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return errors.Wrap(err, "collection not found")
	}

	um.ID = bson.NewObjectId()
	err = col.Insert(um)
	if err != nil {
		return errors.Wrap(err, "error while adding new message")
	}

	return nil
}

func GetLastMessages(messagesCount int) ([]UserMessage, error) {
	col, err := database.GetCollection(collectionName)
	if err != nil {
		return nil, errors.Wrap(err, "collection not found")
	}

	result := make([]UserMessage, 0, 1)
	err = col.Find(nil).Sort("time").Limit(messagesCount).All(&result)
	if err != nil {
		return nil, errors.Wrap(err, "cant get last messages")
	}

	return result, nil
}
