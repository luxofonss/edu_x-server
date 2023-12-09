package coursemodel

import "server/common"

const SectionEntityName = "Section"

type Section struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
	CourseId        int    `json:"course_id" gorm:"column:course_id;"`
	Course          Course `json:"course" gorm:"foreignKey:CourseId;"`
}

func (Section) TableName() string { return "sections" }

func (s *Section) Mask(isAdminOrOwner bool) {
	s.GenUID(common.DbTypeSection)
}
