package config

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/pkg/errors"
	"log"
	"strings"
	"time"
)

var instance *ServerConfig

// структура конфига статики
type StaticServerConfig struct {
	MaxUploadSizeMB int64  `json:"max_upload_size_mb"`
	UploadPath      string `json:"upload_path"`
	MaxUploadSize   int64
}

// структура конфига КОРС
type CORSConfig struct {
	Origin      string   `json:"allow-origin"`
	Credentials bool     `json:"allow-credentials"`
	Methods     []string `json:"allow-methods"`
	Headers     []string `json:"allow-headers"`
	MaxAge      int      `json:"max-age"`
}

// https://robreid.io/json-time-duration/
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	d.Duration, err = time.ParseDuration(strings.Trim(string(b), `"`))
	return
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}

// структура конфига кук
type CookieConfig struct {
	CookieName      string   `json:"cookie_name"`
	HttpOnly        bool     `json:"http_only"`
	CookieDuration  int64    `json:"cookie_time_hours"`
	ServerPrefix    string   `json:"server_prefix"`
	CookieTimeHours Duration `json:"cookie_time"`
}

// структура конфига базы юзеров
type DBUserConfig struct {
	MongoPort         string `json:"mongo_port"`
	DatabaseName      string `json:"database_name"`
	CollectionName    string `json:"collection_name"`
	GenerateFakeUsers bool   `json:"generate_fake_users"`
	TruncateTable     bool   `json:"truncate_table"`
}

// структура конфига генератора фейковых юзеров
type FakeUsersConfig struct {
	UsersCount int    `json:"users_count"`
	Lang       string `json:"lang"`
	MaxScore   int    `json:"max_score"`
}

// структура сервера, собирает все вышеперечисленные структуры
type ServerConfig struct {
	StaticServerConfig StaticServerConfig
	CORSConfig         CORSConfig
	CookieConfig       CookieConfig
	DBUserConfig       DBUserConfig
	FakeUsersConfig    FakeUsersConfig

	configPath string
}

// считывание всех конфигов по пути `configsDir`
func (sc *ServerConfig) New(configsDir string) *ServerConfig {
	sc.configPath = configsDir

	// конфиг статик сервера
	err := config_reader.ReadConfigFile(configsDir, "static_server_config.json", &sc.StaticServerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading static_server_config config"))
	}
	sc.StaticServerConfig.MaxUploadSize = sc.StaticServerConfig.MaxUploadSizeMB * 1024 * 1024

	log.Println("Configs->Static server config = ", sc.StaticServerConfig)

	// конфиг корса
	err = config_reader.ReadConfigFile(configsDir, "cors_config.json", &sc.CORSConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading CORS config"))
	}

	log.Println("Configs->CORS config = ", sc.CORSConfig)

	// конфиг кук
	err = config_reader.ReadConfigFile(configsDir, "cookie_config.json", &sc.CookieConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	}
	log.Println("Configs->Cookie config = ", sc.CookieConfig)

	// конфиг бд юзеров (монго)
	err = config_reader.ReadConfigFile(configsDir, "db_user_config.json", &sc.DBUserConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading DB User config"))
	}
	log.Println("Configs->DB User config = ", sc.DBUserConfig)

	// конфиг генерации фейковых юзеров
	err = config_reader.ReadConfigFile(configsDir, "user_faker_config.json", &sc.FakeUsersConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading user faker config"))
	}
	log.Println("Configs->User faker config = ", sc.FakeUsersConfig)

	// инстанс сервера
	instance = sc

	return sc
}

func GetInstance() *ServerConfig {
	return instance
}
