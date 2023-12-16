package common

import "github.com/google/uuid"

const (
	DbTypeCourse            = 1
	DbTypeUser              = 2
	DbTypeSection           = 3
	DbTypeLecture           = 4
	DbTypeAssignment        = 5
	DbTypeQuestion          = 6
	DbTypeQuestionChoice    = 7
	DbTypeQuestionAnswer    = 8
	DbTypeAssignmentAttempt = 9
	TokenExpireTime         = 60 * 60 // 1 hour
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() uuid.UUID
	GetEmail() string
}
