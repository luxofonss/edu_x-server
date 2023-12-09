package coursepg

import "gorm.io/gorm"

type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *courseRepo {
	return &courseRepo{db: db}
}
