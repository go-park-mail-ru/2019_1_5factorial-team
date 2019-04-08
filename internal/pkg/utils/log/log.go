package log

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/rossmcdonald/telegram_hook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	Fatal = logrus.Fatalf
	Warn  = logrus.Warn
)

func InitLogs() {
	// настраиваем logrus (по всему проекту)
	logrus.SetOutput(os.Stdout)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   config.Get().LogrusConfig.DisableColors,
		FullTimestamp:   config.Get().LogrusConfig.FullTimestamp,
		TimestampFormat: config.Get().LogrusConfig.TimestampFormat,
	})
	//log.SetFormatter(&log.JSONFormatter{
	//	TimestampFormat: config.Get().LogrusConfig.TimestampFormat,
	//	PrettyPrint: true,
	//})

	if config.Get().LogrusConfig.AppName != "" {
		// тележка <3
		hook, err := telegram_hook.NewTelegramHook(
			config.Get().LogrusConfig.AppName,
			config.Get().LogrusConfig.AuthToken,
			config.Get().LogrusConfig.TargetID,
			telegram_hook.WithAsync(config.Get().LogrusConfig.Async),
			telegram_hook.WithTimeout(config.Get().LogrusConfig.Timeout.Duration),
		)
		if err != nil {
			logrus.Fatalf("Encountered error when creating Telegram hook: %s", err)
		}
		logrus.AddHook(hook)
	}
}

func LoggerWithoutAuth(funcName string, req *http.Request) *logrus.Entry {
	ctxLogger := logrus.WithFields(logrus.Fields{
		"req":        req.URL,
		"method":     req.Method,
		"host":       req.Host,
		"remoteAddr": req.RemoteAddr,
		"func":       funcName,
	})

	return ctxLogger
}

func LoggerWithAuth(funcName string, req *http.Request) *logrus.Entry {
	ctxLogger := logrus.WithFields(logrus.Fields{
		"req":        req.URL,
		"method":     req.Method,
		"host":       req.Host,
		"remoteAddr": req.RemoteAddr,
		"func":       funcName,
		"userID":     req.Context().Value("userID"),
		"auth":       req.Context().Value("authorized"),
	})

	return ctxLogger
}
