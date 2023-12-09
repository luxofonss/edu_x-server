package userpostgres

import (
	"context"

	usrmodel "server/modules/user/model"
)

func (repo *userRepo) GetProfile(ctx context.Context, userId int) (*usrmodel.User, error) {
	// not implemented yet
	return nil, nil
}
