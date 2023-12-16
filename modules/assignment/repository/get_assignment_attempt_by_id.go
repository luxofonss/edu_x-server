package assignmentrepo

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAssigmentAttemptById(ctx context.Context, assignmentAttemptId uuid.UUID) (*assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var assigment *assignmentmodel.AssignmentAttempt

	err := db.Where("id = ?", assignmentAttemptId).First(&assigment).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.AssignmentAttemptEntityName, err)
	}

	return assigment, nil
}
