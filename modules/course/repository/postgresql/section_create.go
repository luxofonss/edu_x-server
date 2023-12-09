package coursepg

import (
	"context"

	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateSection(ctx context.Context, data *coursemodel.Section, userId int) (*coursemodel.Section, error) {
	db := repo.db.Table(coursemodel.Section{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
