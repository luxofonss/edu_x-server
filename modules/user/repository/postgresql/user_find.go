package userpostgres

import (
	"context"

	"gorm.io/gorm"
	"server/common"
	usermodel "server/modules/user/model"
)

func (repo *userRepo) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := repo.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
