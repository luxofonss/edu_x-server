package coursebiz

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

type CourseAttemptByCodeRepo interface {
	CreateCourseEnroll(ctx context.Context, courseId uuid.UUID, price float64, userId uuid.UUID, studentId int) (bool, error)
	GetCourseWithCondition(
		ctx context.Context,
		filter *coursemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*coursemodel.Course, error)
}

type courseAttemptByCodeBiz struct {
	repo CourseAttemptByCodeRepo
}

func NewCourseAttemptByCodeBiz(repo CourseAttemptByCodeRepo) *courseAttemptByCodeBiz {
	return &courseAttemptByCodeBiz{repo: repo}
}

func (biz *courseAttemptByCodeBiz) AttemptCourseByCode(ctx context.Context, code string, userId uuid.UUID) (bool, error) {
	paging := common.Paging{}
	paging.Fulfill()

	course, err := biz.repo.GetCourseWithCondition(ctx, &coursemodel.Filter{Code: &code}, &paging, "UserEnrollments")
	if len(course) == 0 {
		return false, errors.New("Course not found!")
	}

	fmt.Println(course[0].UserEnrollments)
	if &course[0].EndDate != nil {
		if common.CompareTimeNow(*course[0].EndDate) {
			return false, errors.New("Course expired!")
		}
	}

	res, err := biz.repo.CreateCourseEnroll(ctx, course[0].Id, course[0].Price, userId, len(course[0].UserEnrollments)+1)
	if err != nil {
		return false, err
	}

	return res, nil
}
