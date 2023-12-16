package gincourse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursemodel "server/modules/course/model"
	coursepg "server/modules/course/repository/postgresql"
)

func CreateSection(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId := c.MustGet(common.CurrentUser).(common.Requester)

		courseId, err := uuid.Parse(c.Param("courseId"))
		if err != nil {
			panic(err)
		}

		var data *coursemodel.Section
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewCreateSectionBiz(courseRepo)

		createdSection, err := biz.CreateSection(c.Request.Context(), data, courseId, ownerId.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(createdSection))
	}
}
