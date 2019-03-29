package config

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/config_reader"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/fileproc"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/pkg/errors"
	"log"
)
var instance *ServerConfig

type ServerConfig struct {
	StaticServerConfig fileproc.StaticServerConfig
	CORSConfig         middleware.CORSConfig
	CookieConfig       session.CookieConfig

	configPath string
}

func (sc *ServerConfig) New(configsDir string) *ServerConfig {
	sc.configPath = configsDir

	// конфиг статик сервера
	err := config_reader.ReadConfigFile(configsDir, "static_server_config.json", &sc.StaticServerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading static_server_config config"))
	}
	sc.StaticServerConfig.MaxUploadSize = sc.StaticServerConfig.MaxUploadSizeMB * 1024 * 1024

	log.Println("New Server->Static server config = ", sc.StaticServerConfig)

	// конфиг корса
	err = config_reader.ReadConfigFile(configsDir, "cors_config.json", &sc.CORSConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading CORS config"))
	}

	log.Println("New Server->CORS config = ", sc.CORSConfig)

	// конфиг кук
	err = config_reader.ReadConfigFile(configsDir, "cookie_config.json", &sc.CookieConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
	}
	log.Println("New Server->Cookie config = ", sc.CookieConfig)

	// инстанс сервера
	instance = sc

	return sc
}

func GetInstance() *ServerConfig {
	return instance
}