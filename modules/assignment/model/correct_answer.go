package assignmentmodel

import "server/common"

const QuestionCorrectAnswerEntityName = "QuestionCorrectAnswer"

type QuestionCorrectAnswer struct {
	common.SimpleSqlModel `json:",inline"`
	CorrectTextAnswer     string    `json:"correct_text_answer" gorm:"column:correct_text_answer;"`
	AnswerExplain         string    `json:"answer_explain" gorm:"column:answer_explain;"`
	QuestionId            int       `json:"question_id" gorm:"column:question_id;primaryKey"`
	Question              *Question `json:"question" gorm:"foreignKey:QuestionId;"`
}

func (QuestionCorrectAnswer) TableName() string { return "correct_answers" }
