package coursebiz

import (
	"github.com/google/uuid"
)

type CourseEnrollStatusUpdateRepo interface {
	UpdateCourseEnrollStatus(courseId, userId uuid.UUID, status string) error
}

type updateCourseEnrollStatusBiz struct {
	courseEnrollStatusUpdateRepo CourseEnrollStatusUpdateRepo
}

func NewUpdateCourseEnrollStatusBiz(courseEnrollStatusUpdateRepo CourseEnrollStatusUpdateRepo) *updateCourseEnrollStatusBiz {
	return &updateCourseEnrollStatusBiz{courseEnrollStatusUpdateRepo: courseEnrollStatusUpdateRepo}
}

func (biz *updateCourseEnrollStatusBiz) UpdateCourseEnrollStatus(courseId, userId uuid.UUID, status string) error {
	err := biz.courseEnrollStatusUpdateRepo.UpdateCourseEnrollStatus(courseId, userId, status)
	if err != nil {
		return err
	}

	return nil
}
