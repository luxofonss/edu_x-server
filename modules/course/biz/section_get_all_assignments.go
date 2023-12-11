package coursebiz

import (
	"context"

	assignmentmodel "server/modules/assignment/model"
)

type GetAssignmentSectionRepo interface {
	GetAssignmentBySectionId(ctx context.Context, id int) ([]*assignmentmodel.Assignment, error)
}

type getAssignmentsInSectionBiz struct {
	assignmentRepo GetAssignmentSectionRepo
}

func NewGetAssignmentsInSectionBiz(assignmentRepo GetAssignmentSectionRepo) *getAssignmentsInSectionBiz {
	return &getAssignmentsInSectionBiz{assignmentRepo: assignmentRepo}
}

func (biz *getAssignmentsInSectionBiz) GetAssignmentsInSection(ctx context.Context, sectionId int) ([]*assignmentmodel.Assignment, error) {
	assignments, err := biz.assignmentRepo.GetAssignmentBySectionId(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
