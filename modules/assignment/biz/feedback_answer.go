package assignmentbiz

import (
	"context"
	assignmentmodel "server/modules/assignment/model"
)

type FeedbackAnswerRepo interface {
	CreateFeedbackAnswer(ctx context.Context, data *assignmentmodel.Feedback) (*assignmentmodel.Feedback, error)
}

type feedbackAnswerBiz struct {
	repo FeedbackAnswerRepo
}

func NewFeedbackAnswerBiz(repo FeedbackAnswerRepo) *feedbackAnswerBiz {
	return &feedbackAnswerBiz{repo: repo}
}

func (biz *feedbackAnswerBiz) CreateFeedbackAnswer(ctx context.Context, data *assignmentmodel.Feedback) (*assignmentmodel.Feedback, error) {
	feedback, err := biz.repo.CreateFeedbackAnswer(ctx, data)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}
