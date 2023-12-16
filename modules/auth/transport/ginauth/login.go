package ginauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	"server/libs/hasher"
	"server/libs/token-provider/jwt"
	authbiz "server/modules/auth/biz"
	authmodel "server/modules/auth/model"
	authrepo "server/modules/auth/repository/postgresql"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainSQLDbConnection()

		var data authmodel.AuthLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// setup dependencies
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		md5 := hasher.NewMD5Hash()

		loginRepo := authrepo.NewAuthRepo(db)

		biz := authbiz.NewLoginBiz(loginRepo, md5, tokenProvider, common.TokenExpireTime)

		token, err := biz.Login(c.Request.Context(), &data)
		c.SetCookie("token", token.Token, 3600, "/", "localhost", false, true)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
