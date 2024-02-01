package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	DOB       string `json:"dob" gorm:"column:dob;"`
	Gender    string `json:"gender" gorm:"column:gender;"`
}

func (SimpleUser) TableName() string {
	return "users"
}
