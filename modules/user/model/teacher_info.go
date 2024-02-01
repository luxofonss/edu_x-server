package usrmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const TeacherInfoEntityName = "TeacherInfo"

type TeacherInfo struct {
	common.SQLModel  `json:",inline"`
	UserId           uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;" validate:"required"`
	EduQualification string    `json:"edu_qualification" gorm:"column:edu_qualification;" validate:"required"`
	Biography        string    `json:"biography" gorm:"column:biography;" validate:"required"`
	User             *User     `json:"user" gorm:"foreignKey:UserId;"`
}

func (TeacherInfo) TableName() string {
	return "teacher_infos"
}
