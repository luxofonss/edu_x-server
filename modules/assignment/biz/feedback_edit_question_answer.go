package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type FeedbackEditQuestionAnswerRepo interface {
	GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (
		*assignmentmodel.AssignmentAttempt,
		error,
	)
	GetQuestionById(ctx context.Context, questionId uuid.UUID, moreKeys ...string) (
		*assignmentmodel.Question,
		error,
	)
	UpdateQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (
		*assignmentmodel.QuestionAnswer,
		error,
	)
	GetAnswer(ctx context.Context, questionId uuid.UUID, userId uuid.UUID) (
		*assignmentmodel.QuestionAnswer,
		error,
	)
}

type feedbackEditQuestionAnswerBiz struct {
	repo FeedbackEditQuestionAnswerRepo
}

func NewFeedbackEditQuestionAnswer(repo FeedbackEditQuestionAnswerRepo) *feedbackEditQuestionAnswerBiz {
	return &feedbackEditQuestionAnswerBiz{repo: repo}
}

func (biz *feedbackEditQuestionAnswerBiz) FeedbackEditQuestionAnswer(
	ctx context.Context,
	data *assignmentmodel.QuestionAnswer,
) (*assignmentmodel.QuestionAnswer, error) {
	// Get assignment
	assignmentAttempt, err := biz.repo.GetAssigmentAttemptById(ctx, data.AssignmentAttemptId)
	if err != nil {
		return nil, err
	}

	//  Check if question is in assignment
	var assignmentId uuid.UUID
	assignmentId = assignmentAttempt.AssignmentId

	question, err := biz.repo.GetQuestionById(ctx, data.QuestionId, "Choices", "CorrectAnswers")
	if err != nil {
		return nil, err
	}

	if question.AssignmentId != assignmentId {
		return nil, assignmentmodel.ErrQuestionNotInAssignment
	}

	// Add answer
	questionAnswer, err := biz.repo.UpdateQuestionAnswer(ctx, data)
	if err != nil {
		return nil, err
	}

	return questionAnswer, nil
}
