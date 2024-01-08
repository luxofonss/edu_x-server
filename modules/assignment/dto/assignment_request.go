package assignmentdto

type FeedbackQuestionAnswerRequest struct {
	Message    string `json:"message"`
	FeedbackId string `json:"feedback_id"`
}
