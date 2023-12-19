package assignmentrepo

import (
	"context"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateAssignmentAttempt(ctx context.Context, data *assignmentmodel.AssignmentAttemptCreate) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	attemptData := assignmentmodel.AssignmentAttempt{
		AssignmentId: data.AssignmentId,
		UserId:       data.UserId,
	}

	if err := db.Create(&attemptData).Error; err != nil {
		return nil, err
	}

	return &attemptData, nil
}
