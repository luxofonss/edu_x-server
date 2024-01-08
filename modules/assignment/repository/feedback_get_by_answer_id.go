package assignmentrepo

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

func (repo *assignmentRepo) GetFeedbacksByAnswerId(ctx context.Context, answerId uuid.UUID) ([]*assignmentmodel.Feedback, error) {
	var data []*assignmentmodel.Feedback
	if err := repo.db.Where("question_answer_id = ?", answerId).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
