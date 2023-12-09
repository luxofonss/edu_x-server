package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/common"
	"server/libs/appctx"
	userpostgres "server/modules/user/repository/postgresql"
)

func RequiredTeacher(ctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := ctx.GetMainSQLDbConnection()
		userRepo := userpostgres.NewUserRepo(db)

		teacherInfo, err := userRepo.FindTeacherInfoByUserId(c.Request.Context(), requester.GetUserId())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				panic(common.NewCustomError(err, "Only teacher can create course!", "CANNOT_CREATE_COURSE"))
			} else {
				panic(err)
			}
		}

		if teacherInfo.DeletedAt != nil {
			panic(common.ErrNoPermission(nil))
		}

		c.Next()
	}
}
