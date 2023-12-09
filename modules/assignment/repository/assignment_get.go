package assignmentrepo

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetAssignment(ctx context.Context, id int) (*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data assignmentmodel.Assignment

	db = db.Preload("Questions").Preload("Questions.Choices").Preload("Questions.CorrectAnswers")
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *assignmentRepo) GetAssignmentByCourseId(ctx context.Context, id int) (*assignmentmodel.Assignment, error) {
	db := repo.db.Table(assignmentmodel.Assignment{}.TableName())
	var data assignmentmodel.Assignment

	db = db.Preload("Questions").Preload("Questions.Choices").Preload("Questions.CorrectAnswers")
	if err := db.Where("course_id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
