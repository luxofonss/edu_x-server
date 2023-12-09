package usrmodel

import "server/common"

const TeacherInfoEntityName = "TeacherInfo"

type TeacherInfo struct {
	common.SQLModel  `json:",inline"`
	UserId           int    `json:"user_id" gorm:"column:user_id;" validate:"required"`
	EduQualification int    `json:"edu_qualification" gorm:"column:edu_qualification;" validate:"required"`
	Biography        string `json:"biography" gorm:"column:biography;" validate:"required"`
}

func (TeacherInfo) TableName() string {
	return "teacher_infos"
}
