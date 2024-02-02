package assignmentdto

type FeedbackQuestionAnswerRequest struct {
	Message    string `json:"message"`
	FeedbackId string `json:"feedback_id,omitempty"`
	Id         string `json:"id"`
	Type       string `json:"type"`
}

type LongAnswerScoreRequest struct {
	Point int `json:"point"`
}
