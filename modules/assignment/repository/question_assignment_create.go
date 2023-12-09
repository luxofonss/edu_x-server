package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateManyQuestionAssignment(ctx context.Context, data []*assignmentmodel.QuestionAssignment) error {
	db := repo.db.Table(assignmentmodel.QuestionAssignment{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
