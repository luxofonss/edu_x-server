package postgresql

import "gorm.io/gorm"

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}
