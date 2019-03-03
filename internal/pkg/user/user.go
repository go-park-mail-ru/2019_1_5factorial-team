package user

type User struct {
	Email string
	Nickname string
	Password string
	Score int
	AvatarType string
	AvatarLink string
}

func CreateUser(nickname string, email string, password string) error {
	// TODO(smet1): добавить валидацию
	addUser(DatabaseUser{
		Email: email,
		Nickname: nickname,
		Password: password,
	})

	return nil
}


