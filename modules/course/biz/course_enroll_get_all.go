package coursebiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

type CourseGetAllEnrollmentsRepo interface {
	GetAllCourseEnrollmentsByCourseId(courseId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error)
}

type getAllEnrollmentsCourseBiz struct {
	courseGetAllEnrollmentsRepo CourseGetAllEnrollmentsRepo
}

func NewGetAllEnrollmentsCourseBiz(courseGetAllEnrollmentsRepo CourseGetAllEnrollmentsRepo) *getAllEnrollmentsCourseBiz {
	return &getAllEnrollmentsCourseBiz{courseGetAllEnrollmentsRepo: courseGetAllEnrollmentsRepo}
}

func (biz *getAllEnrollmentsCourseBiz) GetAllEnrollmentsCourse(ctx context.Context, courseId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error) {
	paging := common.Paging{}
	paging.Fulfill()
	course, err := biz.courseGetAllEnrollmentsRepo.GetAllCourseEnrollmentsByCourseId(
		courseId,
	)
	if err != nil {
		return nil, err
	}

	return course, nil
}
