package gincourse

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursemodel "server/modules/course/model"
	coursepg "server/modules/course/repository/postgresql"
)

func CreateLecture(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainSQLDbConnection()

		ownerId := c.MustGet(common.CurrentUser).(common.Requester)
		courseId := c.Param("courseId")
		courseIdInt, err := strconv.Atoi(courseId)
		sectionId := c.Param("sectionId")
		sectionIdInt, err := strconv.Atoi(sectionId)
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

		createdLecture, err := biz.CreateLecture(c.Request.Context(), data, courseIdInt, sectionIdInt, ownerId.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(createdLecture))
	}
}
