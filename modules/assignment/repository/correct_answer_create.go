package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateCorrectAnswer(ctx context.Context, data *assignmentmodel.QuestionCorrectAnswer) error {
	db := repo.db.Table(assignmentmodel.QuestionCorrectAnswer{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
