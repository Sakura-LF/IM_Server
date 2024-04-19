package logic

import (
	"IM_Server/im_auth/auth_api/internal/svc"
	"IM_Server/im_auth/auth_api/internal/types"
	"IM_Server/utils/jwt"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp *types.Response, err error) {
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
	key := fmt.Sprintf("logout_%s", payLoad.UserID)
	result, err := l.svcCtx.Redis.Get(context.Background(), key).Result()
	fmt.Println("Result", result)
	fmt.Println("Err", err)
	if err == nil {
		return &types.Response{
			Code: 200,
			Data: nil,
			Msg:  "认证失败",
		}, err
	}

	return &types.Response{
		Code: 200,
		Data: nil,
		Msg:  "认证完成",
	}, nil
}
