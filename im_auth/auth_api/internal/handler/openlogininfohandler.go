package handler

import (
	"net/http"

	"IM_Server/im_auth/auth_api/internal/logic"
	"IM_Server/im_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
