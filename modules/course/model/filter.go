package coursemodel

import "github.com/google/uuid"

type Filter struct {
	Id        uuid.UUID `json:"id,omitempty" form:"id"`
	TeacherId uuid.UUID `json:"teacher_id,omitempty" form:"teacher_id"`
	Code      *string   `json:"code,omitempty" form:"code"`
	Status    string    `json:"status,omitempty" form:"status"`
}
