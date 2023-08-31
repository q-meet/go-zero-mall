package logic

import (
	"context"

	"go-zero/mall/userscore/internal/svc"
	"go-zero/mall/userscore/types/userscore"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(in *userscore.UserScoreRequest) (*userscore.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	l.WithContext(l.ctx).Infof("SaveUserScore:%d", in.UserId)
	return &userscore.UserScoreResponse{
		UserId: in.UserId,
		Score:  in.Score + 1,
		Name:   "Name",
	}, nil
}
