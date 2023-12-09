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
	common.SQLModel `json:",inline"`
	Title           string      `json:"title" gorm:"column:title;"`
	Description     string      `json:"description" gorm:"column:description;"`
	TotalPoint      int         `json:"total_point" gorm:"column:total_point;"`
	TeacherId       *int        `json:"teacher_id" gorm:"column:teacher_id;"`
	SchoolId        *int        `json:"school_id" gorm:"column:school_id;"`
	SubjectId       int         `json:"subject_id" gorm:"column:subject_id;"`
	Questions       []*Question `json:"questions" gorm:"many2many:question_assignment;"`
}

func (Assignment) TableName() string { return "assignments" }

func (a *Assignment) Mask(isAdminOrOwner bool) {
	a.GenUID(common.DbTypeAssignment)
}

type AssignmentCreate struct {
	StartTime  *string           `json:"start_time"`
	EndTime    *string           `json:"end_time"`
	CourseId   *int              `json:"course_id" gorm:"column:course_id;"`
	LectureId  *int              `json:"lecture_id" gorm:"column:lecture_id;"`
	SectionId  *int              `json:"section_id" gorm:"column:section_id;"`
	Type       *AssignmentType   `json:"type"`
	Assignment Assignment        `json:"assignment"`
	Questions  []*QuestionCreate `json:"questions"`
}
