package uploadmodel

import "server/common"

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "this file is not an image", "ErrFileIsNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "cannot save file", "ErrCannotSaveFile")
}
