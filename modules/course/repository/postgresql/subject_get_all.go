package coursepg

import (
	coursemodel "server/modules/course/model"
)

func (repo *courseRepo) GetAllSubject() ([]*coursemodel.Subject, error) {
	var subjects []*coursemodel.Subject
	db := repo.db

	if err := db.Find(&subjects).Error; err != nil {
		return nil, err
	}

	return subjects, nil
}
