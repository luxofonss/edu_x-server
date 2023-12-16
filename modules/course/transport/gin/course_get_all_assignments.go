package gincourse

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentrepo "server/modules/assignment/repository"
	coursebiz "server/modules/course/biz"
)

func GetAllAssignments(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId, err := uuid.Parse(c.Param("courseId"))
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
