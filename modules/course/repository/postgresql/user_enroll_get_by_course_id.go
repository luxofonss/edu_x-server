package coursepg

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetAllCourseEnrollmentsByCourseId(courseId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error) {
	var result []*coursemodel.UserEnrollCourse
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())
	db.Preload("User")
	if err := db.Where("course_id = ?", courseId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
