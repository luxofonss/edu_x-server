package assignmentbiz

import (
	"context"
	"fmt"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AttemptAssignmentRepo interface {
	CheckMultipleAttempt(ctx context.Context, assignmentPlacementId int) (bool, error)
	GetAllAttemptInAssignment(ctx context.Context, assignmentPlacementId int, userId int) ([]assignmentmodel.AssignmentAttempt, error)
	CreateAssignmentAttempt(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptBiz struct {
	attemptAssignmentRepo AttemptAssignmentRepo
}

func NewAttemptAssignmentBiz(attemptAssignmentRepo AttemptAssignmentRepo) *assignmentAttemptBiz {
	return &assignmentAttemptBiz{attemptAssignmentRepo: attemptAssignmentRepo}
}

func (biz *assignmentAttemptBiz) AttemptAssignment(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error) {
	canMultipleAttempt, err := biz.attemptAssignmentRepo.CheckMultipleAttempt(ctx, *data.AssignmentPlacementId)
	if err != nil {
		return nil, err
	}

	fmt.Println("canMultipleAttempt:: ", canMultipleAttempt)
	if !canMultipleAttempt {
		allAttempts, err := biz.attemptAssignmentRepo.GetAllAttemptInAssignment(ctx, *data.AssignmentPlacementId, data.UserId)
		if err != nil {
			return nil, err
		}

		fmt.Println("allAttempts:: ", allAttempts)
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
