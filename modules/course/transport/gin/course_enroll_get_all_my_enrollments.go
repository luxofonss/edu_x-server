package gincourse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursepg "server/modules/course/repository/postgresql"
)

func GetAllMyEnrolledCourses(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		getAllMyEnrolledCoursesBiz := coursebiz.NewGetAllMyEnrollmentsCourseBiz(courseRepo)

		res, err := getAllMyEnrolledCoursesBiz.GetAllMyEnrollments(c.Request.Context(), requester.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
