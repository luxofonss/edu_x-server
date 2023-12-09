package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;index"`
	FakeId    *UID       `json:"id" gorm:"-"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)

	m.FakeId = &uid
}

type SimpleSqlModel struct {
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
