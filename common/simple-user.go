package common

type SimpleUser struct {
	SQLModel    `json:",inline"`
	LastName    string `json:"last_name" gorm:"column:last_name;"`
	FirstName   string `json:"first_name" gorm:"column:first_name;"`
	DOB         string `json:"dob" gorm:"column:dob;"`
	Gender      string `json:"gender" gorm:"column:gender;"`
	Email       string `json:"email" gorm:"column:email;"`
	Username    string `json:"username" gorm:"column:username;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
	Role        string `json:"role" gorm:"column:role;"`
}

func (SimpleUser) TableName() string {
	return "users"
}
