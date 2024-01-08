package assignmentrepo

import (
	"context"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateFeedbackAnswer(ctx context.Context, data *assignmentmodel.Feedback) (*assignmentmodel.Feedback, error) {
	db := repo.db

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
