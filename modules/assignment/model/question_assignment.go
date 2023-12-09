package assignmentmodel

import (
	"gorm.io/gorm"
	"server/common"
)

const QuestionAssignmentEntityName = "QuestionAssignment"

type QuestionAssignment struct {
	common.SimpleSqlModel `json:",inline"`
	AssignmentId          int `json:"assignment_id" gorm:"column:assignment_id;"`
	QuestionId            int `json:"question_id" gorm:"column:question_id;"`
	Order                 int `json:"order" gorm:"column:order;"`
	Point                 int `json:"point" gorm:"column:point;"`
}

func (QuestionAssignment) TableName() string { return "question_assignment" }

func (QuestionAssignment) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.Select("AssignmentId", "QuestionId", "Order", "Point")

	return
}
