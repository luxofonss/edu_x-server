package userpostgres

import (
	"context"
	"fmt"

	"server/common"
	usermodel "server/modules/user/model"
)

func (repo *userRepo) CreateUser(ctx context.Context, data *usermodel.User) (*usermodel.User, error) {
	db := repo.db.Table(usermodel.User{}.TableName())

	fmt.Print("data.Id:: ", data.Id)
	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
