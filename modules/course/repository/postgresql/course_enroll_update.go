package coursepg

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) UpdateCourseEnrollStatus(courseEnrollId uuid.UUID, status string) error {
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())

	if err := db.Where("id = ?", courseEnrollId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
