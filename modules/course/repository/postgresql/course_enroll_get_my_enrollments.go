package coursepg

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetAllMyEnrollments(userId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error) {
	var result []*coursemodel.UserEnrollCourse
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())

	db = db.Preload("Course").Preload("Course.Teacher").Preload("Course.Subject")
	if err := db.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
