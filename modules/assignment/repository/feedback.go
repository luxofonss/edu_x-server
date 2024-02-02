package assignmentrepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) CreateFeedbackAnswer(ctx context.Context, data *assignmentmodel.Feedback) (*assignmentmodel.Feedback, error) {
	db := repo.db

	fmt.Println("data", data.Id, data.FeedbackId)

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"message"}),
	}).Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}

func (repo *assignmentRepo) GetFeedbacksByAnswerId(ctx context.Context, answerId uuid.UUID) ([]*assignmentmodel.Feedback, error) {
	var data []*assignmentmodel.Feedback
	if err := repo.db.Where("question_answer_id = ?", answerId).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
