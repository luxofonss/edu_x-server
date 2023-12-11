package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	usrbiz "server/modules/user/biz"
	usrrepo "server/modules/user/repository/postgresql"
)

func GetProfile(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		db := ctx.GetMainSQLDbConnection()
		repo := usrrepo.NewUserRepo(db)
		biz := usrbiz.NewProfileBiz(repo)

		user, err := biz.GetProfile(c.Request.Context(), u.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
