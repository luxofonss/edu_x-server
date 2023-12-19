package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const AssignmentEntityName = "Assignment"

type AssignmentType string

const (
	MidTerm AssignmentType = "mid"
	Final   AssignmentType = "final"
	Regular AssignmentType = "regular"
)

type Assignment struct {
	common.SQLModel   `json:",inline"`
	Title             string               `json:"title" gorm:"column:title;"`
	Description       string               `json:"description" gorm:"column:description;"`
	TotalPoint        int                  `json:"total_point" gorm:"column:total_point;"`
	MultipleAttempt   bool                 `json:"multiple_attempt" gorm:"column:multiple_attempt;"`
	StartTime         *string              `json:"start_time" gorm:"column:start_time;"`
	EndTime           *string              `json:"end_time" gorm:"column:end_time;"`
	Type              AssignmentType       `json:"type" gorm:"column:type;"`
	PlacementId       uuid.UUID            `json:"placement_id" gorm:"column:placement_id;type:uuid;"`
	TeacherId         uuid.UUID            `json:"teacher_id" gorm:"column:teacher_id;type:uuid;default:NULL;"`
	SchoolId          *uuid.UUID           `json:"school_id" gorm:"column:school_id;type:uuid;default:NULL;"`
	SubjectId         uuid.UUID            `json:"subject_id" gorm:"column:subject_id;type:uuid;"`
	Questions         []*Question          `json:"questions" gorm:"foreignkey:AssignmentId;"`
	AssignmentAttempt []*AssignmentAttempt `json:"assignment_attempts" gorm:"foreignKey:AssignmentId"`
}

func (Assignment) TableName() string { return "assignments" }

type AssignmentCreate struct {
	Assignment      `,json:"inline"`
	MultipleAttempt bool           `json:"multiple_attempt" gorm:"column:multiple_attempt;"`
	StartTime       *string        `json:"start_time" gorm:"column:start_time;"`
	EndTime         *string        `json:"end_time" gorm:"column:end_time;"`
	Type            AssignmentType `json:"type" gorm:"column:type;"`
	PlacementId     uuid.UUID      `json:"placement_id" gorm:"column:placement_id;type:uuid;"`
}

func (AssignmentCreate) TableName() string {
	return "assignments"
}

var (
	ErrQuestionNotInAssignment = common.NewCustomError(nil, "question not in assignment", "ErrQuestionNotInAssignment")
)
