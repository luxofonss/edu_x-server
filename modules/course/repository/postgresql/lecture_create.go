package coursepg

import (
	"context"

	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateLecture(ctx context.Context, data *coursemodel.Lecture) (*coursemodel.Lecture, error) {
	db := repo.db.Table(coursemodel.Lecture{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
