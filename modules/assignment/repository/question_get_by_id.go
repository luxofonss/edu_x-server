package assignmentrepo

import (
	"context"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetQuestionById(ctx context.Context, questionId int) (*assignmentmodel.Question, error) {
	db := repo.db.Table(assignmentmodel.Question{}.TableName())

	var question *assignmentmodel.Question

	err := db.Where("id = ?", questionId).First(&question).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.QuestionEntityName, err)
	}

	return question, nil
}
