package userpostgres

import (
	"context"
	"server/common"

	"github.com/google/uuid"
	usrmodel "server/modules/user/model"
)

func (repo *userRepo) GetProfile(ctx context.Context, userId uuid.UUID) (*usrmodel.User, error) {
	var user usrmodel.User
	err := repo.db.Model(&user).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(usrmodel.UserEntityName, err)
	}

	return &user, nil
}
