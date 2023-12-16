package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
)

const SubjectEntityName = "Subject"

type Subject struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
	ThumbNailUrl    string `json:"thumbnail_url" gorm:"column:thumbnail_url;"`
}

type SimpleSubjectGet struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;type:uuid;primary_key;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (Subject) TableName() string { return "subjects" }

func (SimpleSubjectGet) TableName() string { return "subjects" }
