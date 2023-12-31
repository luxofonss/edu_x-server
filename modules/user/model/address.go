package usrmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const AddressEntityName = "Address"

type Address struct {
	common.SQLModel `json:",inline"`
	UserId          uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;"`
	Country         string    `json:"country" gorm:"column:country;"`
	Province        string    `json:"province" gorm:"column:province;"`
	District        string    `json:"district" gorm:"column:district;"`
	Ward            string    `json:"ward" gorm:"column:ward;"`
	HouseNumber     string    `json:"house_number" gorm:"column:house_number;"`
}

func (Address) TableName() string { return "addresses" }
