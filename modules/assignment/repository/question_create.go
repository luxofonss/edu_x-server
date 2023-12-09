package assignmentrepo

import (
	"context"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateManyQuestion(ctx context.Context, data []*assignmentmodel.Question) error {
	db := repo.db.Table(assignmentmodel.Question{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (repo *assignmentRepo) CreateQuestion(ctx context.Context, data *assignmentmodel.Question) (*assignmentmodel.Question, error) {
	db := repo.db.Table(assignmentmodel.Question{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
