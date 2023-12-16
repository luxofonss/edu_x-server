package coursebiz

import (
	"context"
	coursemodel "server/modules/course/model"
)

type GetAllSubjectRepo interface {
	GetAllSubject() ([]*coursemodel.Subject, error)
}

type getAllSubjectBiz struct {
	repo GetAllSubjectRepo
}

func NewGetAllSubjectBiz(repo GetAllSubjectRepo) *getAllSubjectBiz {
	return &getAllSubjectBiz{repo: repo}
}

func (biz *getAllSubjectBiz) GetAllSubject(ctx context.Context) ([]*coursemodel.Subject, error) {
	subjects, err := biz.repo.GetAllSubject()
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
