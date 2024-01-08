package coursepg

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateCourseEnroll(ctx context.Context, courseId uuid.UUID, price float64, userId uuid.UUID, studentId int) (bool, error) {
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())

	data := coursemodel.UserEnrollCourse{
		UserId:    userId,
		CourseId:  courseId,
		Price:     price,
		StudentId: &studentId,
	}

	if err := db.Create(&data).Error; err != nil {
		return false, err
	}
	return true, nil
}

var (
	ErrCourseExpired = common.ErrCannotCreateEntity("Course", errors.New("Token expired!"))
)
