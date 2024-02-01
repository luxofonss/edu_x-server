package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentAttemptGetRepo interface {
	GetAssignmentAttemptById(ctx context.Context, id uuid.UUID) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptGetBiz struct {
	assignmentAttemptGetRepo AssignmentAttemptGetRepo
}

func NewAssignmentAttemptGetBiz(assignmentAttemptGetRepo AssignmentAttemptGetRepo) *assignmentAttemptGetBiz {
	return &assignmentAttemptGetBiz{assignmentAttemptGetRepo: assignmentAttemptGetRepo}
}

func (biz *assignmentAttemptGetBiz) GetAssignmentAttempt(ctx context.Context, id uuid.UUID) (*assignmentmodel.AssignmentAttempt, error) {
	assignmentAttempt, err := biz.assignmentAttemptGetRepo.GetAssignmentAttemptById(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttempt{}.TableName(), err)
	}
	return assignmentAttempt, nil
}
