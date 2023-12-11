package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAssignment(ctx context.Context, id int) (*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data assignmentmodel.Assignment

	db = db.Preload("Questions").Preload("Questions.Choices").Preload("Questions.Questions")
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *assignmentRepo) GetAssignmentByCourseId(ctx context.Context, id int) ([]*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data []*assignmentmodel.Assignment

	db = db.Joins("JOIN assignment_placement ON assignment_placement.assignment_id = assignments.id").
		Where("assignment_placement.course_id = ?", id)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *assignmentRepo) GetAssignmentBySectionId(ctx context.Context, id int) ([]*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data []*assignmentmodel.Assignment

	db = db.Joins("JOIN assignment_placement ON assignment_placement.assignment_id = assignments.id").
		Where("assignment_placement.section_id = ?", id)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *assignmentRepo) GetAssignmentByLectureId(ctx context.Context, id int) ([]*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data []*assignmentmodel.Assignment

	db = db.Joins("JOIN assignment_placement ON assignment_placement.assignment_id = assignments.id").
		Where("assignment_placement.lecture_id = ?", id)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
