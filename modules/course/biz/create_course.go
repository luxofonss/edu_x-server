package coursebiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	coursedto "server/modules/course/dto"
	coursemodel "server/modules/course/model"
)

type CourseRepo interface {
	CreateCourse(ctx context.Context, data *coursemodel.Course) (*coursemodel.Course, error)
	CreateManyCourseInfo(ctx context.Context, data []*coursemodel.CourseInfo) error
}

type createCourseBiz struct {
	courseRepo CourseRepo
}

func NewCreateCourseBiz(courseRepo CourseRepo) *createCourseBiz {
	return &createCourseBiz{courseRepo: courseRepo}
}

func (biz *createCourseBiz) CreateCourse(ctx context.Context, data *coursedto.CourseCreateRequest, teacherId uuid.UUID) (*coursemodel.Course, error) {
	//data.TeacherId = teacherId
	//data.IsVerified = false
	//data.Code = common.GenCourseCode(6)

	course := data.ToCourseModel()
	course.TeacherId = teacherId
	course.IsVerified = false
	course.Code = common.GenCourseCode(6)

	createdCourse, err := biz.courseRepo.CreateCourse(ctx, course)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(coursemodel.CourseEntityName, err)
	}

	return createdCourse, nil
}
