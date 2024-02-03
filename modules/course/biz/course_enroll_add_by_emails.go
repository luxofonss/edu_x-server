package coursebiz

import (
	"context"
	"github.com/google/uuid"
	usermodel "server/modules/user/model"
)

type AddUsersToCourseRepo interface {
	AddUsersToCourseByIds(ctx context.Context, courseId uuid.UUID, ids []uuid.UUID) error
}

type GetUsersByEmailsRepo interface {
	GetUsersByEmails(ctx context.Context, emails []string) ([]*usermodel.User, error)
}

type AddUsersToCourseByEmailsBiz struct {
	repo     AddUsersToCourseRepo
	userRepo GetUsersByEmailsRepo
}

func NewAddUsersToCourseByEmailsBiz(repo AddUsersToCourseRepo, userRepo GetUsersByEmailsRepo) *AddUsersToCourseByEmailsBiz {
	return &AddUsersToCourseByEmailsBiz{repo: repo, userRepo: userRepo}
}

func (biz *AddUsersToCourseByEmailsBiz) AddUsersToCourseByEmails(ctx context.Context, courseId uuid.UUID, emails []string) error {
	users, err := biz.userRepo.GetUsersByEmails(ctx, emails)
	if err != nil {
		return err
	}

	var userIds []uuid.UUID
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}

	return biz.repo.AddUsersToCourseByIds(ctx, courseId, userIds)
}
