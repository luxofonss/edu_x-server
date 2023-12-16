package coursebiz

import (
	"context"

	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type GetAssignmentSectionRepo interface {
	GetAssignmentBySectionId(ctx context.Context, id uuid.UUID) ([]*assignmentmodel.Assignment, error)
}

type getAssignmentsInSectionBiz struct {
	assignmentRepo GetAssignmentSectionRepo
}

func NewGetAssignmentsInSectionBiz(assignmentRepo GetAssignmentSectionRepo) *getAssignmentsInSectionBiz {
	return &getAssignmentsInSectionBiz{assignmentRepo: assignmentRepo}
}

func (biz *getAssignmentsInSectionBiz) GetAssignmentsInSection(ctx context.Context, sectionId uuid.UUID) ([]*assignmentmodel.Assignment, error) {
	assignments, err := biz.assignmentRepo.GetAssignmentBySectionId(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
