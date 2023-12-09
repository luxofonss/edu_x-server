package ginauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	"server/libs/hasher"
	usrbiz "server/modules/auth/biz"
	authrepo "server/modules/auth/repository/postgresql"
	usermodel "server/modules/user/model"
	userpostgres "server/modules/user/repository/postgresql"
)

//func Register(appCtx appctx.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		db := appCtx.GetMainSQLDbConnection()
//
//		var requestBody authmodel.RegisterRequest
//		if err := c.ShouldBind(&requestBody); err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//		// setup dependencies
//		md5 := hasher.NewMD5Hash()
//		authRepo := authrepo.NewAuthRepo(db)
//		userRepo := userpostgres.NewUserRepo(db)
//		biz := usrbiz.NewRegisterBiz(authRepo, userRepo, md5)
//
//		userData, err := biz.Register(c.Request.Context(), &requestBody)
//		if err != nil {
//			panic(common.ErrInternal(err))
//		}
//
//		userData.Mask(false)
//		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userData.FakeId.String()))
//	}
//}

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainSQLDbConnection()

		var requestBody usermodel.User
		if err := c.ShouldBind(&requestBody); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		// setup dependencies
		md5 := hasher.NewMD5Hash()
		authRepo := authrepo.NewAuthRepo(db)
		userRepo := userpostgres.NewUserRepo(db)
		biz := usrbiz.NewRegisterBiz(authRepo, userRepo, md5)

		userData, err := biz.Register(c.Request.Context(), &requestBody)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		userData.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userData.FakeId.String()))
	}
}
