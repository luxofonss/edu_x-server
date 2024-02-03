package gincourse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursedto "server/modules/course/dto"
	coursepg "server/modules/course/repository/postgresql"
	usrrepo "server/modules/user/repository/postgresql"
)

func AddUsersToCourseByEmails(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId := uuid.MustParse(c.Param("courseId"))

		var data coursedto.CourseEnrollAddByEmailsRequest

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMainSQLDbConnection()
		repo := coursepg.NewCourseRepo(db)
		userRepo := usrrepo.NewUserRepo(db)

		biz := coursebiz.NewAddUsersToCourseByEmailsBiz(repo, userRepo)

		err := biz.AddUsersToCourseByEmails(c.Request.Context(), courseId, data.Emails)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"message": "success"})
	}
}
