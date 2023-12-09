package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateAssignmentPlacement(ctx context.Context, data *assignmentmodel.AssignmentPlacement) error {
	db := repo.db.Table(assignmentmodel.AssignmentPlacement{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
