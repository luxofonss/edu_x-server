package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AttemptAssignmentRepo interface {
	CheckMultipleAttempt(ctx context.Context, assignmentId uuid.UUID) (bool, error)
	GetAllAttemptInAssignment(ctx context.Context, assignmentId uuid.UUID, userId uuid.UUID) ([]assignmentmodel.AssignmentAttempt, error)
	CreateAssignmentAttempt(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptBiz struct {
	attemptAssignmentRepo AttemptAssignmentRepo
}

func NewAttemptAssignmentBiz(attemptAssignmentRepo AttemptAssignmentRepo) *assignmentAttemptBiz {
	return &assignmentAttemptBiz{attemptAssignmentRepo: attemptAssignmentRepo}
}

func (biz *assignmentAttemptBiz) AttemptAssignment(
	ctx context.Context,
	data *assignmentmodel.AssignmentAttemptCreate,
) (*assignmentmodel.AssignmentAttempt, error) {
	canMultipleAttempt, err := biz.attemptAssignmentRepo.CheckMultipleAttempt(ctx, data.AssignmentId)
	if err != nil {
		return nil, err
	}

	if !canMultipleAttempt {
		allAttempts, err := biz.attemptAssignmentRepo.GetAllAttemptInAssignment(ctx, data.AssignmentId, data.UserId)
		if err != nil {
			return nil, err
		}

		if len(allAttempts) > 0 {
			return nil, common.NewCustomError(err, "You can attempt this assignment one time!", "CREATE_ASSIGMENT")
		}
	}

	assignmentAttempt, err := biz.attemptAssignmentRepo.CreateAssignmentAttempt(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttempt, nil
}
