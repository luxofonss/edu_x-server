package userpostgres

import (
	"context"

	"server/common"
	usrmodel "server/modules/user/model"
)

func (repo *userRepo) CreateLearnerInfo(ctx context.Context, data *usrmodel.LearnerInfo) error {
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
