package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
	usermodel "server/modules/user/model"
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
	SubjectId       uuid.UUID                     `json:"subject_id" gorm:"column:subject_id;type:uuid;"`
	Grade           int                           `json:"grade" gorm:"column:grade;"`
	Code            string                        `json:"code" gorm:"column:code;"`
	TeacherId       uuid.UUID                     `json:"teacher_id" gorm:"column:teacher_id;type:uuid;"`
	SchoolId        *uuid.UUID                    `json:"school_id" gorm:"column:school_id;type:uuid;default:NULL;"`
	Sections        []*Section                    `json:"sections" gorm:"foreignKey:CourseId;"`
	CourseInfos     []*CourseInfo                 `json:"course_infos" gorm:"foreignKey:CourseId;"`
	Assignment      []*assignmentmodel.Assignment `json:"assignments" gorm:"many2many:assignment_placement;"`
	Subject         *Subject                      `json:"subject" gorm:"foreignKey:SubjectId;"`
	Teacher         *usermodel.User               `json:"teacher" gorm:"foreignKey:TeacherId;"`
}

func (Course) TableName() string { return "courses" }

type CourseCreate struct {
	Course      Course        `json:"course"`
	CourseInfos []*CourseInfo `json:"course_infos" gorm:"foreignKey:CourseId;"`
	Sections    []*Section    `json:"sections" gorm:"foreignKey:CourseId;"`
}

type CourseGet struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Description     string             `json:"description" gorm:"column:description;"`
	BackgroundImg   string             `json:"background_img" gorm:"column:background_img;"`
	StartDate       string             `json:"start_date" gorm:"column:start_date;"`
	EndDate         string             `json:"end_date" gorm:"column:end_date;"`
	Price           float64            `json:"price" gorm:"column:price;"`
	Currency        string             `json:"currency" gorm:"column:currency;"`
	Level           CourseLevel        `json:"level" gorm:"column:level;"`
	IsVerified      bool               `json:"is_verified" gorm:"column:is_verified;"`
	SubjectId       uuid.UUID          `json:"subject_id" gorm:"column:subject_id;type:uuid;"`
	Grade           int                `json:"grade" gorm:"column:grade;"`
	TeacherId       uuid.UUID          `json:"-" gorm:"column:teacher_id;type:uuid;"`
	Subject         *SimpleSubjectGet  `json:"subject" gorm:"foreignKey:SubjectId;"`
	Teacher         *common.SimpleUser `json:"teacher" gorm:"foreignKey:TeacherId;"`
	Sections        []*Section         `json:"sections" gorm:"foreignKey:CourseId;"`
}

var (
	ErrCourseNotFound = common.NewCustomError(nil, "course not found", "ErrCourseNotFound")
)
