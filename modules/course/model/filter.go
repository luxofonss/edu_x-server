package coursemodel

import "github.com/google/uuid"

type Filter struct {
	TeacherId uuid.UUID `json:"teacher_id,omitempty" form:"teacher_id"`
}
