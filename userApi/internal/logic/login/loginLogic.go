package login

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero/mall/user/types/user"
	"strconv"
	"time"

	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userInfo, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{Name: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	// todo: add your logic here and delete this line
	auth := l.svcCtx.Config.Auth
	userId, _ := strconv.ParseInt(userInfo.Id, 10, 64)
	result, err := l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, userId)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Id:       userId,
		Name:     userInfo.Name,
		Token:    result,
		ExpireAt: strconv.FormatInt(auth.AccessExpire, 10),
	}, nil
}

func (l *LoginLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
