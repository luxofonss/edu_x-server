package assignmentrepo

import (
	"context"
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateAssignment(ctx context.Context, data *assignmentmodel.AssignmentCreate, teacherId uuid.UUID) error {
	data.TeacherId = teacherId
	//data.SchoolId.UUID = uuid.Nil

	db := repo.db.Table(assignmentmodel.AssignmentCreate{}.TableName())
	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
