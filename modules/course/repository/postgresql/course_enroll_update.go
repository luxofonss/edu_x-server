package coursepg

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) UpdateCourseEnrollStatus(courseId, userId uuid.UUID, status string) error {
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())

	if err := db.Where("course_id = ? AND user_id = ?", courseId, userId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
