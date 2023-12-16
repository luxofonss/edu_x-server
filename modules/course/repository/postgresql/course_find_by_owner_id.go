package coursepg

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) FindCourseByOwnerId(ctx context.Context, teacherId uuid.UUID, courseId uuid.UUID) (*coursemodel.Course, error) {
	db := repo.db.Table(coursemodel.Course{}.TableName())

	var courses *coursemodel.Course
	if err := db.Where("teacher_id = ? AND id = ?", teacherId, courseId).First(&courses).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return courses, nil
}
