package assignmentmodel

import (
	"github.com/google/uuid"
	"time"
)

const FeedbackEntityName = "Feedbacks"

type Feedback struct {
	Id               uuid.UUID   `json:"id" gorm:"primaryKey;unique;column:id;type:uuid;default:uuid_generate_v4()"`
	DeletedAt        *time.Time  `json:"deleted_at" gorm:"column:deleted_at;default:NULL"`
	CreatedAt        *time.Time  `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP()"`
	UpdatedAt        *time.Time  `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
	UserId           uuid.UUID   `json:"user_id" gorm:"column:user_id;type:uuid;"`
	Message          string      `json:"message" gorm:"column:message;"`
	Type             string      `json:"type" gorm:"column:type;"`
	QuestionAnswerId uuid.UUID   `json:"question_answer_id" gorm:"column:question_answer_id;type:uuid;"`
	FeedbackId       *uuid.UUID  `json:"feedback_id" gorm:"column:feedback_id;type:uuid;default:NULL;"`
	Feedbacks        []*Feedback `json:"feedbacks" gorm:"foreignKey:FeedbackId;"`
}
