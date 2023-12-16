package coursebiz

import (
	"context"

	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentRepo interface {
	GetAssignmentByCourseId(ctx context.Context, id uuid.UUID) ([]*assignmentmodel.Assignment, error)
}

type getAssignmentsInCourseBiz struct {
	assignmentRepo AssignmentRepo
}

func NewGetAssignmentsInCourseBiz(assignmentRepo AssignmentRepo) *getAssignmentsInCourseBiz {
	return &getAssignmentsInCourseBiz{assignmentRepo: assignmentRepo}
}

func (biz *getAssignmentsInCourseBiz) GetAssignmentsInCourse(ctx context.Context, courseId uuid.UUID) ([]*assignmentmodel.Assignment, error) {
	assignments, err := biz.assignmentRepo.GetAssignmentByCourseId(ctx, courseId)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
