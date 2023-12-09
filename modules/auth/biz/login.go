package authbiz

import (
	"context"

	tokenprovider "server/libs/token-provider"
	authmodel "server/modules/auth/model"
	usrmodel "server/modules/user/model"
)

type Hasher interface {
	Hash(data string) string
}
type LoginRepo interface {
	GetAuth(ctx context.Context, data *authmodel.AuthLogin) (*authmodel.Auth, error)
}

type LoginBiz struct {
	loginRepo     LoginRepo
	hasher        Hasher
	tokenProvider tokenprovider.Provider
	expiry        int
}

func NewLoginBiz(loginRepo LoginRepo, hasher Hasher, tokenProvider tokenprovider.Provider, expiry int) *LoginBiz {
	return &LoginBiz{loginRepo: loginRepo, hasher: hasher, tokenProvider: tokenProvider, expiry: expiry}
}

func (biz *LoginBiz) Login(ctx context.Context, requestData *authmodel.AuthLogin) (*tokenprovider.Token, error) {
	auth, err := biz.loginRepo.GetAuth(ctx, requestData)
	if err != nil {
		return nil, usrmodel.ErrEmailOrPasswordInvalid
	}
	hashedPwd := biz.hasher.Hash(requestData.Password + auth.Salt)

	if auth.Password != hashedPwd {
		return nil, usrmodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: auth.UserId,
		Email:  auth.Email,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
