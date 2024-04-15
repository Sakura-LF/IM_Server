package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int32       `json:"code"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		body := Body{
			Code:    500,
			Message: err,
			Data:    nil,
		}
		httpx.WriteJson(w, http.StatusOK, body)
	}
	body := Body{
		Code:    200,
		Message: err,
		Data:    res,
	}
	httpx.WriteJson(w, http.StatusOK, body)

}
