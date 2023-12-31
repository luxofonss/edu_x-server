package userpostgres

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	usermodel "server/modules/user/model"
)

func (repo *userRepo) FindTeacherInfoByUserId(ctx context.Context, userId uuid.UUID) (*usermodel.TeacherInfo, error) {
	db := repo.db.Table(usermodel.TeacherInfo{}.TableName())
	var teacherInfo usermodel.TeacherInfo
	if err := db.Where("user_id = ?", userId).First(&teacherInfo).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &teacherInfo, nil
}
