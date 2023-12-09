package coursemodel

import (
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
	SectionId       int                           `json:"section_id" gorm:"column:section_id;"`
	Assignment      []*assignmentmodel.Assignment `json:"assignments" gorm:"many2many:assignment_placement;"`
}

func (Lecture) TableName() string { return "lectures" }

func (l *Lecture) Mask(isAdminOrOwner bool) {
	l.GenUID(common.DbTypeLecture)
}
