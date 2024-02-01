package coursebiz

import (
	"github.com/google/uuid"
)

type CourseEnrollStatusUpdateRepo interface {
	UpdateCourseEnrollStatus(courseEnrollId uuid.UUID, status string) error
}

type updateCourseEnrollStatusBiz struct {
	courseEnrollStatusUpdateRepo CourseEnrollStatusUpdateRepo
}

func NewUpdateCourseEnrollStatusBiz(courseEnrollStatusUpdateRepo CourseEnrollStatusUpdateRepo) *updateCourseEnrollStatusBiz {
	return &updateCourseEnrollStatusBiz{courseEnrollStatusUpdateRepo: courseEnrollStatusUpdateRepo}
}

func (biz *updateCourseEnrollStatusBiz) UpdateCourseEnrollStatus(courseEnrollId uuid.UUID, status string) error {
	err := biz.courseEnrollStatusUpdateRepo.UpdateCourseEnrollStatus(courseEnrollId, status)
	if err != nil {
		return err
	}

	return nil
}
