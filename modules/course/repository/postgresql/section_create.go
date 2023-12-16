package coursepg

import (
	"context"

	"github.com/google/uuid"
	"server/common"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) CreateSection(ctx context.Context, data *coursemodel.Section, userId uuid.UUID) (*coursemodel.Section, error) {
	db := repo.db.Table(coursemodel.Section{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
