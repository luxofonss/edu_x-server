package assignmentrepo

import (
	"context"
	assignmentmodel "server/modules/assignment/model"

	"github.com/google/uuid"
)

func (repo *assignmentRepo) CheckMultipleAttempt(ctx context.Context, assignmentId uuid.UUID) (bool, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())

	var assignment assignmentmodel.Assignment

	err := db.Where("id = ?", assignmentId).First(&assignment).Error
	if err != nil {
		return false, err
	}

	return assignment.MultipleAttempt, nil
}
