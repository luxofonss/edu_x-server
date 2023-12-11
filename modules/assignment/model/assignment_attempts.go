package assignmentmodel

import "server/common"

const AssignmentAttemptEntityName = "AssignmentAttempt"

type AssignmentAttempt struct {
	common.SQLModel       `json:",inline"`
	UserId                int                  `json:"user_id" gorm:"column:user_id;"`
	AssignmentPlacementId *int                 `json:"assignment_placement_id" gorm:"column:assignment_placement_id;"`
	AssignmentId          *int                 `json:"assignment_id" gorm:"column:assignment_id;"`
	Point                 int                  `json:"point" gorm:"column:point;"`
	TeacherComment        string               `json:"teacher_comment" gorm:"column:teacher_comment;"`
	FinishedAt            *string              `json:"finished_at" gorm:"column:finished_at;"`
	Assignment            *Assignment          `json:"assignment" gorm:"foreignKey:AssignmentId;"`
	AssignmentPlacement   *AssignmentPlacement `json:"assignment_placement" gorm:"foreignkey:AssignmentPlacementId;"`
}

func (AssignmentAttempt) TableName() string {
	return "assignment_attempts"
}

type AssignmentAttemptCreate struct {
	UserId                int  `json:"user_id" gorm:"column:user_id;"`
	AssignmentPlacementId *int `json:"assignment_placement_id" gorm:"column:assignment_placement_id;"`
	AssignmentId          *int `json:"assignment_id" gorm:"column:assignment_id;"`
}
