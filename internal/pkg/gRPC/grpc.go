package gRPC

import (
	"context"
	auth "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
)

func InitAuthClient() error {
	grpcConn, err := auth.CreateConnection()
	if err != nil {
		log.Error(errors.Wrap(err, "cant create auth service client, GRPC.InitAuthClient"))

		return errors.Wrap(err, "cant create auth service client, GRPC.InitAuthClient")
	}

	auth.AuthGRPCClient = auth.NewAuthCheckerClient(grpcConn)

	_, err = auth.AuthGRPCClient.Ping(context.Background(), &auth.Nothing{})
	if err != nil {
		log.Error(errors.Wrap(err, "cant ping auth service client, GRPC.InitAuthClient"))

		return errors.Wrap(err, "cant ping auth service client, GRPC.InitAuthClient")
	}

	return nil
}
