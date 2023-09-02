package interceptor

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rpc-common/errorx"
	"rpc-common/trace"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	logger := trace.WithContext(ctx)

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			logger.Errorf("【RPC-SRV-ERR】 %+v", err)
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			//转成 grpc err
			err = status.Error(codes.Code(e.GetErrCode()), err.Error()) // e.GetErrMsg())
		} else if e, ok := causeErr.(*errorx.BizError); ok { //自定义错误类型
			logger.Errorf("【RPC-SRV-ERR】 %+v", err)
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			//转成 grpc err
			err = status.Error(codes.Code(e.GetCode()), err.Error()) // e.GetErrMsg())
		} else {
			//logx.Infof("【RPC-SRV-ERR】 %+v", err)
			logger.Errorf("【RPC-SRV-ERR】 %+v", err)
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

func ExampleUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// TODO: fill your logic here
	return handler(ctx, req)
}
func ExampleStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// TODO: fill your logic here
	return handler(srv, ss)
}
