package assignmentbiz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	"time"
)

type AssignmentAttemptGetResultRepo interface {
	GetAssignmentAttemptById(ctx context.Context, id uuid.UUID) (*assignmentmodel.AssignmentAttempt, error)
	UpdateAssignmentAttempt(
		ctx context.Context,
		data *assignmentmodel.AssignmentAttempt,
	) (*assignmentmodel.AssignmentAttempt, error)
}

type assignmentAttemptGetResultBiz struct {
	assignmentAttemptGetResultRepo AssignmentAttemptGetResultRepo
}

func NewAssignmentAttemptGetResultBiz(assignmentAttemptGetResultRepo AssignmentAttemptGetResultRepo) *assignmentAttemptGetResultBiz {
	return &assignmentAttemptGetResultBiz{assignmentAttemptGetResultRepo: assignmentAttemptGetResultRepo}
}

func (biz *assignmentAttemptGetResultBiz) GetAssignmentAttemptResult(ctx context.Context, id uuid.UUID) (*assignmentmodel.AssignmentAttempt, error) {
	assignmentAttempt, err := biz.assignmentAttemptGetResultRepo.GetAssignmentAttemptById(ctx, id)
	if err != nil {
		return nil, err
	}

	if assignmentAttempt.AssignmentTimeMillis != 0 && assignmentAttempt.FinishedAt == nil {
		assignmentCreatedAt, err := time.Parse(common.DateString, assignmentAttempt.CreatedAt.String())

		assignmentTimeMillis := assignmentAttempt.AssignmentTimeMillis
		maxSubmitTime := assignmentCreatedAt.Add(time.Duration(assignmentTimeMillis) * time.Millisecond)

		fmt.Println("maxSubmitTime:: ", maxSubmitTime, time.Now())
		if time.Now().Before(maxSubmitTime) {
			return nil, common.NewCustomError(err, "This attempt is in time!", "SUBMIT_QUESTION")
		}
	}

	if assignmentAttempt.Point == nil {
		// handle calculate result
		fmt.Println("Calculate result")
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

		_, err := biz.assignmentAttemptGetResultRepo.UpdateAssignmentAttempt(ctx, newAssignmentAttempt)
		if err != nil {
			return nil, err
		}
		return newAssignmentAttempt, nil
	}

	return assignmentAttempt, nil
}
