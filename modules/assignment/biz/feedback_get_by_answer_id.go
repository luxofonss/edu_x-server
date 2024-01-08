package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type GetFeedbackByAnswerIdRepo interface {
	GetFeedbacksByAnswerId(ctx context.Context, answerId uuid.UUID) ([]*assignmentmodel.Feedback, error)
}

type getFeedbackByAnswerIdBiz struct {
	repo GetFeedbackByAnswerIdRepo
}

func NewGetFeedbackByAnswerIdBiz(repo GetFeedbackByAnswerIdRepo) *getFeedbackByAnswerIdBiz {
	return &getFeedbackByAnswerIdBiz{repo: repo}
}

func (biz *getFeedbackByAnswerIdBiz) GetFeedbackByAnswerId(ctx context.Context, answerId uuid.UUID) ([]*assignmentmodel.Feedback, error) {
	// :TODO check if user can access this feedback

	return biz.repo.GetFeedbacksByAnswerId(ctx, answerId)
}
