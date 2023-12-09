package userpostgres

import (
	"context"

	"server/common"
	usrmodel "server/modules/user/model"
)

func (repo *userRepo) CreateTeacherInfo(ctx context.Context, data *usrmodel.TeacherInfo) error {
	db := repo.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
