package grouplogic

import (
	"context"
	"github.com/lambertstu/shortlink-user-rpc/internal/svc"
	"github.com/lambertstu/shortlink-user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByUsernameLogic {
	return &GetGroupByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupByUsernameLogic) GetGroupByUsername(in *user.GetGroupRequest) (*user.GetGroupResponse, error) {
	groups, err := l.svcCtx.GroupModel.GetGroupByUsername(l.ctx, in.GetUsername())
	if err != nil || groups == nil {
		return &user.GetGroupResponse{}, err
	}

	var groupDataList []*user.GroupData
	for _, group := range *groups {
		groupData := &user.GroupData{
			Gid:      group.GID,
			Name:     group.Name,
			Username: group.Username,
		}
		groupDataList = append(groupDataList, groupData)
	}

	return &user.GetGroupResponse{
		Data: groupDataList,
	}, nil
}
