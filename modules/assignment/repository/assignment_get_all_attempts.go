package assignmentrepo

import (
	"context"

	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAllAttemptInAssignment(ctx context.Context, assignmentId uuid.UUID, userId uuid.UUID) ([]assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var attempts []assignmentmodel.AssignmentAttempt

	err := db.Where("assignment_id = ? AND user_id = ?", assignmentId, userId).Find(&attempts).Error
	if err != nil {
		return nil, err
	}

	return attempts, nil
}
