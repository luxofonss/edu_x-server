package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateAssignment(ctx context.Context, data *assignmentmodel.AssignmentCreate, teacherId int) error {
	data.TeacherId = &teacherId
	db := repo.db.Table(assignmentmodel.AssignmentCreate{}.TableName())
	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
