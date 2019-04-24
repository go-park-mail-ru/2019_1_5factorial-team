package auth

import (
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/gRPC/auth"
	"github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/utils/log"
	"github.com/pkg/errors"
)

func Run()  {
	err := session.GRPCServer()
	if err != nil {
		log.Error(errors.Wrap(err, "cant start auth service"))
	}
}
