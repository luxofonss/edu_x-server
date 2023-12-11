package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAllAttemptInAssignment(ctx context.Context, assignmentPlacementId int, userId int) ([]assignmentmodel.AssignmentAttempt, error) {
	db := repo.db.Table(assignmentmodel.AssignmentAttempt{}.TableName())

	var attempts []assignmentmodel.AssignmentAttempt

	err := db.Where("assignment_placement_id = ? AND user_id = ?", assignmentPlacementId, userId).Find(&attempts).Error
	if err != nil {
		return nil, err
	}

	return attempts, nil
}
