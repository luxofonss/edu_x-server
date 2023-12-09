package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	"server/libs/token-provider/jwt"
	userpostgres "server/modules/user/repository/postgresql"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authentication header",
		"ErrWrongAuthHeader")
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(ctx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(ctx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := ctx.GetMainSQLDbConnection()
		repo := userpostgres.NewUserRepo(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := repo.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if user.DeletedAt != nil {
			panic(common.ErrNoPermission(nil))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)

		c.Next()
	}
}
