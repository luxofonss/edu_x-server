package assignmentmodel

import "server/common"

const AssignmentPlacementEntityName = "AssignmentPlacement"

type AssignmentPlacement struct {
	common.SQLModel `json:",inline"`
	AssignmentId    int    `json:"assignment_id" gorm:"column:assignment_id;"`
	CourseId        *int   `json:"course_id" gorm:"column:course_id;"`
	LectureId       *int   `json:"lecture_id" gorm:"column:lecture_id;"`
	SectionId       *int   `json:"section_id" gorm:"column:section_id;"`
	StartTime       string `json:"start_time" gorm:"column:start_time;"`
	EndTime         string `json:"end_time" gorm:"column:end_time;"`
}

func (AssignmentPlacement) TableName() string { return "assignment_placement" }
