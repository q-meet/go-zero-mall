package login

import (
	"net/http"
	"rpc-common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/mall/user/Api/internal/logic/login"
	"go-zero/mall/user/Api/internal/svc"
	"go-zero/mall/user/Api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			// 错误类型不对
			result.ParamErrorResult(r, w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
