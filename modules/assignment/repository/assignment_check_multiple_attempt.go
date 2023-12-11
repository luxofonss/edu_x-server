package assignmentrepo

import (
	"context"
	"fmt"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CheckMultipleAttempt(ctx context.Context, assignmentPlacementId int) (bool, error) {
	db := repo.db.Table(assignmentmodel.AssignmentPlacement{}.TableName())

	var assignment assignmentmodel.AssignmentPlacement

	err := db.Where("id = ?", assignmentPlacementId).First(&assignment).Error
	if err != nil {
		return false, err
	}

	fmt.Println("assignment:: ", assignment.CanMultipleAttempt)
	return assignment.CanMultipleAttempt, nil
}
