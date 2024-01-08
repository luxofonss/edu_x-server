package assignmentbiz

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentGetRepo interface {
	GetAssignmentById(ctx context.Context, id uuid.UUID) (*assignmentmodel.Assignment, error)
}

type assignmentGetBiz struct {
	assignmentGetRepo AssignmentGetRepo
}

func NewAssignmentGetBiz(assignmentGetRepo AssignmentGetRepo) *assignmentGetBiz {
	return &assignmentGetBiz{assignmentGetRepo: assignmentGetRepo}
}

func (biz *assignmentGetBiz) GetAssignmentById(ctx context.Context, id uuid.UUID) (*assignmentmodel.Assignment, error) {
	assignment, err := biz.assignmentGetRepo.GetAssignmentById(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.Assignment{}.TableName(), err)
	}
	return assignment, nil
}
