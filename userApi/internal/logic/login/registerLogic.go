package login

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"
	"go-zero/mall/user/types/user"
	"go-zero/mall/userscore/types/userscore"
	"strconv"
	// 下面这行导入gozero的dtm驱动
	//_ "github.com/dtm-labs/driver-gozero"
)

// dtm已经通过前面的配置，注册到下面这个地址，因此在dtmgrpc中使用该地址
var dtmServer = "etcd://localhost:2379/dtmservice"

//var dtmServer = "http://localhost:36789/api/dtmsvr"

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	//消息型
	userRequest := &user.UserRequest{
		Name:   req.Username,
		Gender: req.Gender,
	}
	//gid := dtmgrpc.MustGenGid(dtmServer)
	//l.WithContext(l.ctx).Info("gid:" + gid)
	//msgGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	//userServer, err := l.svcCtx.Config.UserRpc.BuildTarget()
	//if err != nil {
	//	return nil, err
	//}
	//userScoreServer, err := l.svcCtx.Config.UserScoreRpc.BuildTarget()
	//if err != nil {
	//	return nil, err
	//}
	//msgGrpc.Add(userServer+"/user.User/SaveUser", userServer+"/user.User/saveUserCallback", userRequest)
	userResponse, err := l.svcCtx.UserRpc.SaveUser(l.ctx, userRequest)
	if err != nil {
		return nil, err
	}
	userId, _ := strconv.ParseInt(userResponse.Id, 10, 64)
	scoreRequest := &userscore.UserScoreRequest{
		UserId: 100,
		Score:  10,
	}
	//msgGrpc.Add(userScoreServer+"/userscore.Userscore/SaveUserScore", "", scoreRequest)
	score, err := l.svcCtx.UserScoreRpc.SaveUserScore(l.ctx, scoreRequest)
	if err != nil {
		return nil, err
	}
	//msgGrpc.WaitResult = true
	//err = msgGrpc.Submit()
	//if err != nil {
	//	fmt.Println("-----------------------")
	//	fmt.Println(err)
	//	return nil, errors.New(err.Error())
	//}
	logx.Infof("register add score %d,user_id %d", score.Score, userId)
	return &types.RegisterResp{
		Id:   userId,
		Name: req.Username,
	}, nil
}

/*
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	userResponse, err := l.svcCtx.UserRpc.SaveUser(l.ctx, &user.UserRequest{
		Name:   req.Username,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	userId, _ := strconv.ParseInt(userResponse.Id, 10, 64)
	scoreRequest := &userscore.UserScoreRequest{
		UserId: userId,
		Score:  10,
	}
	score, err := l.svcCtx.UserScoreRpc.SaveUserScore(l.ctx, scoreRequest)
	if err != nil {
		return nil, err
	}
	logx.WithContext(l.ctx).Infof("register add score %d", score.Score)
	return &types.RegisterResp{
		Id:   userId,
		Name: req.Username,
	}, nil
}*/
