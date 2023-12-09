package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Database error", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "Internal error", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err, "Cannot list "+entity, "CANNOT_LIST_"+entity)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err, "Cannot delete "+entity, "CANNOT_DELETE_"+entity)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err, "Cannot update "+entity, "CANNOT_UPDATE_"+entity)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err, "Cannot get "+entity, "CANNOT_GET_"+entity)
}

func ErrEntityDeleted(entity string) *AppError {
	return NewCustomError(nil, entity+" was deleted", entity+"_DELETED")
}

func ErrEntityExisted(entity string) *AppError {
	return NewCustomError(nil, entity+" was existed", entity+"_EXISTED")
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, entity+" not found", entity+"_NOT_FOUND")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, "Cannot create "+entity, "CANNOT_CREATE_"+entity)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, "You don't have permission", "NO_PERMISSION")
}

var RecordNotFound = errors.New("record not found")
