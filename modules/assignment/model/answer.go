package assignmentmodel

import "server/common"

const QuestionAnswerEntityName = "QuestionAnswer"

type QuestionAnswer struct {
	common.SimpleSqlModel `json:",inline"`
	UserId                int    `json:"user_id" gorm:"column:user_id;"`
	QuestionId            int    `json:"question_id" gorm:"column:question_id;"`
	SelectedOptionId      int    `json:"selected_option_id" gorm:"column:selected_option_id;"`
	TextAnswer            string `json:"text_answer" gorm:"column:text_answer;"`
}

func (QuestionAnswer) TableName() string { return "question_answers" }
