package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

const SectionEntityName = "Section"

type Section struct {
	common.SQLModel `json:",inline"`
	Name            string                        `json:"name" gorm:"column:name;"`
	Description     string                        `json:"description" gorm:"column:description;"`
	CourseId        uuid.UUID                     `json:"course_id" gorm:"column:course_id;type:uuid;"`
	Course          Course                        `json:"course" gorm:"foreignKey:CourseId;"`
	Assignment      []*assignmentmodel.Assignment `json:"assignments" gorm:"many2many:assignment_placement;"`
	Lecture         []*Lecture                    `json:"lectures" gorm:"foreignKey:SectionId;"`
}

func (Section) TableName() string { return "sections" }
