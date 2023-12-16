package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const AssignmentPlacementEntityName = "AssignmentPlacement"

type AssignmentPlacement struct {
	common.SQLModel    `json:",inline"`
	AssignmentId       uuid.UUID            `json:"assignment_id" gorm:"column:assignment_id;type:uuid;"`
	CourseId           uuid.UUID            `json:"course_id" gorm:"column:course_id;type:uuid;default:NULL;"`
	LectureId          uuid.UUID            `json:"lecture_id" gorm:"column:lecture_id;type:uuid;default:NULL;"`
	CanMultipleAttempt bool                 `json:"multiple_attempts" gorm:"column:multiple_attempts;"`
	SectionId          uuid.UUID            `json:"section_id" gorm:"column:section_id;type:uuid;default:NULL;"`
	StartTime          string               `json:"start_time" gorm:"column:start_time;"`
	EndTime            string               `json:"end_time" gorm:"column:end_time;"`
	Type               AssignmentType       `json:"type" gorm:"column:type;"`
	Assignment         *Assignment          `json:"assignment" gorm:"foreignkey:AssignmentId;association_foreignkey:Id"`
	AssignmentAttempt  []*AssignmentAttempt `json:"assignment_attempts" gorm:"foreignKey:AssignmentPlacementId"`
}

func (AssignmentPlacement) TableName() string { return "assignment_placement" }
