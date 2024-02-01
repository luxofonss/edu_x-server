package gincourse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursedto "server/modules/course/dto"
	coursepg "server/modules/course/repository/postgresql"
)

func UpdateCourseEnroll(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var courseEnrollUpdateRequest coursedto.CourseEnrollUpdateRequest
		if err := c.ShouldBindJSON(&courseEnrollUpdateRequest); err != nil {
			panic(err)
		}

		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		updateCourseEnrollBiz := coursebiz.NewUpdateCourseEnrollStatusBiz(courseRepo)

		courseEnrollId, err := uuid.Parse(courseEnrollUpdateRequest.CourseEnrollId)
		if err != nil {
			panic(err)
		}

		err = updateCourseEnrollBiz.UpdateCourseEnrollStatus(courseEnrollId, courseEnrollUpdateRequest.Status)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("OK"))
	}
}
