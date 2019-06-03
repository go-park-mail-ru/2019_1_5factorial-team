package log

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/rossmcdonald/telegram_hook"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
	"math/rand"
	"net/http"
	"os"
)

var (
	Fatal   = logrus.Fatalf
	Warn    = logrus.Warn
	Printf  = logrus.Printf
	Println = logrus.Println

	Error  = logrus.Error
	Errorf = logrus.Errorf
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
		// socks5 proxy for telegram
		httpTransport := &http.Transport{}
		httpClient := &http.Client{Transport: httpTransport}
		dialer, err := proxy.SOCKS5(
			config.Get().LogrusConfig.ProxyNetwork,
			fmt.Sprintf("%s:%s", config.Get().LogrusConfig.ProxyIP, config.Get().LogrusConfig.ProxyPort),
			nil,
			proxy.Direct,
		)
		httpTransport.Dial = dialer.Dial

		// тележка <3
		hook, err := telegram_hook.NewTelegramHookWithClient(
			config.Get().LogrusConfig.AppName,
			config.Get().LogrusConfig.AuthToken,
			config.Get().LogrusConfig.TargetID,
			httpClient,
			telegram_hook.WithAsync(config.Get().LogrusConfig.Async),
			telegram_hook.WithTimeout(config.Get().LogrusConfig.Timeout.Duration),
		)
		if err != nil {
			logrus.Fatalf("Encountered error when creating Telegram hook: %s", err)
		}
		logrus.AddHook(hook)

		logrus.Warnf("logrus using telegram hook, acc = %s", config.Get().LogrusConfig.AppName)
		logrus.Errorf("hello %s! I'm alive and ready to sent you errors", config.Get().LogrusConfig.TargetID)
	}
}

func LoggerWithoutAuth(req *http.Request) *logrus.Entry {
	ctxLogger := logrus.WithFields(logrus.Fields{
		"ID_log":     generateID(),
		"req":        req.URL,
		"method":     req.Method,
		"host":       req.Host,
		"remoteAddr": req.RemoteAddr,
	})

	return ctxLogger
}

func LoggerWithAuth(req *http.Request) *logrus.Entry {
	ctxLogger := logrus.WithFields(logrus.Fields{
		"ID_log":     generateID(),
		"req":        req.URL,
		"method":     req.Method,
		"host":       req.Host,
		"remoteAddr": req.RemoteAddr,
		"userID":     req.Context().Value("userID"),
		"auth":       req.Context().Value("authorized"),
	})

	return ctxLogger
}

func generateID() interface{} {
	return rand.Intn(1000)
}
