package assignmentbiz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	"time"
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
	assignment, err := biz.attemptAssignmentRepo.GetAssignmentById(ctx, data.AssignmentId)

	allAttempts, err := biz.attemptAssignmentRepo.GetAllAttemptInAssignment(ctx, data.AssignmentId, data.UserId)
	if err != nil {
		return nil, err
	}

	fmt.Println("allAttempts:: ", allAttempts)

	if len(allAttempts) > 0 {
		lastAttempt := allAttempts[len(allAttempts)-1]
		assignmentTimeMillis := lastAttempt.AssignmentTimeMillis

		fmt.Println("assignmentTimeMillis:: ", assignmentTimeMillis)

		if assignmentTimeMillis != 0 {
			assignmentCreatedAt, err := time.Parse(common.DateString, lastAttempt.CreatedAt.String())
			if err != nil {
				return nil, err
			}

			maxSubmitTime := assignmentCreatedAt.Add(time.Duration(assignmentTimeMillis) * time.Millisecond)

			timeNow := time.Now().Add(time.Hour * 7) // GTM +7
			if timeNow.Before(maxSubmitTime) {
				return &lastAttempt, nil
			}

		}

		if !assignment.MultipleAttempt {
			return nil, common.NewCustomError(err, "You can attempt this assignment one time!", "CREATE_ASSIGMENT")
		}
	}

	fmt.Println("add new attempt")

	data.AssignmentTimeMillis = int64(assignment.Time) * 60 * 1000 // convert minute to millisecond (1 minute = 60 second = 60 * 1000 millisecond

	assignmentAttempt, err := biz.attemptAssignmentRepo.CreateAssignmentAttempt(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttempt, nil
}
