package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const QuestionAnswerEntityName = "QuestionAnswer"

type QuestionAnswer struct {
	common.SQLModel     `json:",inline"`
	UserId              uuid.UUID       `json:"user_id" gorm:"column:user_id;type:uuid;"`
	QuestionId          uuid.UUID       `json:"question_id" gorm:"column:question_id;type:uuid;"`
	AssignmentAttemptId uuid.UUID       `json:"assignment_attempt_id" gorm:"column:assignment_attempt_id;type:uuid;"`
	SelectedOptionId    *uuid.UUID      `json:"selected_option_id" gorm:"column:selected_option_id;type:uuid;"`
	TextAnswer          string          `json:"text_answer" gorm:"column:text_answer;"`
	Score               *int            `json:"score" gorm:"column:score;"`
	Question            *Question       `json:"question" gorm:"foreignKey:QuestionId;"`
	Feedback            []*Feedback     `json:"feedback" gorm:"foreignKey:QuestionAnswerId;"`
	Choice              *QuestionChoice `json:"option" gorm:"foreignKey:SelectedOptionId;"`
}

func (QuestionAnswer) TableName() string { return "question_answers" }
