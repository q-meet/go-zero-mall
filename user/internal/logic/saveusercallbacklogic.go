package logic

import (
	"context"

	"go-zero/mall/user/internal/svc"
	"go-zero/mall/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserCallbackLogic {
	return &SaveUserCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserCallbackLogic) SaveUserCallback(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line

	l.WithContext(l.ctx).Info("我收到了 saveUserCallback")
	l.WithContext(l.ctx).Infof("参数是：%#v", in)
	return &user.UserResponse{}, nil
}
