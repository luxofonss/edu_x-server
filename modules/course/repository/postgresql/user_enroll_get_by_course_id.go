package coursepg

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetAllCourseEnrollmentsByCourseId(courseId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error) {
	var result []*coursemodel.UserEnrollCourse
	if err := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName()).Where("course_id = ?", courseId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
