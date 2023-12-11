package assignmentrepo

import (
	"context"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAssignmentByAssignmentPlacementId(
	ctx context.Context,
	AssignmentPlacementId int,
) (*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.AssignmentPlacement{}.TableName())

	db.Preload("Assignment")

	var assignment *assignmentmodel.AssignmentPlacement

	err := db.Where("id = ?", AssignmentPlacementId).First(&assignment).Error
	if err != nil {
		return nil, common.ErrCannotGetEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return assignment.Assignment, nil
}
