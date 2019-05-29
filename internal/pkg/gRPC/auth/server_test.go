package session

import (
	"context"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"testing"
)

func TestGRPCServer(t *testing.T) {
	config.Get().AuthGRPCConfig.Port = "1111111111111111111111"
	err := GRPCServer()

	if err == nil {
		t.Error("err expected")
	}

	config.Get().AuthGRPCConfig.Port = "5000"
	go GRPCServer()
}

//func TestCreateConnection(t *testing.T) {
//	config.Get().AuthGRPCConfig.Port = "1111111111111111111111"
//	_, err := CreateConnection()
//	if err == nil {
//		t.Error("err not expected")
//	}
//}

func TestAuth_CreateSession(t *testing.T) {
	grpcConn, err := CreateConnection()
	if err != nil {
		log.Error(errors.Wrap(err, "cant create auth service client, GRPC.InitAuthClient"))
	}

	AuthGRPCClient := NewAuthCheckerClient(grpcConn)

	ctx := context.Background()
	_, err = AuthGRPCClient.CreateSession(ctx, &UserID{
		ID: "kek",
	})
	if err != nil {
		t.Error("error expected")
	}
}
