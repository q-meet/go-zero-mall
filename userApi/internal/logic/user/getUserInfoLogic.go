package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero/mall/user/types/user"
	"strconv"

	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("userId")
	logx.Infof("get token content: %s \n ", value)
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: strconv.FormatInt(req.Id, 10),
	})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		Message: "success",
		Data: types.GetUserInfo{
			Id:   req.Id,
			Name: userResponse.Name,
			Desc: userResponse.Gender,
		},
	}, nil
}

func (l *GetUserInfoLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
