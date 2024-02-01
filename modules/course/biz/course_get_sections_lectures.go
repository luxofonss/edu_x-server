package coursebiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

type CourseGetSectionLectureRepo interface {
	GetCourseWithCondition(
		ctx context.Context,
		filter *coursemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*coursemodel.Course, error)
}

type getSectionLectureBiz struct {
	repo CourseGetSectionLectureRepo
}

func NewGetSectionLectureBiz(repo CourseGetSectionLectureRepo) *getSectionLectureBiz {
	return &getSectionLectureBiz{repo: repo}
}

func (biz *getSectionLectureBiz) GetSectionLecture(ctx context.Context, courseId uuid.UUID) ([]*coursemodel.Section, error) {
	paging := common.Paging{}
	paging.Fulfill()
	courses, err := biz.repo.GetCourseWithCondition(ctx, &coursemodel.Filter{Id: courseId}, &paging, "Sections", "Sections.Lectures")
	if err != nil {
		return nil, err
	}

	if len(courses) == 0 {
		return nil, nil
	}

	return courses[0].Sections, nil
}
