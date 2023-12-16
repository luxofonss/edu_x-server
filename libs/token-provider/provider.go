package tokenprovider

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"server/common"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
}

type Provider interface {
	Generate(payload TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrorNotFound      = common.NewCustomError(errors.New("token not found"), "token not found", "ErrNotFound")
	ErrorEncodingToken = common.NewCustomError(errors.New("error encoding token"), "error encoding token", "ErrEncodingToken")
	ErrorInvalidToken  = common.NewCustomError(errors.New("invalid token"), "invalid token", "ErrInvalidToken")
)
