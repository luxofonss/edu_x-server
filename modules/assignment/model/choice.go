package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const QuestionChoiceEntityName = "QuestionChoice"

type QuestionChoice struct {
	common.SQLModel `json:",inline"`
	Content         string    `json:"content" gorm:"column:content;"`
	Order           int       `json:"order" gorm:"column:order;"`
	IsCorrect       bool      `json:"is_correct" gorm:"column:is_correct;"`
	QuestionId      uuid.UUID `json:"question_id" gorm:"column:question_id;type:uuid;"`
	Question        *Question `json:"question" gorm:"foreignKey:QuestionId;"`
}

func (QuestionChoice) TableName() string { return "question_choices" }
