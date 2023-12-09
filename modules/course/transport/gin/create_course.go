package gincourse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursemodel "server/modules/course/model"
	coursepg "server/modules/course/repository/postgresql"
)

func CreateCourse(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId := c.MustGet(common.CurrentUser).(common.Requester)

		var data *coursemodel.CourseCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewCreateCourseBiz(courseRepo)

		createdCourse, err := biz.CreateCourse(c.Request.Context(), data, ownerId.GetUserId())
		if err != nil {
			panic(err)
		}

		createdCourse.Mask(true)
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(createdCourse))
	}
}
