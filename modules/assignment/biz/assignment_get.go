package assignmentbiz

import (
	"context"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentGetRepo interface {
	GetAssignment(ctx context.Context, id int) (*assignmentmodel.Assignment, error)
}

type assignmentGetBiz struct {
	assignmentGetRepo AssignmentGetRepo
}

func NewAssignmentGetBiz(assignmentGetRepo AssignmentGetRepo) *assignmentGetBiz {
	return &assignmentGetBiz{assignmentGetRepo: assignmentGetRepo}
}

func (biz *assignmentGetBiz) GetAssignment(ctx context.Context, id int) (*assignmentmodel.Assignment, error) {
	assignment, err := biz.assignmentGetRepo.GetAssignment(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.Assignment{}.TableName(), err)
	}

	return assignment, nil
}
