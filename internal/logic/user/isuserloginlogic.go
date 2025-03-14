package userlogic

import (
	"context"
	"github.com/lambertstu/shortlink-user-rpc/pkg/constant"

	"github.com/lambertstu/shortlink-user-rpc/internal/svc"
	"github.com/lambertstu/shortlink-user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserLoginLogic {
	return &IsUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsUserLoginLogic) IsUserLogin(in *user.IsUserLoginRequest) (*user.IsUserLoginResponse, error) {
	hasLoginMap, _ := l.svcCtx.Redis.Hgetall(constant.USER_LOGIN_KEY + in.GetUsername())
	if len(hasLoginMap) == 0 || hasLoginMap["token"] != in.Token {
		return &user.IsUserLoginResponse{
			Success: false,
		}, nil
	}
	return &user.IsUserLoginResponse{
		Success: true,
	}, nil
}
