package assignmentmodel

import "server/common"

const AssignmentTurnInEntityName = "AssignmentTurnIn"

type AssignmentTurnIn struct {
	common.SimpleSqlModel `json:",inline"`
	UserId                int         `json:"user_id" gorm:"column:user_id;"`
	AssignmentId          int         `json:"assignment_id" gorm:"column:assignment_id;"`
	Point                 int         `json:"point" gorm:"column:point;"`
	TeacherComment        string      `json:"teacher_comment" gorm:"column:teacher_comment;"`
	FinishedAt            string      `json:"finished_at" gorm:"column:finished_at;"`
	Assignment            *Assignment `json:"assignment" gorm:"foreignkey:AssignmentId;association_foreignkey:ID"`
}

func (AssignmentTurnIn) TableName() string {
	return "assignment_turn_ins"
}
