package coursebiz

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

type CourseLectureRepo interface {
	CreateLecture(ctx context.Context, data *coursemodel.Lecture) (*coursemodel.Lecture, error)
	FindCourseByOwnerId(ctx context.Context, teacherId uuid.UUID, courseId uuid.UUID) (*coursemodel.Course, error)
	GetSectionCourseInfo(ctx context.Context, sectionId uuid.UUID) (*coursemodel.Section, error)
}

type createLectureBiz struct {
	courseRepo CourseLectureRepo
}

func NewCreateLectureBiz(courseRepo CourseLectureRepo) *createLectureBiz {
	return &createLectureBiz{courseRepo: courseRepo}
}

func (biz *createLectureBiz) CreateLecture(
	ctx context.Context,
	data *coursemodel.Lecture,
	courseId uuid.UUID,
	sectionId uuid.UUID,
	userId uuid.UUID,
) (*coursemodel.Lecture, error) {
	sectionCourse, err := biz.courseRepo.GetSectionCourseInfo(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	if sectionCourse == nil {
		return nil, coursemodel.ErrCourseNotFound
	}

	if sectionCourse.CourseId != courseId {
		return nil, coursemodel.ErrCourseNotFound
	}

	if sectionCourse.Course.TeacherId != userId {
		return nil, common.ErrNoPermission(errors.New("you don't have permission to create lecture"))
	}

	data.SectionId = sectionId
	lecture, err := biz.courseRepo.CreateLecture(ctx, data)
	if err != nil {
		return nil, err
	}

	return lecture, nil
}
