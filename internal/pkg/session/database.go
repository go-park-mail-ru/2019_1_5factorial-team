package session

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

//type TokenDatabase struct {
//	Token string
//	UserId int64
//}

var once sync.Once
var tokens map[string]int64
var mu *sync.Mutex

func init() {
	once.Do(func() {
		fmt.Println("init tokens map")

		tokens = make(map[string]int64)
		mu = &sync.Mutex{}
	})
}

func SetToken(id int64) (string, error) {
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

	tokens[token] = id

	return token, nil
}

func GetId(token string) (int64, error) {
	defer mu.Unlock()
	mu.Lock()

	if i, ok := tokens[token]; !ok {
		return 0, errors.New("token not found")
	} else {
		return i, nil
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