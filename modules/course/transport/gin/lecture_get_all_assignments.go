package gincourse

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	assignmentrepo "server/modules/assignment/repository"
	coursebiz "server/modules/course/biz"
)

func GetAllAssignmentsInLecture(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		lectureIdStr := c.Param("lectureId")
		lectureId, err := strconv.Atoi(lectureIdStr)
		if err != nil {
			panic(err)
		}

		fmt.Println(lectureId)
		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := coursebiz.NewGetAssignmentsInLectureBiz(assignmentRepo)

		assignments, err := biz.GetAssignmentsInLecture(c.Request.Context(), lectureId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignments))
	}
}
