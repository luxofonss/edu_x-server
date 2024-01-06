package coursebiz

import (
	"context"
	"server/common"
	coursemodel "server/modules/course/model"

	"github.com/google/uuid"
)

type CourseGetOneRepo interface {
	GetCourseWithCondition(
		ctx context.Context,
		filter *coursemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*coursemodel.CourseGet, error)
}

type getOneCourseByIdBiz struct {
	courseGetOneRepo CourseGetOneRepo
}

func NewGetOneCourseBiz(courseGetOneRepo CourseGetOneRepo) *getOneCourseByIdBiz {
	return &getOneCourseByIdBiz{courseGetOneRepo: courseGetOneRepo}
}

func (biz *getOneCourseByIdBiz) GetOneCourseById(ctx context.Context, courseId uuid.UUID) (*coursemodel.CourseGet, error) {
	paging := common.Paging{}
	paging.Fulfill()
	course, err := biz.courseGetOneRepo.GetCourseWithCondition(ctx, &coursemodel.Filter{Id: courseId}, &paging)
	if err != nil {
		return nil, err
	}

	if len(course) == 0 {
		return nil, common.ErrEntityNotFound(coursemodel.CourseEntityName, nil)
	}

	return course[0], nil
}
