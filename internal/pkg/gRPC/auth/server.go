package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const (
	address = "auth-go"
	port    = "5000"
)

var AuthGRPCClient AuthCheckerClient
var once sync.Once

type Auth struct {
}

func GetClient() AuthCheckerClient {
	once.Do(func() {
		grpcConn, err := CreateConnection()
		if err != nil {
			logrus.Error(errors.Wrap(err, "cant create auth service client"))

			return
		}

		AuthGRPCClient = NewAuthCheckerClient(grpcConn)
	})

	return AuthGRPCClient
}

func CreateConnection() (*grpc.ClientConn, error) {
	grcpConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", address, port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Error(errors.Wrap(err, "cant connect to grpc"))

		return nil, errors.Wrap(err, "cant connect to grpc")
	}
	//defer grcpConn.Close()

	return grcpConn, nil
}

func GRPCServer() error {
	port := "5000"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Error(err)
		return err
	}
	grpcServer := grpc.NewServer()
	RegisterAuthCheckerServer(grpcServer, &Auth{})

	log.Println(fmt.Sprintf("start serve auth grpc in %s port", port))
	return grpcServer.Serve(lis)
}

func (a *Auth) CreateSession(ctx context.Context, userID *UserID) (*Cookie, error) {
	randToken, expiration, err := session.SetToken(userID.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Auth gRPC")
	}

	return &Cookie{
		Token:      randToken,
		Expiration: expiration.Format(time.RFC3339),
	}, nil
}

func (a *Auth) DeleteSession(ctx context.Context, cookie *Cookie) (*Nothing, error) {
	err := session.DeleteToken(cookie.Token)
	if err != nil {
		return &Nothing{}, err
	}
	return &Nothing{}, nil
}

func (a *Auth) UpdateSession(ctx context.Context, cookie *Cookie) (*Cookie, error) {
	updatedToken, err := session.UpdateToken(cookie.Token)
	if err != nil {
		return nil, errors.Wrap(err, "cant update token in auth grpc")
	}
	return &Cookie{
		Token:      updatedToken.Token,
		Expiration: updatedToken.CookieExpiredTime.Format(time.RFC3339),
	}, nil
}

func (a *Auth) GetIDFromSession(ctx context.Context, cookie *Cookie) (*UserID, error) {
	uId, err := session.GetId(cookie.Token)
	if err != nil {
		return &UserID{}, errors.Wrap(err, "cant find session in auth grpc")
	}

	return &UserID{ID: uId}, nil
}

// getid from session
