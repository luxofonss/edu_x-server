package postgresql

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"server/common"
	authmodel "server/modules/auth/model"
)

func (repo *authRepo) GetAuth(ctx context.Context, data *authmodel.AuthLogin) (*authmodel.Auth, error) {
	db := repo.db.Table(authmodel.AuthLogin{}.TableName())
	var auth authmodel.Auth
	if err := db.Where("email = ?", data.Email).First(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(authmodel.EntityName, err)
		}
		return nil, err
	}
	return &auth, nil
}
