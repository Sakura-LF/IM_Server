package handler

import (
	"net/http"

	"IM_Server/im_auth/auth_api/internal/logic"
	"IM_Server/im_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.Authentication(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err, func(w http.ResponseWriter, err error) {
				httpx.WriteJsonCtx(r.Context(), w, resp.Code, resp)
			})
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
