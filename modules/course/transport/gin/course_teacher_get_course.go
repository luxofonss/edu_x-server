package gincourse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursemodel "server/modules/course/model"
	coursepg "server/modules/course/repository/postgresql"
)

func GetAllCourses(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId := c.MustGet(common.CurrentUser).(common.Requester)

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		pagingData.Fulfill()

		var filter coursemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		filter.TeacherId = ownerId.GetUserId()
		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewGetCourseBiz(courseRepo)

		courses, err := biz.GetAllMyCourses(c.Request.Context(), &pagingData, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(courses))
	}
}
