package session

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
	"time"
)

const (
	CookieName      string = "token"
	HttpOnly        bool   = true
	CookieTimeHours        = 10
	ServerPrefix    string = "/api"
)

//// https://robreid.io/json-time-duration/
//type ConfigDuration struct {
//	time.Duration
//}
//
//func (d *ConfigDuration) UnmarshalJSON(b []byte) (err error) {
//	d.Duration, err = time.ParseDuration(strings.Trim(string(b), `"`))
//	return
//}
//
//func (d ConfigDuration) MarshalJSON() (b []byte, err error) {
//	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
//}

//type CookieConfig struct {
//	CookieName      string         `json:"cookie_name"`
//	HttpOnly        bool           `json:"http_only"`
//	CookieDuration  int64          `json:"cookie_time_hours"`
//	ServerPrefix    string         `json:"server_prefix"`
//	CookieTimeHours ConfigDuration `json:"cookie_time"`
//}

//var CookieConf = CookieConfig{}

//func init() {
//	err := config_reader.ReadConfigFile("cookie_config.json", &CookieConf)
//	if err != nil {
//		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
//	}
//	fmt.Println(CookieConf)
//
//	// чтобы не копипастить везде (это время жизни куки в часах, чтобы добавлять к уже созданной)
//	CookieConf.CookieTimeHours = time.Duration(CookieConf.CookieDuration * int64(time.Hour))
//}

func CreateHttpCookie(value string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Path:     config.GetInstance().CookieConfig.ServerPrefix,
		Name:     config.GetInstance().CookieConfig.CookieName,
		Value:    value,
		Expires:  expiration,
		HttpOnly: config.GetInstance().CookieConfig.HttpOnly,
	}
}

func UpdateHttpCookie(cookie *http.Cookie, expiration time.Time) {
	cookie.Path = config.GetInstance().CookieConfig.ServerPrefix
	cookie.Name = config.GetInstance().CookieConfig.CookieName
	cookie.Expires = expiration
	cookie.HttpOnly = config.GetInstance().CookieConfig.HttpOnly
}
