package assignmentrepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateAssignmentAttempt(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	attemptData := assignmentmodel.AssignmentAttempt{
		AssignmentId:         data.AssignmentId,
		UserId:               data.UserId,
		AssignmentTimeMillis: data.AssignmentTimeMillis,
	}

	if err := db.Create(&attemptData).Error; err != nil {
		return nil, err
	}

	return &attemptData, nil
}

func (repo *assignmentRepo) CheckMultipleAttempt(ctx context.Context, assignmentId uuid.UUID) (bool, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())

	var assignment assignmentmodel.Assignment

	err := db.Where("id = ?", assignmentId).First(&assignment).Error
	if err != nil {
		return false, err
	}

	return assignment.MultipleAttempt, nil
}

func (repo *assignmentRepo) GetAllAttemptInAssignment(ctx context.Context, assignmentId uuid.UUID, userId uuid.UUID) ([]assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var attempts []assignmentmodel.AssignmentAttempt

	err := db.Order("created_at asc").Where("assignment_id = ? AND user_id = ?", assignmentId, userId).Find(&attempts).Error
	if err != nil {
		return nil, err
	}

	return attempts, nil
}

func (repo *assignmentRepo) GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var assigment *assignmentmodel.AssignmentAttempt

	err := db.Where("id = ?", assignmentAttemptId).First(&assigment).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assigment, nil
}

func (repo *assignmentRepo) GetAnswer(ctx context.Context, questionId uuid.UUID, userId uuid.UUID) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	var answer *assignmentmodel.QuestionAnswer

	err := db.Where("question_id = ? AND user_id = ?", questionId, userId).First(&answer).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return answer, nil
}

func (repo *assignmentRepo) SubmitQuestionAnswer(
	ctx context.Context,
	data *assignmentmodel.QuestionAnswer,
) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "question_id"}, {Name: "assignment_attempt_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"selected_option_id", "text_answer"}),
	}).Create(&data).Error; err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return data, nil
}

func (repo *assignmentRepo) GetAssignmentAttemptById(
	ctx context.Context,
	id uuid.UUID,
) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var assignmentAttempt *assignmentmodel.AssignmentAttempt

	if err := db.Where("id = ?", id).First(&assignmentAttempt).Error; err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	var assignmentAttemptRes *assignmentmodel.AssignmentAttempt

	db.Preload("Assignment").
		Preload("Assignment.Questions").
		Preload("Assignment.Questions.Choices").
		Preload("Assignment.Questions.Answers", "user_id = ? AND assignment_attempt_id = ?", assignmentAttempt.UserId, assignmentAttempt.Id).
		Preload("Assignment.Questions.Answers.Feedback")

	if err := db.Where("assignment_attempts.id = ?", id).First(&assignmentAttemptRes).Error; err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttemptRes, nil
}

func (repo *assignmentRepo) GetAssignmentAttempt(
	ctx context.Context,
	filter *assignmentmodel.AssignmentAttemptFilter,
	paging *common.Paging,
	moreKeys ...string,
) ([]*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	if filter != nil {
		if filter.Id != uuid.Nil {
			db = db.Where("id = ?", filter.Id)
		}

		if filter.AssignmentId != uuid.Nil {
			db = db.Where("assignment_id = ?", filter.AssignmentId)
		}
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var assignmentAttempt []*assignmentmodel.AssignmentAttempt
	err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&assignmentAttempt).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assignmentAttempt, nil
}

func (repo *assignmentRepo) UpdateAssignmentAttempt(
	ctx context.Context,
	data *assignmentmodel.AssignmentAttempt,
) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	if err := db.Where("id = ?", data.Id).Updates(&data).Error; err != nil {
		return nil, common.ErrCannotUpdateEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return data, nil
}

func (repo *assignmentRepo) GetAllAttemptByAssignmentId(
	ctx context.Context,
	assignmentId uuid.UUID,
	moreKeys ...string,
) ([]assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var attempts = []assignmentmodel.AssignmentAttempt{}

	err := db.Where("assignment_id = ?", assignmentId).Find(&attempts).Error
	if err != nil {
		return nil, err
	}

	return attempts, nil
}

func (repo *assignmentRepo) UpdateQuestionAnswer(
	ctx context.Context,
	data *assignmentmodel.QuestionAnswer,
) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	fmt.Println("data:: ", data.Id)
	if err := db.Where("id = ? ", data.Id).Updates(&data).Error; err != nil {
		return nil, common.ErrCannotUpdateEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return data, nil
}

func (repo *assignmentRepo) UpdateLongAnswerScore(ctx context.Context, questionAnswerId uuid.UUID, point int) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	if err := db.Where("id = ?", questionAnswerId).Update("score", point).Error; err != nil {
		return nil, common.ErrCannotUpdateEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	db2 := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())
	var questionAnswer *assignmentmodel.QuestionAnswer

	if err := db2.Where("id = ?", questionAnswerId).First(&questionAnswer).Error; err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	fmt.Println("questionAnswer:: ", questionAnswer.Score)
	return questionAnswer, nil
}
