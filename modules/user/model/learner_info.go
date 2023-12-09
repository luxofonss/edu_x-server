package usrmodel

import "server/common"

const LearnerInfoEntityName = "LearnerInfo"

type LearnerType string

const (
	STUDENT LearnerType = "student"
	PUPIL   LearnerType = "pupil"
	OTHER   LearnerType = "other"
)

type LearnerInfo struct {
	common.SQLModel `json:",inline"`
	UserId          int         `json:"user_id" gorm:"column:user_id;" validate:"required"`
	Type            LearnerType `json:"type" gorm:"column:type;" validate:"required"`
	Grade           int         `json:"grade" gorm:"column:grade;"`
	School          string      `json:"school" gorm:"column:school;"`
}

func (LearnerInfo) TableName() string {
	return "learner_info"
}
