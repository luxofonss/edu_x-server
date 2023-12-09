package coursemodel

import "server/common"

const LectureEntityName = "Lecture"

type Lecture struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
	Background      string `json:"background" gorm:"column:background;"`
	VideoUrl        string `json:"video_url" gorm:"column:video_url;"`
	SectionId       int    `json:"section_id" gorm:"column:section_id;"`
}

func (Lecture) TableName() string { return "lectures" }

func (l *Lecture) Mask(isAdminOrOwner bool) {
	l.GenUID(common.DbTypeLecture)
}
