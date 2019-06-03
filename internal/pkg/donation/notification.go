package donation

import (
	"crypto/sha1"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"net/http"
)

// extended info from api request
type NotificationExtendedInfo struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// struct with info, available on notification
type Notification struct {
	OperationID string
	Amount      int
	NotificationExtendedInfo
}

func checkNotification(req *http.Request) bool {
	form := req.PostForm
	checkString := fmt.Sprintf("%s&%s&%s&%s&%s&%s&%s&%s&%s",
		form.Get("notification_type"),
		form.Get("operation_id"),
		form.Get("amount"),
		form.Get("currency"),
		form.Get("datetime"),
		form.Get("sender"),
		form.Get("codepro"),
		config.Get().DonationConfig.NotificationSecret,
		form.Get("label"),
	)

	originalHash := form.Get("sha1_hash")

	calculatedHash := checkNotificationString(checkString)

	return originalHash == calculatedHash
}

func checkNotificationString(checkString string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(checkString)))
}
