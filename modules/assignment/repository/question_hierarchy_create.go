package assignmentrepo

import (
	"context"
	"fmt"

	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateQuestionHierarchy(ctx context.Context, data *assignmentmodel.QuestionHierarchy) error {
	db := repo.db.Table(assignmentmodel.QuestionHierarchy{}.TableName())
	fmt.Println("data:: ", data)
	if err := db.Create(&data).Error; err != nil {
		fmt.Println("err:: ", err)
		return err
	}

	return nil
}
