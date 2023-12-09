package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateChoice(ctx context.Context, data *assignmentmodel.QuestionChoice) error {
	db := repo.db.Table(assignmentmodel.QuestionChoice{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
