package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
	assignmentmodel "server/modules/assignment/model"
)

const LectureEntityName = "Lecture"

type Lecture struct {
	common.SQLModel `json:",inline"`
	Name            string                        `json:"name" gorm:"column:name;"`
	Description     string                        `json:"description" gorm:"column:description;"`
	Background      string                        `json:"background" gorm:"column:background;"`
	VideoUrl        string                        `json:"video_url" gorm:"column:video_url;"`
	SectionId       uuid.UUID                     `json:"section_id" gorm:"column:section_id;type:uuid;"`
	Assignment      []*assignmentmodel.Assignment `json:"assignments" gorm:"many2many:assignment_placement;"`
}

func (Lecture) TableName() string { return "lectures" }
