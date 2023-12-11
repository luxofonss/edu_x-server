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

func GetAllAssignments(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseIdParam := c.Param("courseId")
		courseId, err := strconv.Atoi(courseIdParam)
		if err != nil {
			panic(err)
		}

		fmt.Println(courseId)
		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := coursebiz.NewGetAssignmentsInCourseBiz(assignmentRepo)

		assignments, err := biz.GetAssignmentsInCourse(c.Request.Context(), courseId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignments))
	}
}
