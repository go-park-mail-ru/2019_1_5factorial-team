package session

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type DatabaseToken struct {
	UserId            int64
	CookieIssuedTime  time.Time
	CookieExpiredTime time.Time
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

func SetToken(id int64) (string, time.Time, error) {
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
		CookieIssuedTime:  now,
		CookieExpiredTime: now.Add(cookieTimeHours * time.Hour),
	}

	return token, now.Add(cookieTimeHours * time.Hour), nil
}

func UpdateToken(token string) (DatabaseToken, error) {
	defer mu.Unlock()

	mu.Lock()

	// updating values in map not via ptrs
	tmpToken := tokens[token]
	tmpToken.CookieExpiredTime = time.Now().Add(cookieTimeHours * time.Minute)
	tokens[token] = tmpToken

	return tokens[token], nil
}

// при взятии токена, проверяет его на время
func GetId(token string) (int64, error) {
	defer mu.Unlock()
	mu.Lock()

	if i, ok := tokens[token]; !ok {
		return 0, errors.New("token not found")
	} else {
		if i.CookieExpiredTime.Unix() < time.Now().Unix() {
			return 0, errors.New("your's session has been expired, relogin please")
		} else {
			return i.UserId, nil
		}
	}
}

func DeleteToken(token string) error {
	defer mu.Unlock()
	mu.Lock()

	if _, ok := tokens[token]; !ok {
		return errors.New("token not found")
	}

	delete(tokens, token)

	return nil
}