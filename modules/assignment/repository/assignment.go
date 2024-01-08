package assignmentrepo

import (
	"context"
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

// CREATE
func (repo *assignmentRepo) CreateAssignment(ctx context.Context, data *assignmentmodel.AssignmentCreate, teacherId uuid.UUID) error {
	data.TeacherId = teacherId
	//data.SchoolId.UUID = uuid.Nil

	db := repo.db.Table(assignmentmodel.AssignmentCreate{}.TableName())
	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// GET
func (repo *assignmentRepo) GetAssignmentById(ctx context.Context, id uuid.UUID) (*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data assignmentmodel.Assignment

	db = db.Preload("Questions").Preload("Questions.Choices").Preload("Questions.Questions")
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *assignmentRepo) GetAssignmentByPlacementId(ctx context.Context, id uuid.UUID) ([]*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data []*assignmentmodel.Assignment

	db = db.Where("placement_id = ?", id)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
