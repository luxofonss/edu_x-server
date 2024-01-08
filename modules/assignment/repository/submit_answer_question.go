package assignmentrepo

import (
	"context"
	"gorm.io/gorm/clause"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

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
