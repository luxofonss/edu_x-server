package assignmentdto

import (
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type QuestionAnswerRequest struct {
	QuestionId          string   `json:"question_id"`
	AssignmentAttemptId string   `json:"assignment_attempt_id"`
	SelectedOptionId    []string `json:"selected_option_id,omitempty"`
	UserId              string   `json:"user_id"`
	TextAnswer          string   `json:"text_answer,omitempty"`
}

func (data *QuestionAnswerRequest) ToQuestionAnswerEntity() ([]*assignmentmodel.QuestionAnswer, error) {
	var result []*assignmentmodel.QuestionAnswer

	for _, optionId := range data.SelectedOptionId {
		if optionId != "" {
			id := uuid.MustParse(optionId)
			result = append(result, &assignmentmodel.QuestionAnswer{
				UserId:              uuid.MustParse(data.UserId),
				QuestionId:          uuid.MustParse(data.QuestionId),
				AssignmentAttemptId: uuid.MustParse(data.AssignmentAttemptId),
				SelectedOptionId:    &id,
				TextAnswer:          data.TextAnswer,
			})

		}
	}

	return result, nil
}
