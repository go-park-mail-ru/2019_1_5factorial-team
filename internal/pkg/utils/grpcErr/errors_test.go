package grpcErr

import (
	"google.golang.org/grpc/codes"
	"net/http"
	"testing"
)

var casesGetHTTPStatus = []struct {
	status codes.Code
	want   int
}{
	{
		status: codes.Internal,
		want:   http.StatusInternalServerError,
	},
	{
		status: codes.AlreadyExists,
		want:   http.StatusConflict,
	},
	{
		status: codes.DataLoss,
		want:   http.StatusInternalServerError,
	},
}

func TestGetHTTPStatus(t *testing.T) {
	for i, val := range casesGetHTTPStatus {
		res := GetHTTPStatus(val.status)

		if res != val.want {
			t.Error("#", i, "RES expected:", val.want, "have:", res)
			continue
		}
	}
}
