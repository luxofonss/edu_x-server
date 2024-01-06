package usrmodel

import (
	"github.com/google/uuid"
	"server/common"
	authmodel "server/modules/auth/model"
	coursemodel "server/modules/course/model"
)

const UserEntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string                `json:"email" gorm:"column:email;" validate:"required, email"`
	Username        string                `json:"username" gorm:"column:username;" validate:"required"`
	FirstName       string                `json:"first_name" gorm:"column:first_name;" validate:"required"`
	LastName        string                `json:"last_name" gorm:"column:last_name;" validate:"required"`
	PhoneNumber     string                `json:"phone_number" gorm:"column:phone_number;" validate:"required"`
	Gender          string                `json:"gender" gorm:"column:gender" validate:"required, eq=male|eq=female|eq=other"`
	Role            string                `json:"role" gorm:"column:role;"`
	Avatar          string                `json:"avatar" gorm:"column:avatar;type:json"`
	Dob             string                `json:"dob" gorm:"column:dob;type:timestamp;"`
	Verified        bool                  `json:"verified" gorm:"column:verified;type:boolean;default:false;"`
	Password        string                `json:"password" gorm:"-;" validate:"required, min=6,max=32"`
	Address         *Address              `json:"address" gorm:"foreignKey:UserId;"`
	LearnerInfo     *LearnerInfo          `json:"learner_info" gorm:"foreignKey:UserId;"`
	TeacherInfo     *TeacherInfo          `json:"teacher_info" gorm:"foreignKey:UserId;references:Id"`
	Auth            *authmodel.Auth       `json:"auth" gorm:"foreignKey:UserId"`
	Courses         []*coursemodel.Course `json:"courses" gorm:"many2many:user_enroll_course;"`
}

func (User) TableName() string { return "users" }

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetUserId() uuid.UUID {
	return u.Id
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(nil, "email or password invalid", "ErrEmailOrPasswordInvalid")
	ErrEmailExisted           = common.NewCustomError(nil, "email has already existed", "ErrEmailExisted")
)
