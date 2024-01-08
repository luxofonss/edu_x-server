package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
)

const UserEnrollCourseEntityName = "UserEnrollCourse"

type EnrollStatus string

const (
	PENDING  EnrollStatus = "PENDING"
	ACTIVE   EnrollStatus = "ACTIVE"
	INACTIVE EnrollStatus = "INACTIVE"
)

type UserEnrollCourse struct {
	common.SimpleSqlModel `json:",inline"`
	UserId                uuid.UUID     `json:"user_id" gorm:"primaryKey;column:user_id;"`
	CourseId              uuid.UUID     `json:"course_id" gorm:"primaryKey;column:course_id;"`
	Price                 float64       `json:"price" gorm:"column:price;"`
	StudentId             *int          `json:"student_id" gorm:"column:student_id;"`
	Status                *EnrollStatus `json:"status" gorm:"column:status;default:PENDING;"`
	Course                *Course       `json:"course" gorm:"foreignKey:CourseId;"`
}

func (UserEnrollCourse) TableName() string { return "user_enroll_course" }
