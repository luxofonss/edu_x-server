package userpostgres

import (
	"context"
	"server/common"
	usermodel "server/modules/user/model"
)

func (repo *userRepo) GetUsersByEmails(ctx context.Context, emails []string) ([]*usermodel.User, error) {
	db := repo.db.Table(usermodel.User{}.TableName())

	var users []*usermodel.User

	err := db.Where("email IN (?)", emails).Find(&users).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.UserEntityName, err)
	}

	return users, nil
}
