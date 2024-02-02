package assignmentbiz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentAttemptLongAnswerScoreRepo interface {
	UpdateLongAnswerScore(
		ctx context.Context,
		questionAnswerId uuid.UUID,
		point int,
	) (*assignmentmodel.QuestionAnswer, error)
	GetAssignmentAttemptById(ctx context.Context, id uuid.UUID) (*assignmentmodel.AssignmentAttempt, error)
	UpdateAssignmentAttempt(
		ctx context.Context,
		data *assignmentmodel.AssignmentAttempt,
	) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptLongAnswerScoreBiz struct {
	assignmentAttemptLongAnswerScoreRepo AssignmentAttemptLongAnswerScoreRepo
}

func NewAssignmentAttemptLongAnswerScoreBiz(
	assignmentAttemptLongAnswerScoreRepo AssignmentAttemptLongAnswerScoreRepo,
) *assignmentAttemptLongAnswerScoreBiz {
	return &assignmentAttemptLongAnswerScoreBiz{assignmentAttemptLongAnswerScoreRepo: assignmentAttemptLongAnswerScoreRepo}
}

func (biz *assignmentAttemptLongAnswerScoreBiz) ScoreAssignmentAttemptLongAnswer(
	ctx context.Context,
	assignmentAttemptId uuid.UUID,
	questionAnswerId uuid.UUID,
	point int,
) error {

	questionAnswer, err := biz.assignmentAttemptLongAnswerScoreRepo.UpdateLongAnswerScore(ctx, questionAnswerId, point)
	if err != nil {
		return err
	}

	assignmentAttempt, err := biz.assignmentAttemptLongAnswerScoreRepo.GetAssignmentAttemptById(ctx, questionAnswer.AssignmentAttemptId)
	if err != nil {
		return err
	}

	if assignmentAttempt.Id != assignmentAttemptId {
		return common.NewCustomError(nil, "Assignment attempt not found", "SCORE_LONG_ANSWER")
	}

	newAssignmentAttempt := assignmentAttempt
	newAssignmentAttempt.Point = new(int)
	for _, question := range assignmentAttempt.Assignment.Questions {
		if question.Type == assignmentmodel.SingleChoice {
			if len(question.Answers) > 0 && question.Answers[0].SelectedOptionId != nil {
				for _, choice := range question.Choices {
					if choice.Id == *question.Answers[0].SelectedOptionId && choice.IsCorrect {
						fmt.Println("question.Point:: ", *question.Point)
						*newAssignmentAttempt.Point += *question.Point
					}
				}
			}
		} else if question.Type == assignmentmodel.MultipleChoice {
			// TODO: handle multi choice
		} else if question.Type == assignmentmodel.ShortAnswer {
			for _, answer := range question.CorrectAnswers {
				if answer.CorrectTextAnswer == question.Answers[0].TextAnswer {
					*newAssignmentAttempt.Point += *question.Point
				}
			}
		} else if question.Type == assignmentmodel.LongAnswer {
			if question.Answers[0].Score != nil {
				*newAssignmentAttempt.Point += *question.Answers[0].Score
			}
		}
	}

	_, err = biz.assignmentAttemptLongAnswerScoreRepo.UpdateAssignmentAttempt(ctx, newAssignmentAttempt)
	if err != nil {
		return err
	}

	return nil
}
