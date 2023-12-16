package coursebiz

import (
	"context"

	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type GetAssignmentLectureRepo interface {
	GetAssignmentByLectureId(ctx context.Context, id uuid.UUID) ([]*assignmentmodel.Assignment, error)
}

type getAssignmentsInLectureBiz struct {
	assignmentRepo GetAssignmentLectureRepo
}

func NewGetAssignmentsInLectureBiz(assignmentRepo GetAssignmentLectureRepo) *getAssignmentsInLectureBiz {
	return &getAssignmentsInLectureBiz{assignmentRepo: assignmentRepo}
}

func (biz *getAssignmentsInLectureBiz) GetAssignmentsInLecture(ctx context.Context, lectureId uuid.UUID) ([]*assignmentmodel.Assignment, error) {
	assignments, err := biz.assignmentRepo.GetAssignmentByLectureId(ctx, lectureId)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
