package interceptor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rpc-common/errorx"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			//转成 grpc err
			err = status.Error(codes.Code(e.GetErrCode()), err.Error()) // e.GetErrMsg())
		} else if e, ok := causeErr.(*errorx.BizError); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			//转成 grpc err
			err = status.Error(codes.Code(e.GetCode()), err.Error()) // e.GetErrMsg())
		} else {
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	/*
		resp, err = handler(ctx, req)
		if err != nil {
			causeErr := errors.Cause(err)                  // err类型
			if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
				logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
				logx.Infof("【RPC-SRV-ERR】 %+v", err)
				//转成 grpc err
				err = status.Error(codes.Code(e.GetErrCode()), err.Error()) // e.GetErrMsg())
			} else {
				logx.Infof("【RPC-SRV-ERR】 %+v", err)
				logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			}
		}
	*/
	return resp, err
}
