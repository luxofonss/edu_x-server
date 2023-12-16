package assignmentbiz

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type SubmitQuestionAnswerRepo interface {
	GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (*assignmentmodel.AssignmentAttempt, error)
	GetAssignmentByAssignmentPlacementId(
		ctx context.Context,
		assignmentAttemptId uuid.UUID,
	) (*assignmentmodel.Assignment, error)
	GetQuestionById(ctx context.Context, questionId uuid.UUID) (*assignmentmodel.Question, error)
	SubmitQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (*assignmentmodel.QuestionAnswer, error)
}

type submitQuestionAnswerBiz struct {
	repo SubmitQuestionAnswerRepo
}

func NewSubmitQuestionAnswerBiz(repo SubmitQuestionAnswerRepo) *submitQuestionAnswerBiz {
	return &submitQuestionAnswerBiz{repo: repo}
}

func (biz *submitQuestionAnswerBiz) SubmitQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (*assignmentmodel.QuestionAnswer, error) {
	assignmentAttempt, err := biz.repo.GetAssigmentAttemptById(ctx, data.AssignmentAttemptId)
	if err != nil {
		return nil, err
	}

	// TODO: check valid time to submit question

	var assignmentId uuid.UUID

	if assignmentAttempt.AssignmentId != uuid.Nil {
		assignmentId = assignmentAttempt.AssignmentId
	} else {
		assignment, err := biz.repo.GetAssignmentByAssignmentPlacementId(ctx, assignmentAttempt.AssignmentPlacementId)
		if err != nil {
			return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentEntityName, err)
		}

		assignmentId = assignment.Id
	}

	question, err := biz.repo.GetQuestionById(ctx, data.QuestionId)
	if err != nil {
		return nil, err
	}

	if question.AssignmentId != assignmentId {
		return nil, assignmentmodel.ErrQuestionNotInAssignment
	}

	// add answer
	questionAnswer, err := biz.repo.SubmitQuestionAnswer(ctx, data)
	if err != nil {
		return nil, err
	}

	return questionAnswer, nil
}
