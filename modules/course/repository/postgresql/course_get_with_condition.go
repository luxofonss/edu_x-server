package coursepg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetCourseWithCondition(
	ctx context.Context,
	filter *coursemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]*coursemodel.CourseGet, error) {
	db := repo.db.Table(coursemodel.Course{}.TableName())

	if filter != nil {
		if filter.TeacherId != uuid.Nil {
			db = db.Where("teacher_id = ?", filter.TeacherId)
		}

		if filter.Id != uuid.Nil {
			db = db.Where("id = ?", filter.Id)
		}

		fmt.Println(filter)

		if filter.Code != nil {

			db = db.Where("code = ?", *filter.Code)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	fmt.Println("here")

	var courses []*coursemodel.CourseGet
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&courses).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return courses, nil
}
