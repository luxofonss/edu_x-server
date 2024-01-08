package gincourse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursedto "server/modules/course/dto"
	coursepg "server/modules/course/repository/postgresql"
)

func GetAllActiveCourse(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		getAllActiveCourseBiz := coursebiz.NewGetAllActiveCourseBiz(courseRepo)

		courses, err := getAllActiveCourseBiz.GetAllActiveCourse(c.Request.Context())
		if err != nil {
			panic(err)
		}

		var response []coursedto.SimpleCourseResponse
		for i := range courses {
			response = append(response, coursedto.ToSimpleCourseResponse(*courses[i]))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
