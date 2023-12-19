package coursepg

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetCourseWithCondition(
	ctx context.Context,
	filter *coursemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]coursemodel.CourseGet, error) {
	db := repo.db.Table(coursemodel.Course{}.TableName())

	if v := filter; v != nil {
		if v.TeacherId != uuid.Nil {
			db = db.Where("teacher_id = ?", v.TeacherId)
		}

		if v.Id != uuid.Nil {
			db = db.Where("id = ?", v.Id)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var courses []coursemodel.CourseGet
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&courses).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return courses, nil
}
