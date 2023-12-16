package usrbiz

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	usrmodel "server/modules/user/model"
)

type ProfileRepo interface {
	GetProfile(ctx context.Context, userId uuid.UUID) (*usrmodel.User, error)
}

type ProfileBiz struct {
	repo ProfileRepo
}

func NewProfileBiz(repo ProfileRepo) *ProfileBiz {
	return &ProfileBiz{repo: repo}
}

func (biz *ProfileBiz) GetProfile(ctx context.Context, userId uuid.UUID) (*usrmodel.User, error) {
	profile, err := biz.repo.GetProfile(ctx, userId)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usrmodel.UserEntityName, err)
	}

	return profile, nil
}
