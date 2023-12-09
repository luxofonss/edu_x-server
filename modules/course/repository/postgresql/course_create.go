package coursepg

import (
	"context"

	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateCourse(ctx context.Context, data *coursemodel.Course) (*coursemodel.Course, error) {
	db := repo.db.Table(coursemodel.Course{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
