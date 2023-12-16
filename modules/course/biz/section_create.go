package coursebiz

import (
	"context"

	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

type CourseSectionRepo interface {
	CreateSection(ctx context.Context, data *coursemodel.Section, userId uuid.UUID) (*coursemodel.Section, error)
	FindCourseByOwnerId(ctx context.Context, teacherId uuid.UUID, courseId uuid.UUID) (*coursemodel.Course, error)
}

type createSectionBiz struct {
	courseRepo CourseSectionRepo
}

func NewCreateSectionBiz(courseRepo CourseSectionRepo) *createSectionBiz {
	return &createSectionBiz{courseRepo: courseRepo}
}

func (biz *createSectionBiz) CreateSection(ctx context.Context, data *coursemodel.Section, courseId uuid.UUID, userId uuid.UUID) (*coursemodel.Section, error) {
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
