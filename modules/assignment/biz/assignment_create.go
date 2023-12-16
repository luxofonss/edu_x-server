package assignmentbiz

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentRepo interface {
	CreateAssignment(ctx context.Context, data *assignmentmodel.AssignmentCreate, teacherId uuid.UUID) error
	CreateChoice(ctx context.Context, data *assignmentmodel.QuestionChoice) error
	CreateCorrectAnswer(ctx context.Context, data *assignmentmodel.QuestionCorrectAnswer) error
	CreateQuestion(ctx context.Context, data *assignmentmodel.Question) (*assignmentmodel.Question, error)
	CreateAssignmentPlacement(ctx context.Context, data *assignmentmodel.AssignmentPlacement) error
}

type createAssignmentBiz struct {
	assignmentRepo AssignmentRepo
}

func NewAssignmentCreateBiz(assignmentRepo AssignmentRepo) *createAssignmentBiz {
	return &createAssignmentBiz{assignmentRepo: assignmentRepo}
}

func (biz *createAssignmentBiz) CreateAssignment(
	ctx context.Context,
	data *assignmentmodel.AssignmentCreate,
	teacherId uuid.UUID,
) (*assignmentmodel.AssignmentCreate, error) {

	data.AssignmentPlacement = append(data.AssignmentPlacement, &assignmentmodel.AssignmentPlacement{
		StartTime: *data.StartTime,
		EndTime:   *data.EndTime,
		Type:      data.Type,
		CourseId:  data.CourseId,
		LectureId: data.LectureId,
		SectionId: data.SectionId,
	})

	for _, question := range data.Questions {
		question.TeacherId = teacherId

		for _, childQuestion := range question.Questions {
			childQuestion.TeacherId = teacherId
			childQuestion.AssignmentId = uuid.Nil
		}
	}

	if err := biz.assignmentRepo.CreateAssignment(ctx, data, teacherId); err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentEntityName, err)
	}

	fmt.Println("Hello")

	return data, nil
}
