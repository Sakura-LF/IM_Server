package logic

import (
	"context"

	"IM_Server/im_auth/auth_api/internal/svc"
	"IM_Server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	return &types.LoginResponse{
		Code: 200,
		Data: types.LoginInfo{
			Token: "XXX",
		},
		Msg: "登录成功",
	}, nil
}
