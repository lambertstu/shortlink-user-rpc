// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"github.com/lambertstu/shortlink-user-rpc/internal/logic/user"
	"github.com/lambertstu/shortlink-user-rpc/internal/svc"
	"github.com/lambertstu/shortlink-user-rpc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := userlogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := userlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Logout(ctx context.Context, in *user.LogoutRequest) (*user.LogoutResponse, error) {
	l := userlogic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *UserServer) GetUserByUsername(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {
	l := userlogic.NewGetUserByUsernameLogic(ctx, s.svcCtx)
	return l.GetUserByUsername(in)
}

func (s *UserServer) IsExistUser(ctx context.Context, in *user.IsExistUserRequest) (*user.IsExistUserResponse, error) {
	l := userlogic.NewIsExistUserLogic(ctx, s.svcCtx)
	return l.IsExistUser(in)
}

func (s *UserServer) IsUserLogin(ctx context.Context, in *user.IsUserLoginRequest) (*user.IsUserLoginResponse, error) {
	l := userlogic.NewIsUserLoginLogic(ctx, s.svcCtx)
	return l.IsUserLogin(in)
}

func (s *UserServer) UpsertUserInfo(ctx context.Context, in *user.UpsertUserInfoRequest) (*user.UpsertUserInfoResponse, error) {
	l := userlogic.NewUpsertUserInfoLogic(ctx, s.svcCtx)
	return l.UpsertUserInfo(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	l := userlogic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}
