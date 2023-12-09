package assignmentmodel

import "server/common"

const QuestionChoiceEntityName = "QuestionChoice"

type QuestionChoice struct {
	common.SQLModel `json:",inline"`
	Content         string    `json:"content" gorm:"column:content;"`
	Order           int       `json:"order" gorm:"column:order;"`
	IsCorrect       bool      `json:"is_correct" gorm:"column:is_correct;"`
	AnswerExplain   string    `json:"answer_explain" gorm:"column:answer_explain;"`
	QuestionId      int       `json:"question_id" gorm:"column:question_id;"`
	Question        *Question `json:"question" gorm:"foreignKey:QuestionId;"`
}

func (QuestionChoice) TableName() string { return "question_choices" }

func (q *QuestionChoice) Mask(isAdminOrOwner bool) {
	q.GenUID(common.DbTypeQuestionChoice)
}
