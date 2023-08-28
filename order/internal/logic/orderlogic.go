package logic

import (
	"context"
	"fmt"
	"go-zero/mall/user/types/user"
	"time"

	"go-zero/mall/order/internal/svc"
	"go-zero/mall/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	saveUser, err := l.svcCtx.UserRpc.SaveUser(l.ctx, &user.UserRequest{
		Name:   fmt.Sprintf("yy_%d_%s", time.Now().Unix(), req.Name),
		Gender: "woman",
	})
	if err != nil {
		return nil, err
	}
	l.Logger.Info("__________")
	l.Logger.Info(saveUser)
	return &types.Response{
		Message: fmt.Sprintf("id:%s,name:%s,gender:%s", saveUser.Id, saveUser.Name, saveUser.Gender),
	}, nil
}
