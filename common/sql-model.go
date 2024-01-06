package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"primaryKey;unique;column:id;type:uuid;default:uuid_generate_v4()"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;default:NULL"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
}

type SimpleSqlModel struct {
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;default:NULL"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
}
