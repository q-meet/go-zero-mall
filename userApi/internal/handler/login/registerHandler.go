package login

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"rpc-common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/mall/user/Api/internal/logic/login"
	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			// 错误类型不对
			logx.WithContext(r.Context()).Error("参数")
			result.ParamErrorResult(r, w, err)
			return
		}

		l := login.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)

		result.HttpResult(r, w, resp, err)
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}*/
	}
}
