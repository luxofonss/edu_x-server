package authbiz

import (
	"context"

	"server/common"
	authmodel "server/modules/auth/model"
	usrmodel "server/modules/user/model"
)

type UserRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usrmodel.User, error)
	CreateUser(ctx context.Context, data *usrmodel.User) (*usrmodel.User, error)
	CreateLearnerInfo(ctx context.Context, data *usrmodel.LearnerInfo) error
	CreateTeacherInfo(ctx context.Context, data *usrmodel.TeacherInfo) error
}
type registerBiz struct {
	userRepo UserRepo
	hasher   Hasher
}

func NewRegisterBiz(userRepo UserRepo, hasher Hasher) *registerBiz {
	return &registerBiz{userRepo: userRepo, hasher: hasher}
}

func (biz *registerBiz) Register(ctx context.Context, data *usrmodel.User) (*usrmodel.User, error) {
	user, _ := biz.userRepo.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return nil, usrmodel.ErrEmailExisted
	}

	// Create user auth
	authData := authmodel.Auth{
		Email:    data.Email,
		Password: data.Password,
		AuthType: "email_pwd",
	}

	salt := common.GetSalt(50)
	authData.Password = biz.hasher.Hash(authData.Password + salt)
	authData.Salt = salt
	data.Auth = &authData

	// Create user data
	createdUser, err := biz.userRepo.CreateUser(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(usrmodel.UserEntityName, err)
	}

	// Return result
	return createdUser, nil
}
