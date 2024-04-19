package logic

import (
	"IM_Server/utils/jwt"
	"context"
	"errors"
	"fmt"
	"time"

	"IM_Server/im_auth/auth_api/internal/svc"
	"IM_Server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp *types.Response, err error) {
	if token == "" {
		err = errors.New("请传入token")
		return &types.Response{
			Code: 400,
			Data: nil,
			Msg:  err.Error(),
		}, err
	}

	payLoad, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token错误")
		return &types.Response{
			Code: 400,
			Data: nil,
			Msg:  err.Error(),
		}, err
	}
	// 过期时间
	expiration := payLoad.ExpiresAt.Time.Sub(time.Now())
	key := fmt.Sprintf("logout_%d", payLoad.UserID)
	l.svcCtx.Redis.SetNX(context.Background(), key, "", expiration)

	return &types.Response{
		Code: 200,
		Msg:  "注销成功",
		Data: nil,
	}, nil
}
