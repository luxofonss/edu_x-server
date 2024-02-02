package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	"time"
)

type AssignmentAttemptSubmitRepo interface {
	UpdateAssignmentAttempt(
		ctx context.Context,
		data *assignmentmodel.AssignmentAttempt,
	) (*assignmentmodel.AssignmentAttempt, error)
	GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptSubmitBiz struct {
	repo AssignmentAttemptSubmitRepo
}

func NewAssignmentAttemptSubmitBiz(repo AssignmentAttemptSubmitRepo) *assignmentAttemptSubmitBiz {
	return &assignmentAttemptSubmitBiz{repo: repo}
}

func (biz *assignmentAttemptSubmitBiz) UpdateAssignmentAttempt(
	ctx context.Context,
	assignmentAttemptId uuid.UUID,
	requesterId uuid.UUID,
) (*assignmentmodel.AssignmentAttempt, error) {

	data, err := biz.repo.GetAssigmentAttemptById(ctx, assignmentAttemptId)
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	if data.UserId != requesterId {
		return nil, common.NewCustomError(nil, "You do not have permission!", "assignment_attempt_not_belong_to_user")
	}

	if data.FinishedAt != nil {
		return nil, common.NewCustomError(nil, "Assignment has already finished", "assignment_attempt_already_finished")
	}

	// Check VALID TIME to submit question
	if data.AssignmentTimeMillis != 0 {
		assignmentCreatedAt, err := time.Parse(common.DateString, data.CreatedAt.String())
		assignmentTimeMillis := data.AssignmentTimeMillis
		maxSubmitTime := assignmentCreatedAt.Add(time.Duration(assignmentTimeMillis) * time.Millisecond)

		if time.Now().After(maxSubmitTime) {
			return nil, common.NewCustomError(err, "Time to submit this assignment has expired!", "assignment_attempt_already_finished")
		}
	}

	now := time.Time(time.Now())
	data.FinishedAt = &now
	return biz.repo.UpdateAssignmentAttempt(ctx, data)
}
