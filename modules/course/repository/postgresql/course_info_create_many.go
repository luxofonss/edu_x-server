package coursepg

import (
	"context"

	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateManyCourseInfo(ctx context.Context, data []*coursemodel.CourseInfo) error {
	db := repo.db.Table(coursemodel.CourseInfo{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
