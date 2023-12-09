package coursebiz

import (
	"context"

	coursemodel "server/modules/course/model"
)

type CourseSectionRepo interface {
	CreateSection(ctx context.Context, data *coursemodel.Section, userId int) (*coursemodel.Section, error)
	FindCourseByOwnerId(ctx context.Context, teacherId int, courseId int) (*coursemodel.Course, error)
}

type createSectionBiz struct {
	courseRepo CourseSectionRepo
}

func NewCreateSectionBiz(courseRepo CourseSectionRepo) *createSectionBiz {
	return &createSectionBiz{courseRepo: courseRepo}
}

func (biz *createSectionBiz) CreateSection(ctx context.Context, data *coursemodel.Section, courseId int, userId int) (*coursemodel.Section, error) {
	course, err := biz.courseRepo.FindCourseByOwnerId(ctx, userId, courseId)
	if err != nil {
		return nil, err
	}

	if course == nil {
		return nil, coursemodel.ErrCourseNotFound
	}

	data.CourseId = courseId
	section, err := biz.courseRepo.CreateSection(ctx, data, userId)
	if err != nil {
		return nil, err
	}

	return section, nil
}
