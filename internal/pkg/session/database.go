package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/pkg/errors"
	"sync"
	"time"
)

var collectionName = "session"

const NoTokenFound string = "token not found"

type DatabaseToken struct {
	UserId            string
	CookieExpiredTime time.Time
	//CookieIssuedTime  time.Time
}

var once sync.Once
var tokens map[string]DatabaseToken
var mu *sync.Mutex

func init() {
	once.Do(func() {
		fmt.Println("init tokens map")

		tokens = make(map[string]DatabaseToken)
		mu = &sync.Mutex{}
	})
}

func SetToken(id string) (string, time.Time, error) {
	defer mu.Unlock()

	token := GenerateToken()

	mu.Lock()

	// генерю токен, пока не найдет неиспользованный
LOOP:
	for {
		if _, ok := tokens[token]; !ok {
			break LOOP
		}
	}

	now := time.Now()
	tokens[token] = DatabaseToken{
		UserId:            id,
		CookieExpiredTime: now.Add(config.Get().CookieConfig.CookieTimeHours.Duration),
		//CookieIssuedTime:  now
	}

	return token, now.Add(config.Get().CookieConfig.CookieTimeHours.Duration), nil
}

func UpdateToken(token string) (DatabaseToken, error) {
	defer mu.Unlock()

	mu.Lock()

	// updating values in map not via ptrs
	tmpToken := tokens[token]
	tmpToken.CookieExpiredTime = time.Now().Add(config.Get().CookieConfig.CookieTimeHours.Duration)
	tokens[token] = tmpToken

	return tokens[token], nil
}

// при взятии токена, проверяет его на время
func GetId(token string) (string, error) {
	defer mu.Unlock()
	mu.Lock()

	if i, ok := tokens[token]; !ok {
		return "", errors.New(NoTokenFound)
	} else {
		if i.CookieExpiredTime.Unix() < time.Now().Unix() {
			return "", errors.New("your's session has been expired, relogin please")
		} else {
			return i.UserId, nil
		}
	}
}

func DeleteToken(token string) error {
	defer mu.Unlock()
	mu.Lock()

	if _, ok := tokens[token]; !ok {
		return errors.New(NoTokenFound)
	}

	delete(tokens, token)

	return nil
}
