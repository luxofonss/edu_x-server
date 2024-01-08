package common

import (
	"github.com/google/uuid"
)

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

const DateString = "2006-01-02 15:04:05.999999 -0700 MST"
