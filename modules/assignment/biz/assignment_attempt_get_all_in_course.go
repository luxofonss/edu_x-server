package assignmentbiz

import (
	"context"
	"github.com/google/uuid"
	assignmentmodel "server/modules/assignment/model"
)

type AssignmentAttemptGetAllInCourseRepo interface {
	GetAllAssignmentAttemptInCourse(ctx context.Context, userId uuid.UUID, courseId uuid.UUID) ([]*assignmentmodel.AssignmentAttempt, error)
}

type AssignmentAttemptGetAllInCourseBiz struct {
	repo AssignmentAttemptGetAllInCourseRepo
}

func NewAssignmentAttemptGetAllInCourseBiz(repo AssignmentAttemptGetAllInCourseRepo) *AssignmentAttemptGetAllInCourseBiz {
	return &AssignmentAttemptGetAllInCourseBiz{repo: repo}
}

func (biz *AssignmentAttemptGetAllInCourseBiz) AssignmentAttemptGetAllInCourse(ctx context.Context, userId uuid.UUID, courseId uuid.UUID) ([]*assignmentmodel.AssignmentAttempt, error) {
	return biz.repo.GetAllAssignmentAttemptInCourse(ctx, userId, courseId)
}
