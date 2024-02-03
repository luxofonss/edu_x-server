package coursepg

import (
	"context"
	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) AddUsersToCourseByIds(ctx context.Context, courseId uuid.UUID, userIds []uuid.UUID) error {
	db := repo.db.Table(coursemodel.UserEnrollCourse{}.TableName())
	// Start a new transaction
	tx := db.Begin()

	// Check for errors during the transaction initiation
	if tx.Error != nil {
		return common.ErrDB(tx.Error)
	}

	// Perform your database operations within the transaction
	for _, userId := range userIds {
		err := tx.Create(&coursemodel.UserEnrollCourse{
			CourseId: courseId,
			UserId:   userId,
			Status:   coursemodel.ACTIVE,
			Price:    0,
		}).Error

		// Check for errors during the operation
		if err != nil {
			// If there is an error, rollback the transaction
			tx.Rollback()
			return common.ErrDB(err)
		}
	}

	// If everything is successful, commit the transaction
	tx.Commit()
	return nil
}
