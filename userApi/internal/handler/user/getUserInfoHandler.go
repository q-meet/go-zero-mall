package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/mall/user/Api/internal/logic/user"
	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)

		//result.HttpResult(r, w, resp, err)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

//http返回
//func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
//
//	if err == nil {
//		//成功返回
//		r := Success(resp)
//		httpx.WriteJson(w, http.StatusOK, r)
//	} else {
//		//错误返回
//		errcode := xerr.SERVER_COMMON_ERROR
//		errmsg := "服务器开小差啦，稍后再来试一试"
//
//		causeErr := errors.Cause(err)                // err类型
//		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
//			//自定义CodeError
//			errcode = e.GetErrCode()
//			errmsg = e.GetErrMsg()
//		} else {
//			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
//				grpcCode := uint32(gstatus.Code())
//				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
//					errcode = grpcCode
//					errmsg = gstatus.Message()
//				}
//			}
//		}
//
//		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
//
//		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
//	}
//}
