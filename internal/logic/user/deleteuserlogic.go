package userlogic

import (
	"context"
	"errors"
	model "github.com/lambertstu/shortlink-user-rpc/mongo/user"
	"github.com/lambertstu/shortlink-user-rpc/pkg/constant"
	"github.com/lambertstu/shortlink-user-rpc/pkg/errorcode"

	"github.com/lambertstu/shortlink-user-rpc/internal/svc"
	"github.com/lambertstu/shortlink-user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	updateRes, err := l.svcCtx.UserModel.UpdateUserInfo(l.ctx, &model.User{
		Username:   in.Username,
		DeleteFlag: constant.DEL_FLAG,
	})

	if err != nil || updateRes.MatchedCount == 0 || updateRes.ModifiedCount == 0 {
		if updateRes.MatchedCount == 0 {
			return &user.DeleteUserResponse{
				Success: false,
			}, errors.New(errorcode.UserNotExist.Message())
		} else if updateRes.ModifiedCount == 0 {
			return &user.DeleteUserResponse{
				Success: false,
			}, errors.New("删除信息失败")
		}

		return &user.DeleteUserResponse{
			Success: false,
		}, err
	}
	return &user.DeleteUserResponse{
		Success: true,
	}, nil
}
