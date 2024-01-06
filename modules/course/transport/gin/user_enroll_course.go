package gincourse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursedto "server/modules/course/dto"
	coursepg "server/modules/course/repository/postgresql"
)

func UserEnrollCourse(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("user enroll course")
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data *coursedto.CourseAttemptRequest
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		courseRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewCourseAttemptByCodeBiz(courseRepo)

		res, err := biz.AttemptCourseByCode(c.Request.Context(), data.Code, requester.GetUserId())
		if err != nil {
			panic(err)
		}

		if res == false {
			panic(nil)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse("Attempt course successfully!"))
	}
}
