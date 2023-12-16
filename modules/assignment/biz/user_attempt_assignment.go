package assignmentbiz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AttemptAssignmentRepo interface {
	CheckMultipleAttempt(ctx context.Context, assignmentPlacementId uuid.UUID) (bool, error)
	GetAllAttemptInAssignment(ctx context.Context, assignmentPlacementId uuid.UUID, userId uuid.UUID) ([]assignmentmodel.AssignmentAttempt, error)
	CreateAssignmentAttempt(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error)
	GetAssignmentByAssignmentPlacementId(ctx context.Context, AssignmentPlacementId uuid.UUID) (*assignmentmodel.Assignment, error)
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
	canMultipleAttempt, err := biz.attemptAssignmentRepo.CheckMultipleAttempt(ctx, data.AssignmentPlacementId)
	if err != nil {
		return nil, err
	}

	if !canMultipleAttempt {
		allAttempts, err := biz.attemptAssignmentRepo.GetAllAttemptInAssignment(ctx, data.AssignmentPlacementId, data.UserId)
		if err != nil {
			return nil, err
		}

		if len(allAttempts) > 0 {
			return nil, common.NewCustomError(err, "You can attempt this assignment one time!", "CREATE_ASSIGMENT")
		}
	}

	assignment, err := biz.attemptAssignmentRepo.GetAssignmentByAssignmentPlacementId(ctx, data.AssignmentPlacementId)
	if err != nil {
		return nil, err
	}
	data.AssignmentId = assignment.Id

	fmt.Println(assignment.Id, data.AssignmentPlacementId)

	assignmentAttempt, err := biz.attemptAssignmentRepo.CreateAssignmentAttempt(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttempt, nil
}
