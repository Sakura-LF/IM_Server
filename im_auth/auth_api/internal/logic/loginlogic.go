package logic

import (
	"IM_Server/im_auth/auth_api/internal/svc"
	"IM_Server/im_auth/auth_api/internal/types"
	"IM_Server/im_auth/auth_models"
	"IM_Server/utils/jwt"
	"IM_Server/utils/pwd"
	"context"
	"errors"

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
	var user auth_models.User
	//
	if err = l.svcCtx.DB.Take(&user, "id = ? ", req.UserName).Error; err != nil {
		err = errors.New("用户名或密码错误")
		return
		//return &types.LoginResponse{
		//	Code: 400,
		//	Data: types.LoginInfo{},
		//	Msg:  "用户名或密码错误",
		//}, err
	}
	if !pwd.CheckPwd(user.Pwd, req.Password) {
		err = errors.New("用户名或密码错误")
		//return &types.LoginResponse{
		//	Code: 400,
		//	Data: types.LoginInfo{},
		//	Msg:  "用户名或密码错误",
		//}, err
		return
	}
	// 生成token
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		err = errors.New("服务内部错误")
		return
	}

	return &types.LoginResponse{
		Code: 200,
		Data: types.LoginInfo{
			Token: token,
		},
		Msg: "登录成功",
	}, nil
}
