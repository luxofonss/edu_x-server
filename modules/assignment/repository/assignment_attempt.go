package assignmentrepo

import (
	"context"
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

	err := db.Where("assignment_id = ? AND user_id = ?", assignmentId, userId).Find(&attempts).Error
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

// Submit
func (repo *assignmentRepo) SubmitQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "question_id"}, {Name: "assignment_attempt_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"selected_option_id", "text_answer"}),
	}).Create(&data).Error; err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return data, nil

}
