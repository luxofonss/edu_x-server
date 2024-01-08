package coursebiz

import (
	"context"
	"server/common"
	coursemodel "server/modules/course/model"
)

type CourseGetAllActiveRepo interface {
	GetCourseWithCondition(
		ctx context.Context,
		filter *coursemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*coursemodel.CourseGet, error)
}

type getAllActiveCourseBiz struct {
	courseGetAllActiveRepo CourseGetAllActiveRepo
}

func NewGetAllActiveCourseBiz(courseGetAllActiveRepo CourseGetAllActiveRepo) *getAllActiveCourseBiz {
	return &getAllActiveCourseBiz{courseGetAllActiveRepo: courseGetAllActiveRepo}
}

func (biz *getAllActiveCourseBiz) GetAllActiveCourse(ctx context.Context) ([]*coursemodel.CourseGet, error) {
	paging := common.Paging{}
	paging.Fulfill()
	course, err := biz.courseGetAllActiveRepo.GetCourseWithCondition(
		ctx,
		&coursemodel.Filter{Status: "active"},
		&paging,
		"Teacher", "Subject",
	)
	if err != nil {
		return nil, err
	}

	return course, nil
}
