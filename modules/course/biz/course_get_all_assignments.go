package coursebiz

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

type AssignmentRepo interface {
	GetAssignmentByCourseId(ctx context.Context, id int) (*assignmentmodel.Assignment, error)
}

type getAssignmentsInCourseBiz struct {
	assignmentRepo AssignmentRepo
}

func NewGetAssignmentsInCourseBiz(assignmentRepo AssignmentRepo) *getAssignmentsInCourseBiz {
	return &getAssignmentsInCourseBiz{assignmentRepo: assignmentRepo}
}

func (biz *getAssignmentsInCourseBiz) GetAssignmentsInCourse(ctx context.Context, courseId int) ([]*assignmentmodel.Assignment, error) {
	//assignments, err := biz.assignmentRepo.GetAssignment(ctx, courseId)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
