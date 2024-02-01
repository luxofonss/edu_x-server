package coursebiz

import (
	"context"
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

type CourseGetAllMyEnrollmentsRepo interface {
	GetAllMyEnrollments(userId uuid.UUID) ([]*coursemodel.UserEnrollCourse, error)
}

type getAllMyEnrollmentsCourseBiz struct {
	courseGetAllMyEnrollmentsRepo CourseGetAllMyEnrollmentsRepo
}

func NewGetAllMyEnrollmentsCourseBiz(courseGetAllMyEnrollmentsRepo CourseGetAllMyEnrollmentsRepo) *getAllMyEnrollmentsCourseBiz {
	return &getAllMyEnrollmentsCourseBiz{courseGetAllMyEnrollmentsRepo: courseGetAllMyEnrollmentsRepo}
}

func (biz *getAllMyEnrollmentsCourseBiz) GetAllMyEnrollments(ctx context.Context, userId uuid.UUID) ([]coursemodel.Course, error) {
	enrollments, err := biz.courseGetAllMyEnrollmentsRepo.GetAllMyEnrollments(userId)
	if err != nil {
		return nil, err
	}

	var courses []coursemodel.Course
	for _, enrollment := range enrollments {

		if enrollment.Status == coursemodel.ACTIVE {
			courses = append(courses, *enrollment.Course)
		}
	}

	return courses, nil
}
