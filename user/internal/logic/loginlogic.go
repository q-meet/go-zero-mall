package logic

import (
	"context"
	"errors"
	"go-zero/mall/user/internal/svc"
	"go-zero/mall/user/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	if in.Name == "yy_1693215703_me" {
		return nil, errors.New("id参数不正确")
	}
	userData, err := l.svcCtx.UserRepo.FindByName(context.Background(), in.Name)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     strconv.FormatInt(userData.Id, 10),
		Name:   userData.Name,
		Gender: userData.Gender,
	}, nil
}
