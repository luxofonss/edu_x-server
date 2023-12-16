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

func CreateLecture(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainSQLDbConnection()

		courseId, err := uuid.Parse(c.Param("courseId"))
		if err != nil {
			panic(err)
		}

		sectionId, err := uuid.Parse(c.Param("sectionId"))
		if err != nil {
			panic(err)
		}

		ownerId := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(err)
		}
		var data *coursemodel.Lecture

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// setup dependencies
		courseRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewCreateLectureBiz(courseRepo)

		createdLecture, err := biz.CreateLecture(c.Request.Context(), data, courseId, sectionId, ownerId.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(createdLecture))
	}
}
