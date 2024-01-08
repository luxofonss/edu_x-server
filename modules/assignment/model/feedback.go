package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const FeedbackEntityName = "Feedbacks"

type Feedback struct {
	common.SQLModel  `json:",inline"`
	UserId           uuid.UUID   `json:"user_id" gorm:"column:user_id;type:uuid;"`
	Message          string      `json:"message" gorm:"column:message;"`
	QuestionAnswerId uuid.UUID   `json:"question_answer_id" gorm:"column:question_answer_id;type:uuid;"`
	FeedbackId       *uuid.UUID  `json:"feedback_id" gorm:"column:feedback_id;type:uuid;"`
	Feedbacks        []*Feedback `json:"feedbacks" gorm:"foreignKey:FeedbackId;"`
}
