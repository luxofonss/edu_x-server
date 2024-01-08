package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const AssignmentAttemptEntityName = "AssignmentAttempt"

type AssignmentAttempt struct {
	common.SQLModel      `json:",inline"`
	UserId               uuid.UUID   `json:"user_id" gorm:"column:user_id;type:uuid;"`
	AssignmentId         uuid.UUID   `json:"assignment_id" gorm:"column:assignment_id;type:uuid;"`
	AssignmentTimeMillis int64       `json:"assignment_time_millis" gorm:"column:assignment_time_millis;"`
	Point                int         `json:"point" gorm:"column:point;"`
	TeacherComment       string      `json:"teacher_comment" gorm:"column:teacher_comment;"`
	FinishedAt           *string     `json:"finished_at" gorm:"column:finished_at;"`
	Assignment           *Assignment `json:"assignment" gorm:"foreignKey:AssignmentId;"`
}

func (AssignmentAttempt) TableName() string {
	return "assignment_attempts"
}

type AssignmentAttemptCreate struct {
	common.SQLModel      `json:",inline"`
	UserId               uuid.UUID `json:"user_id" gorm:"column:user_id;"`
	AssignmentId         uuid.UUID `json:"assignment_id" gorm:"column:assignment_id;"`
	AssignmentTimeMillis int64     `json:"assignment_time_millis" gorm:"column:assignment_time_millis;"`
}
