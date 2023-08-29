package logic

import (
	"context"
	"go-zero/mall/user/types/user"

	"go-zero/mall/order/internal/svc"
	"go-zero/mall/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserReply{
		Id:       userResponse.Id,
		UserName: userResponse.Name,
		Gender:   userResponse.Gender,
	}, nil

}
