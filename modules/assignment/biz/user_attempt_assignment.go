package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AttemptAssignmentRepo interface {
	GetAssignmentById(ctx context.Context, id uuid.UUID) (*assignmentmodel.Assignment, error)
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
	var canMultipleAttempt bool = true

	assignment, err := biz.attemptAssignmentRepo.GetAssignmentById(ctx, data.AssignmentId)
	if assignment.MultipleAttempt == false {
		canMultipleAttempt = false
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

	data.AssignmentTimeMillis = int64(assignment.Time) * 60 * 1000 // convert minute to millisecond (1 minute = 60 second = 60 * 1000 millisecond

	assignmentAttempt, err := biz.attemptAssignmentRepo.CreateAssignmentAttempt(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttempt, nil
}
