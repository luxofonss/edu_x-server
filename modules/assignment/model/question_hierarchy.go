package assignmentmodel

import "server/common"

const QuestionHierarchyEntityName = "QuestionHierarchy"

type QuestionHierarchy struct {
	common.SimpleSqlModel `json:",inline"`
	QuestionId            int `json:"question_id" gorm:"column:child_id;"`
	ParentId              int `json:"parent_id" gorm:"column:parent_id;"`
	Order                 int `json:"order" gorm:"column:order;"`
	Point                 int `json:"point" gorm:"column:point;"`
}

func (QuestionHierarchy) TableName() string { return "question_hierarchies" }
