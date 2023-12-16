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
	common.SQLModel     `json:",inline"`
	Title               string                 `json:"title" gorm:"column:title;"`
	Description         string                 `json:"description" gorm:"column:description;"`
	TotalPoint          int                    `json:"total_point" gorm:"column:total_point;"`
	TeacherId           uuid.UUID              `json:"teacher_id" gorm:"column:teacher_id;type:uuid;default:NULL;"`
	SchoolId            *uuid.UUID             `json:"school_id" gorm:"column:school_id;type:uuid;default:NULL;"`
	SubjectId           uuid.UUID              `json:"subject_id" gorm:"column:subject_id;type:uuid;"`
	Questions           []*Question            `json:"questions" gorm:"foreignkey:AssignmentId;"`
	AssignmentAttempt   []*AssignmentAttempt   `json:"assignment_attempts" gorm:"foreignKey:AssignmentId"`
	AssignmentPlacement []*AssignmentPlacement `json:"assignment_placements" gorm:"foreignkey:AssignmentId;association_foreignkey:Id"`
}

func (Assignment) TableName() string { return "assignments" }

type AssignmentCreate struct {
	Assignment `,json:"inline"`
	StartTime  *string        `json:"start_time" gorm:"-"`
	EndTime    *string        `json:"end_time" gorm:"-"`
	CourseId   uuid.UUID      `json:"course_id" gorm:"-;"`
	LectureId  uuid.UUID      `json:"lecture_id" gorm:"-;"`
	SectionId  uuid.UUID      `json:"section_id" gorm:"-;"`
	Type       AssignmentType `json:"type" gorm:"-;"`
}

func (AssignmentCreate) TableName() string {
	return "assignments"
}

var (
	ErrQuestionNotInAssignment = common.NewCustomError(nil, "question not in assignment", "ErrQuestionNotInAssignment")
)
