package logic

import (
	"context"
	"go-zero/mall/user/internal/model"
	"strconv"

	"go-zero/mall/user/internal/svc"
	"go-zero/mall/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserLogic {
	return &SaveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	data := &model.User{
		Name:   in.GetName(),
		Gender: in.GetGender(),
	}
	err := l.svcCtx.UserRepo.Save(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     strconv.FormatInt(data.Id, 10),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}
