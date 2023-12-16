package coursepg

import (
	"context"

	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetSectionCourseInfo(ctx context.Context, sectionId uuid.UUID) (*coursemodel.Section, error) {
	db := repo.db.Table(coursemodel.Section{}.TableName())

	var sectionCourseInfo *coursemodel.Section

	if err := db.Preload("Course").Where("id = ?", sectionId).First(&sectionCourseInfo).Error; err != nil {
		return nil, err
	}

	return sectionCourseInfo, nil
}
