package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

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

// TODO(): вырезать CookieDuration здесь и в конфиге
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

// структура конфига хука под тележку и основного лога
type LogrusConfig struct {
	// тг
	AppName   string   `json:"app_name"`
	AuthToken string   `json:"auth_token"`
	TargetID  string   `json:"target_id"`
	Async     bool     `json:"async"`
	Timeout   Duration `json:"timeout"`

	// socks5 для тг
	ProxyNetwork string `json:"proxy_network"`
	ProxyIP      string `json:"proxy_ip"`
	ProxyPort    string `json:"proxy_port"`

	// логрус
	DisableColors   bool   `json:"disable_colors"`
	FullTimestamp   bool   `json:"full_timestamp"`
	TimestampFormat string `json:"timestamp_format"`
}

// TODO(): есть смысл объединить в 1 файл конфига
// структура сервера, собирает все вышеперечисленные структуры
type ServerConfig struct {
	StaticServerConfig StaticServerConfig
	CORSConfig         CORSConfig
	CookieConfig       CookieConfig
	DBConfig           []DBConfig
	LogrusConfig       LogrusConfig

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
	{
		from: "logrus_config.json",
		to:   &instance.LogrusConfig,
	},
}

// считывание всех конфигов по пути `configsDir`
func Init(configsDir string) error {
	logrus.WithField("func", "config.Init").Info("logs path = ", configsDir)

	for i, val := range configs {
		err := config_reader.ReadConfigFile(configsDir, val.from, val.to)
		if err != nil {
			logrus.WithField("err", err.Error()).Error("config.Init")

			return errors.Wrap(err, "error while reading config")
		}

		//log.Println("Configs->", i, "config = ", val.to)
		logrus.WithField("func", "config.Init").
			Infof("i = %d, from file = %s, config = %v", i, val.from, val.to)
	}

	instance.StaticServerConfig.MaxUploadSize = instance.StaticServerConfig.MaxUploadSizeMB * 1024 * 1024

	return nil
}

func Get() *ServerConfig {
	return instance
}
