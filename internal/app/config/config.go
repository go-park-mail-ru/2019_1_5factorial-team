package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/pkg/errors"
)

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
type DBConfig struct {
	Hostname       string `json:"hostname"`
	MongoPort      string `json:"mongo_port"`
	DatabaseName   string `json:"database_name"`
	CollectionName string `json:"collection_name"`
	TruncateTable  bool   `json:"truncate_table"`
}

// TODO(): есть смысл объединить в 1 файл конфига
// структура сервера, собирает все вышеперечисленные структуры
type ServerConfig struct {
	StaticServerConfig StaticServerConfig
	CORSConfig         CORSConfig
	CookieConfig       CookieConfig
	DBConfig           []DBConfig

	configPath string
}

var instance = &ServerConfig{}

// откуда читать, куда заносить
type valueAndPath struct {
	from string
	to   interface{}
}

var configs = []valueAndPath{
	{
		from: "static_server_config.json",
		to:   &instance.StaticServerConfig,
	},
	{
		from: "cors_config.json",
		to:   &instance.CORSConfig,
	},
	{
		from: "cookie_config.json",
		to:   &instance.CookieConfig,
	},
	{
		from: "db_config.json",
		to:   &instance.DBConfig,
	},
}

// считывание всех конфигов по пути `configsDir`
func Init(configsDir string) error {
	log.Println("Configs->logs path = ", configsDir)

	for i, val := range configs {
		err := config_reader.ReadConfigFile(configsDir, val.from, val.to)
		if err != nil {
			return errors.Wrap(err, "error while reading config")
		}

		log.Println("Configs->", i, "config = ", val.to)
	}

	return nil
}

func Get() *ServerConfig {
	return instance
}
