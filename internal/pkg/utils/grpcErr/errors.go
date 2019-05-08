package grpcErr

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

func GetHTTPStatus(grpcStatusCode codes.Code) int {
	switch grpcStatusCode {
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.AlreadyExists:
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
