package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero/mall/user/internal/svc"
	"go-zero/mall/user/types/user"
	"go.opentelemetry.io/otel/baggage"
	"rpc-common/errorx"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: WithContext(ctx),
		//Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	if in.Id == "1" {
		//return nil, errorx.ErrUserAlreadyRegisterError
		//return nil, errors.Wrapf(errorx.ErrUserAlreadyRegisterError, "用户已经存在 mobile:%s,err:%v", "in.Mobile, err", "22")
	}
	if in.Id == "1" {

		//return nil, errorx.ParamsError
		l.Logger.Infof(errors.Wrapf(errorx.ParamsError, "用户已经存在111 mobile:%s,err:%v", "in.Mobile, err", "22").Error())
		return nil, errors.Wrapf(errorx.ParamsError, "用户已经存在 mobile:%s,err:%v", "in.Mobile, err", "22")
	}
	id, _ := strconv.ParseInt(in.Id, 10, 64)
	userData, err := l.svcCtx.UserRepo.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   userData.Name,
		Gender: userData.Gender,
	}, nil
}

const traceIdKey = "biz-trace-id"

func FromTraceId(ctx context.Context) (string, bool) {
	bg := baggage.FromContext(ctx)
	member := bg.Member(traceIdKey)
	return member.Value(), member.Key() != ""
}

func WithContext(ctx context.Context) logx.Logger {
	traceId, ok := FromTraceId(ctx)
	if !ok {
		return logx.WithContext(ctx)
	}

	return logx.WithContext(ctx).WithFields(logx.Field(traceIdKey, traceId))
}
