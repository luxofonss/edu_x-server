package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
)

type CourseInfoTypes string

const (
	INTEND      CourseInfoTypes = "intend"
	WHO         CourseInfoTypes = "who"
	REQUIREMENT CourseInfoTypes = "requirement"
	WELCOME_MSG CourseInfoTypes = "welcome_msg"
	CONGRAT_MSG CourseInfoTypes = "congrat_msg"
)
const CourseInfoEntityName = "CourseInfo"

type CourseInfo struct {
	common.SQLModel `json:",inline"`
	CourseId        uuid.UUID       `json:"course_id" gorm:"column:course_id;type:uuid;"`
	Content         string          `json:"content" gorm:"column:content;"`
	Type            CourseInfoTypes `json:"type" gorm:"column:type;"`
	Course          Course          `json:"course" gorm:"foreignKey:CourseId;"`
}

func (CourseInfo) TableName() string { return "course_infos" }
