package session

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/app/config"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/session"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2"
	"net"
	"strings"
	"time"
)

const (
	address           = "auth-go"
	port              = "5000"
	MongoConflictCode = 11000
)

var AuthGRPCClient AuthCheckerClient

type Auth struct {
}

func CreateConnection() (*grpc.ClientConn, error) {
	grcpConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", config.Get().AuthGRPCConfig.Hostname, config.Get().AuthGRPCConfig.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Error(errors.Wrap(err, "cant connect to grpc"))

		return nil, errors.Wrap(err, "cant connect to grpc")
	}

	return grcpConn, nil
}

func GRPCServer() error {
	port := config.Get().AuthGRPCConfig.Port
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
		return nil, errors.Wrap(err, "Auth gRPC, AuthGRPC")
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
		return nil, errors.Wrap(err, "cant update token in auth grpc, AuthGRPC")
	}

	return &Cookie{
		Token:      updatedToken.Token,
		Expiration: updatedToken.CookieExpiredTime.Format(time.RFC3339),
	}, nil
}

func (a *Auth) GetIDFromSession(ctx context.Context, cookie *Cookie) (*UserID, error) {
	uId, err := session.GetId(cookie.Token)
	if err != nil {
		return &UserID{}, errors.Wrap(err, "cant find session in auth grpc, AuthGRPC")
	}

	return &UserID{ID: uId}, nil
}

func (a *Auth) CreateUser(ctx context.Context, data *UserNew) (*User, error) {
	u, err := user.CreateUser(data.Nickname, data.Email, data.Password, data.Avatar)
	if err != nil {
		log.Error(errors.Wrap(err, "err in user data, AuthGRPC"))

		if errors.Cause(err).(*mgo.LastError).Code == MongoConflictCode {
			if strings.Contains(errors.Cause(err).(*mgo.LastError).Err, data.Nickname) {
				return &User{}, status.Error(codes.AlreadyExists, "login conflict")

			} else if strings.Contains(errors.Cause(err).(*mgo.LastError).Err, data.Email) {
				return &User{}, status.Error(codes.AlreadyExists, "email conflict")
			}
		}

		return &User{}, status.Error(codes.Internal, err.Error())
	}

	return &User{
		ID:           u.ID.Hex(),
		Email:        u.Email,
		Nickname:     u.Nickname,
		HashPassword: u.HashPassword,
		Score:        int32(u.Score),
		AvatarLink:   u.AvatarLink,
	}, nil
}

func (a *Auth) IdentifyUser(ctx context.Context, data *DataAuth) (*User, error) {
	u, err := user.IdentifyUser(data.Login, data.Password)
	if err != nil {
		log.Error(errors.Wrap(err, "Wrong password or login, AuthGRPC"))

		return &User{}, err
	}

	return &User{
		ID:           u.ID.Hex(),
		Email:        u.Email,
		Nickname:     u.Nickname,
		HashPassword: u.HashPassword,
		Score:        int32(u.Score),
		AvatarLink:   u.AvatarLink,
	}, nil
}

func (a *Auth) GetUserByID(ctx context.Context, u *User) (*User, error) {
	searchingUser, err := user.GetUserById(u.ID)
	if err != nil {
		log.Error(errors.Wrap(err, "user with this id not found, AuthGRPC"))
		return &User{}, err
	}

	return &User{
		ID:           searchingUser.ID.Hex(),
		Email:        searchingUser.Email,
		Nickname:     searchingUser.Nickname,
		HashPassword: searchingUser.HashPassword,
		Score:        int32(searchingUser.Score),
		AvatarLink:   searchingUser.AvatarLink,
	}, nil
}

func (a *Auth) UpdateUser(ctx context.Context, req *UpdateUserReq) (*Nothing, error) {
	// TODO(): сделать ответ, как в CreateUser (status.Code, status.Message)

	//if req.NewPassword != "" {
	//	flagValidNewPassword := validator.ValidUpdatePassword(req.NewPassword)
	//	if !flagValidNewPassword {
	//		//ErrResponse(res, http.StatusBadRequest, "invalid new password")
	//		log.Error(errors.New("invalid new password"))
	//		return &Nothing{}, errors.New("invalid new password")
	//	}
	//}

	err := user.UpdateUser(req.ID, req.NewAvatar, req.OldPassword, req.NewPassword)
	if err != nil {
		log.Error(errors.Wrap(err, "UpdateUser error, AuthGRPC"))

		return &Nothing{}, err
	}

	return &Nothing{}, nil
}

func (a *Auth) GetUsersScores(ctx context.Context, params *ScoresParam) (*Scores, error) {
	leaderboard, err := user.GetUsersScores(int(params.Limit), int(params.Offset))
	if err != nil {
		log.Error(errors.Wrap(err, "get leaderboard error, AuthGRPC"))

		return &Scores{}, err
	}

	scores := make([]*Score, 0, len(leaderboard))
	for _, val := range leaderboard {
		scores = append(scores, &Score{Nickname: val.Nickname, Score: int32(val.Score)})
	}

	return &Scores{Scores: scores}, nil
}

func (a *Auth) GetUsersCount(ctx context.Context, _ *Nothing) (*Num, error) {
	count, err := user.GetUsersCount()
	if err != nil {
		log.Error(err.Error())

		return &Num{}, err
	}

	return &Num{
		Count: int32(count),
	}, nil
}

func (a *Auth) UpdateScore(ctx context.Context, req *UpdateScoreReq) (*Nothing, error) {
	err := user.UpdateScore(req.ID, int(req.Score))
	if err != nil {
		log.Error(errors.Wrap(err, "update score err, AuthGRPC"))

		return &Nothing{}, err
	}

	return &Nothing{}, nil
}

func (a *Auth) Ping(ctx context.Context, _ *Nothing) (*Nothing, error) {
	log.Warn("grpc auth alive")
	// тут еще можно проверять доступность баз

	return &Nothing{}, nil
}
