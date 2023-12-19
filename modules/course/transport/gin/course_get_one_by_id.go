package gincourse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursepg "server/modules/course/repository/postgresql"
)

func GetCourseById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId, err := uuid.Parse(c.Param("courseId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		getOneCourseBiz := coursebiz.NewGetOneCourseBiz(courseRepo)

		res, err := getOneCourseBiz.GetOneCourseById(c.Request.Context(), courseId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
