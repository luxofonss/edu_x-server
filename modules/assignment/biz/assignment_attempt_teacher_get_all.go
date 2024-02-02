package assignmentbiz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	"time"
)

type AssignmentAttemptRepo interface {
	GetAllAttemptByAssignmentId(ctx context.Context, assignmentId uuid.UUID, moreKeys ...string) ([]assignmentmodel.AssignmentAttempt, error)
	GetAssignmentAttemptById(
		ctx context.Context,
		id uuid.UUID,
	) (*assignmentmodel.AssignmentAttempt, error)
	UpdateAssignmentAttempt(
		ctx context.Context,
		data *assignmentmodel.AssignmentAttempt,
	) (*assignmentmodel.AssignmentAttempt, error)
}

type AssignmentAttemptTeacherGetAllBiz struct {
	assignmentAttemptRepo AssignmentAttemptRepo
}

func NewAssignmentAttemptTeacherGetAllBiz(assignmentAttemptRepo AssignmentAttemptRepo) *AssignmentAttemptTeacherGetAllBiz {
	return &AssignmentAttemptTeacherGetAllBiz{assignmentAttemptRepo: assignmentAttemptRepo}
}

func (biz *AssignmentAttemptTeacherGetAllBiz) GetAllAssignmentAttemptByAssignmentId(ctx context.Context, assignmentId uuid.UUID) ([]assignmentmodel.AssignmentAttempt, error) {
	assignmentAttempts, err := biz.assignmentAttemptRepo.GetAllAttemptByAssignmentId(ctx, assignmentId, "User")
	if err != nil {
		return nil, err
	}

	var result []assignmentmodel.AssignmentAttempt
	for _, assignmentAttempt := range assignmentAttempts {
		if assignmentAttempt.Point == nil {
			//CALCULATE TOTAL POINT
			detailAttempt, err := biz.assignmentAttemptRepo.GetAssignmentAttemptById(ctx, assignmentAttempt.Id)
			if err != nil {
				return nil, err
			}

			newAssignmentAttempt := detailAttempt
			newAssignmentAttempt.Point = new(int)
			for _, question := range detailAttempt.Assignment.Questions {
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
				}
			}

			if newAssignmentAttempt.FinishedAt == nil {
				assignmentCreatedAt, _ := time.Parse(common.DateString, newAssignmentAttempt.CreatedAt.String())
				assignmentTimeMillis := newAssignmentAttempt.AssignmentTimeMillis
				maxSubmitTime := assignmentCreatedAt.Add(time.Millisecond * time.Duration(assignmentTimeMillis))
				maxTime := time.Time(maxSubmitTime)
				newAssignmentAttempt.FinishedAt = &maxTime
			}
			_, err = biz.assignmentAttemptRepo.UpdateAssignmentAttempt(ctx, newAssignmentAttempt)
			if err != nil {
				return nil, err
			}
			assignmentAttempt = *newAssignmentAttempt
		}

		result = append(result, assignmentAttempt)
	}
	return result, nil
}
