package assignmentbiz

import (
	"context"
	"server/common"
	"time"

	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type SubmitQuestionAnswerRepo interface {
	GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (
		*assignmentmodel.AssignmentAttempt,
		error,
	)
	GetQuestionById(ctx context.Context, questionId uuid.UUID, moreKeys ...string) (
		*assignmentmodel.Question,
		error,
	)
	SubmitQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (
		*assignmentmodel.QuestionAnswer,
		error,
	)
	GetAnswer(ctx context.Context, questionId uuid.UUID, userId uuid.UUID) (
		*assignmentmodel.QuestionAnswer,
		error,
	)
}

type submitQuestionAnswerBiz struct {
	repo SubmitQuestionAnswerRepo
}

func NewSubmitQuestionAnswerBiz(repo SubmitQuestionAnswerRepo) *submitQuestionAnswerBiz {
	return &submitQuestionAnswerBiz{repo: repo}
}

func (biz *submitQuestionAnswerBiz) SubmitQuestionAnswer(
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

	// Check VALID TIME to submit question
	if assignmentAttempt.AssignmentTimeMillis != 0 {
		assignmentCreatedAt, err := time.Parse(common.DateString, assignmentAttempt.CreatedAt.String())

		assignmentTimeMillis := assignmentAttempt.AssignmentTimeMillis
		maxSubmitTime := assignmentCreatedAt.Add(time.Duration(assignmentTimeMillis) * time.Millisecond)

		if time.Now().After(maxSubmitTime) {
			return nil, common.NewCustomError(err, "Time to submit this assignment has expired!", "SUBMIT_QUESTION")
		}
	}

	// Add answer
	questionAnswer, err := biz.repo.SubmitQuestionAnswer(ctx, data)
	if err != nil {
		return nil, err
	}

	return questionAnswer, nil
}
