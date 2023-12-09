package assignmentrepo

import (
	"gorm.io/gorm"
)

type assignmentRepo struct {
	db *gorm.DB
}

func NewAssignmentRepo(db *gorm.DB) *assignmentRepo {
	return &assignmentRepo{db: db}
}
