package assignmentdto

import (
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	"time"
)

type AnswerResponse struct {
	Id               string  `json:"id"`
	SelectedOptionId *string `json:"selected_option_id,omitempty"`
	TextAnswer       string  `json:"text_answer,omitempty"`
}

type ChoiceResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Order   *int   `json:"order"`
}

type QuestionResponse struct {
	Id       string            `json:"id"`
	Title    string            `json:"title"`
	Image    *string           `json:"image"`
	AudioUrl *string           `json:"audio_url"`
	Type     string            `json:"type"`
	Level    string            `json:"level"`
	Order    *int              `json:"order"`
	Choices  []*ChoiceResponse `json:"choices"`
	Answer   *AnswerResponse   `json:"answer"`
}

type AssignmentResponse struct {
	Id              string              `json:"id"`
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	TotalPoint      int                 `json:"total_point"`
	MultipleAttempt bool                `json:"multiple_attempt"`
	Time            int                 `json:"time"`
	StartTime       *string             `json:"start_time"`
	EndTime         *string             `json:"end_time"`
	Type            string              `json:"type"`
	Questions       []*QuestionResponse `json:"questions"`
}

type AssignmentAttemptResponse struct {
	Id                   string             `json:"id"`
	CreatedAt            string             `json:"created_at"`
	UpdatedAt            string             `json:"updated_at"`
	UserId               string             `json:"user_id"`
	AssignmentId         string             `json:"assignment_id"`
	AssignmentTimeMillis int64              `json:"assignment_time_millis"`
	RemainingTime        int64              `json:"remaining_time"`
	Assignment           AssignmentResponse `json:"assignment"`
}

func ToAssignmentAttemptResponse(assignmentAttempt assignmentmodel.AssignmentAttempt) AssignmentAttemptResponse {
	// Return remaining time
	var remainingTime int64 = 0
	if assignmentAttempt.AssignmentTimeMillis != 0 {
		assignmentCreatedAt, err := time.Parse(common.DateString, assignmentAttempt.CreatedAt.String())
		if err != nil {
			remainingTime = 0
		}
		assignmentTimeMillis := assignmentAttempt.AssignmentTimeMillis
		maxSubmitTime := assignmentCreatedAt.Add(time.Millisecond * time.Duration(assignmentTimeMillis))

		timeNow := time.Now().Add(time.Hour * 7) // GTM +7
		if timeNow.After(maxSubmitTime) {
			remainingTime = 0
		} else {
			remainingTimeDuration := maxSubmitTime.Sub(time.Now())
			remainingTime = int64(remainingTimeDuration.Milliseconds()) - 7*60*60*1000 // Minus 7 hours GTM +7
		}
	}

	var questions []*QuestionResponse
	for _, question := range assignmentAttempt.Assignment.Questions {
		var choices []*ChoiceResponse
		for _, choice := range question.Choices {
			choices = append(choices, &ChoiceResponse{
				Id:      choice.Id.String(),
				Content: choice.Content,
				Order:   &choice.Order,
			})
		}

		var answer *AnswerResponse

		if len(question.Answers) > 0 {

			if question.Answers[0].SelectedOptionId != nil {
				selectedOptionId := question.Answers[0].SelectedOptionId.String()
				answer = &AnswerResponse{
					Id:               question.Answers[0].Id.String(),
					SelectedOptionId: &selectedOptionId,
					TextAnswer:       question.Answers[0].TextAnswer,
				}
			} else {
				answer = &AnswerResponse{
					Id:               question.Answers[0].Id.String(),
					SelectedOptionId: nil,
					TextAnswer:       question.Answers[0].TextAnswer,
				}
			}

		}

		var image *string
		if question.Image != nil {
			image = &question.Image.Url
		}

		questions = append(questions, &QuestionResponse{
			Id:       question.Id.String(),
			Title:    question.Title,
			Image:    image,
			AudioUrl: &question.AudioUrl,
			Type:     string(question.Type),
			Level:    string(question.Level),
			Order:    question.Order,
			Choices:  choices,
			Answer:   answer,
		})
	}

	return AssignmentAttemptResponse{
		Id:                   assignmentAttempt.Id.String(),
		CreatedAt:            assignmentAttempt.CreatedAt.String(),
		UpdatedAt:            assignmentAttempt.UpdatedAt.String(),
		UserId:               assignmentAttempt.UserId.String(),
		AssignmentId:         assignmentAttempt.AssignmentId.String(),
		AssignmentTimeMillis: assignmentAttempt.AssignmentTimeMillis,
		RemainingTime:        remainingTime,
		Assignment: AssignmentResponse{
			Id:              assignmentAttempt.Assignment.Id.String(),
			Title:           assignmentAttempt.Assignment.Title,
			Description:     assignmentAttempt.Assignment.Description,
			TotalPoint:      assignmentAttempt.Assignment.TotalPoint,
			MultipleAttempt: assignmentAttempt.Assignment.MultipleAttempt,
			Time:            assignmentAttempt.Assignment.Time,
			StartTime:       assignmentAttempt.Assignment.StartTime,
			EndTime:         assignmentAttempt.Assignment.EndTime,
			Type:            string(assignmentAttempt.Assignment.Type),
			Questions:       questions,
		},
	}
}
