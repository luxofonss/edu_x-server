package coursemodel

import (
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

const CourseEntityName = "Course"

type CourseLevel string

const (
	BEGINER            CourseLevel = "beginer"
	ELEMENTARY         CourseLevel = "elementary"
	INTERMEDIATE       CourseLevel = "intermediate"
	UPPER_INTERMEDIATE CourseLevel = "upper_intermediate"
	ADVANCED           CourseLevel = "advanced"
	PROFICIENCY        CourseLevel = "proficiency"
)

type Course struct {
	common.SQLModel `json:",inline"`
	Name            string                        `json:"name" gorm:"column:name;"`
	Description     string                        `json:"description" gorm:"column:description;"`
	BackgroundImg   string                        `json:"background_img" gorm:"column:background_img;"`
	StartDate       string                        `json:"start_date" gorm:"column:start_date;"`
	EndDate         string                        `json:"end_date" gorm:"column:end_date;"`
	Price           float64                       `json:"price" gorm:"column:price;"`
	Currency        string                        `json:"currency" gorm:"column:currency;"`
	Level           CourseLevel                   `json:"level" gorm:"column:level;"`
	IsVerified      bool                          `json:"is_verified" gorm:"column:is_verified;"`
	SubjectId       int                           `json:"subject_id" gorm:"column:subject_id;"`
	Grade           int                           `json:"grade" gorm:"column:grade;"`
	Code            string                        `json:"code" gorm:"column:code;"`
	TeacherId       int                           `json:"teacher_id" gorm:"column:teacher_id;"`
	SchoolId        *int                          `json:"school_id" gorm:"column:school_id;"`
	Sections        []Section                     `json:"sections" gorm:"-"`
	CourseInfos     []CourseInfo                  `json:"course_infos" gorm:"-"`
	Assignment      []*assignmentmodel.Assignment `json:"assignments" gorm:"many2many:assignment_placement;"`
}

func (Course) TableName() string { return "courses" }

func (c *Course) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeCourse)
}

type CourseCreate struct {
	Course      Course        `json:"course"`
	CourseInfos []*CourseInfo `json:"course_infos"`
}

var (
	ErrCourseNotFound = common.NewCustomError(nil, "course not found", "ErrCourseNotFound")
)
