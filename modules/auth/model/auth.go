package authmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const EntityName = "Auth"

type AuthType string

const (
	EMAIL    AuthType = "email_pwd"
	FACEBOOK AuthType = "facebook"
	GOOGLE   AuthType = "google"
)

type Auth struct {
	common.SimpleSqlModel `json:",inline"`
	UserId                uuid.UUID `json:"user_id" gorm:"column:user_id;primary_key;type:uuid;"`
	AuthType              AuthType  `json:"auth_type" gorm:"column:auth_type;" validate:"required"`
	ServiceId             string    `json:"service_id" gorm:"column:service_id;"`
	Email                 string    `json:"email" gorm:"column:email;" validate:"required, email"`
	Password              string    `json:"-" gorm:"column:password;" validate:"required, min=6,max=32"`
	Salt                  string    `json:"-" gorm:"column:salt;" validate:"required"`
}

func (Auth) TableName() string { return "auths" }

type AuthLogin struct {
	Username string `json:"username" gorm:"column:username;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (AuthLogin) TableName() string {
	return Auth{}.TableName()
}

type RegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6,max=32"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Dob         string `json:"dob" binding:"required"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(nil, "email or password invalid", "ErrEmailOrPasswordInvalid")
	ErrEmailExisted           = common.NewCustomError(nil, "email has already existed", "ErrEmailExisted")
)
