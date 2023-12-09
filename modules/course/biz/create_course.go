package coursebiz

import (
	"context"

	"server/common"
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

func (biz *createCourseBiz) CreateCourse(ctx context.Context, data *coursemodel.CourseCreate, teacherId int) (*coursemodel.Course, error) {
	var course *coursemodel.Course
	course = &data.Course

	var courseInfos []*coursemodel.CourseInfo
	courseInfos = data.CourseInfos

	course.TeacherId = teacherId
	course.IsVerified = false
	course.SchoolId = nil
	course.Code = common.GenCourseCode(6)

	createdCourse, err := biz.courseRepo.CreateCourse(ctx, course)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(coursemodel.CourseEntityName, err)
	}

	for _, courseInfo := range courseInfos {
		courseInfo.CourseId = createdCourse.Id
	}

	if err := biz.courseRepo.CreateManyCourseInfo(ctx, courseInfos); err != nil {
		return nil, common.ErrCannotCreateEntity(coursemodel.CourseInfoEntityName, err)
	}

	return createdCourse, nil
}
