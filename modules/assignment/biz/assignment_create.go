package assignmentbiz

import (
	"context"
	"errors"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentRepo interface {
	CreateAssignment(ctx context.Context, data *assignmentmodel.Assignment, teacherId int) error
	CreateChoice(ctx context.Context, data *assignmentmodel.QuestionChoice) error
	CreateCorrectAnswer(ctx context.Context, data *assignmentmodel.QuestionCorrectAnswer) error
	CreateQuestionHierarchy(ctx context.Context, data *assignmentmodel.QuestionHierarchy) error
	CreateQuestion(ctx context.Context, data *assignmentmodel.Question) (*assignmentmodel.Question, error)
	CreateManyQuestionAssignment(ctx context.Context, data []*assignmentmodel.QuestionAssignment) error
	CreateAssignmentPlacement(ctx context.Context, data *assignmentmodel.AssignmentPlacement) error
}

type createAssignmentBiz struct {
	assignmentRepo AssignmentRepo
}

func NewAssignmentCreateBiz(assignmentRepo AssignmentRepo) *createAssignmentBiz {
	return &createAssignmentBiz{assignmentRepo: assignmentRepo}
}

func (biz *createAssignmentBiz) createNestedQuestions(
	ctx context.Context,
	questions []*assignmentmodel.QuestionCreate,
	teacherId int,
	parentId *int,
) ([]assignmentmodel.CreatedQuestionRelation, error) {
	var questionIds []assignmentmodel.CreatedQuestionRelation
	for _, question := range questions {
		var questionInfo assignmentmodel.Question
		questionInfo = question.Question
		questionInfo.TeacherId = &teacherId

		createdQuestion, err := biz.assignmentRepo.CreateQuestion(ctx, &questionInfo)
		if err != nil {
			return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionEntityName, err)
		}

		// add question id to questionIds if it is not a sub question
		if parentId == nil {
			questionRelation := assignmentmodel.CreatedQuestionRelation{
				QuestionId: createdQuestion.Id,
				Order:      *createdQuestion.Order,
				Point:      *createdQuestion.Point,
			}

			questionIds = append(questionIds, questionRelation)
		}
		// Create relational between question and parent question
		if parentId != nil {
			questionHierarchy := &assignmentmodel.QuestionHierarchy{
				ParentId:   *parentId,
				QuestionId: createdQuestion.Id,
				Order:      *createdQuestion.Order,
				Point:      *createdQuestion.Point,
			}
			err := biz.assignmentRepo.CreateQuestionHierarchy(ctx, questionHierarchy)
			if err != nil {
				return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionHierarchyEntityName, err)
			}
		}

		if createdQuestion.Type == assignmentmodel.SuperQuestion {
			_, err := biz.createNestedQuestions(ctx, createdQuestion.Questions, teacherId, &createdQuestion.Id)
			if err != nil {
				return nil, err
			}
		} else {
			for _, choice := range question.Choices {
				choice.QuestionId = createdQuestion.Id
				if err := biz.assignmentRepo.CreateChoice(ctx, &choice); err != nil {
					return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionChoiceEntityName, err)
				}
			}

			for _, correctAnswer := range question.CorrectAnswer {
				correctAnswer.QuestionId = createdQuestion.Id
				if err := biz.assignmentRepo.CreateCorrectAnswer(ctx, &correctAnswer); err != nil {
					return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionCorrectAnswerEntityName, err)
				}
			}
		}
	}
	return questionIds, nil
}

func (biz *createAssignmentBiz) CreateAssignment(
	ctx context.Context,
	data *assignmentmodel.AssignmentCreate,
	teacherId int,
) (*assignmentmodel.Assignment, error) {
	var assignment *assignmentmodel.Assignment
	assignment = &data.Assignment

	var questions []*assignmentmodel.QuestionCreate
	questions = data.Questions

	if len(questions) == 0 {
		return nil, common.ErrInvalidRequest(errors.New("Please provide at least one question"))
	}

	if err := biz.assignmentRepo.CreateAssignment(ctx, assignment, teacherId); err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentEntityName, err)
	}

	questionRelations, err := biz.createNestedQuestions(ctx, questions, teacherId, nil)
	if err != nil {
		return nil, err
	}

	for _, questionId := range questionRelations {
		questionAssignment := &assignmentmodel.QuestionAssignment{
			AssignmentId: assignment.Id,
			QuestionId:   questionId.QuestionId,
			Order:        questionId.Order,
			Point:        questionId.Point,
		}
		if err := biz.assignmentRepo.CreateManyQuestionAssignment(ctx, []*assignmentmodel.QuestionAssignment{questionAssignment}); err != nil {
			return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionAssignmentEntityName, err)
		}
	}

	// Create assignment placement
	assignmentPlacement := &assignmentmodel.AssignmentPlacement{}
	assignmentPlacement.StartTime = *data.StartTime
	assignmentPlacement.EndTime = *data.EndTime
	assignmentPlacement.AssignmentId = assignment.Id

	if data.CourseId != nil {
		assignmentPlacement.CourseId = data.CourseId
		if data.SectionId != nil {
			if data.LectureId != nil {
				assignmentPlacement.LectureId = data.LectureId
				assignmentPlacement.SectionId = data.SectionId
			} else {
				assignmentPlacement.SectionId = data.SectionId
			}
		} else {
			assignmentPlacement.CourseId = data.CourseId
		}

		if err := biz.assignmentRepo.CreateAssignmentPlacement(ctx, assignmentPlacement); err != nil {
			return nil, common.ErrCannotCreateEntity(assignmentmodel.AssignmentPlacementEntityName, err)
		}
	}

	return assignment, nil
}
