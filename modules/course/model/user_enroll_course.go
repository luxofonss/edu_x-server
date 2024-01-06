package coursemodel

import (
	"github.com/google/uuid"
	"server/common"
)

const UserEnrollCourseEntityName = "UserEnrollCourse"

type UserEnrollCourse struct {
	common.SimpleSqlModel `json:",inline"`
	UserId                uuid.UUID `json:"user_id" gorm:"primaryKey;column:user_id;"`
	CourseId              uuid.UUID `json:"course_id" gorm:"primaryKey;column:course_id;"`
	Price                 float64   `json:"price" gorm:"column:price;"`
	StudentId             *int      `json:"student_id" gorm:"column:student_id;"`
}

func (UserEnrollCourse) TableName() string { return "user_enroll_course" }
