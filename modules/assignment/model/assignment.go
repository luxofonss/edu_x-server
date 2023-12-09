package assignmentmodel

import "server/common"

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
	TeacherId           *int                   `json:"teacher_id" gorm:"column:teacher_id;"`
	SchoolId            *int                   `json:"school_id" gorm:"column:school_id;"`
	SubjectId           int                    `json:"subject_id" gorm:"column:subject_id;"`
	Questions           []*Question            `json:"questions" gorm:"foreignkey:AssignmentId;"`
	AssignmentTurnIn    []*AssignmentTurnIn    `json:"assignment_turn_ins" gorm:"foreignkey:AssignmentId;association_foreignkey:ID"`
	AssignmentPlacement []*AssignmentPlacement `json:"assignment_placements" gorm:"foreignkey:AssignmentId;association_foreignkey:ID"`
}

func (Assignment) TableName() string { return "assignments" }

func (a *Assignment) Mask(isAdminOrOwner bool) {
	a.GenUID(common.DbTypeAssignment)
}

type AssignmentCreate struct {
	Assignment `,json:"inline"`
	StartTime  *string        `json:"start_time" gorm:"-"`
	EndTime    *string        `json:"end_time" gorm:"-"`
	CourseId   *int           `json:"course_id" gorm:"-;"`
	LectureId  *int           `json:"lecture_id" gorm:"-;"`
	SectionId  *int           `json:"section_id" gorm:"-;"`
	Type       AssignmentType `json:"type" gorm:"-;"`
}

func (AssignmentCreate) TableName() string {
	return "assignments"
}
