package handler

import (
	"net/http"

	"IM_Server/im_auth/auth_api/internal/logic"
	"IM_Server/im_auth/auth_api/internal/svc"
	"IM_Server/im_auth/auth_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			//httpx.WriteJson(w, http.StatusBadRequest)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		//response.Response(r, w, resp, err)
	}
}
