package assignmentrepo

import (
	"context"

	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) SubmitQuestionAnswer(ctx context.Context, data *assignmentmodel.QuestionAnswer) (*assignmentmodel.QuestionAnswer, error) {
	db := repo.db.Table(assignmentmodel.QuestionAnswer{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrCannotCreateEntity(assignmentmodel.QuestionAnswerEntityName, err)
	}

	return data, nil

}
