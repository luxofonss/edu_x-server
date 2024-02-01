package coursebiz

import (
	"context"
	"server/common"
	coursemodel "server/modules/course/model"
)

type GetCourseRepo interface {
	GetCourseWithCondition(
		ctx context.Context,
		filter *coursemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*coursemodel.Course, error)
}

type getCourseBiz struct {
	courseRepo GetCourseRepo
}

func NewGetCourseBiz(courseRepo GetCourseRepo) *getCourseBiz {
	return &getCourseBiz{courseRepo: courseRepo}
}

func (biz *getCourseBiz) GetAllMyCourses(
	ctx context.Context,
	paging *common.Paging,
	filter *coursemodel.Filter,
) ([]*coursemodel.Course, error) {
	courses, err := biz.courseRepo.GetCourseWithCondition(ctx, filter, paging, "Subject", "Teacher")
	if err != nil {
		return nil, common.ErrCannotListEntity(coursemodel.CourseEntityName, err)
	}

	return courses, nil
}
