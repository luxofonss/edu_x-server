package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentrepo "server/modules/assignment/repository"
)

func GetAllAssignmentAttemptInCourse(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		userId := requester.GetUserId()
		courseId := uuid.MustParse(c.Param("courseId"))

		db := appCtx.GetMainSQLDbConnection()
		repo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptGetAllInCourseBiz(repo)

		assignmentAttempts, err := biz.AssignmentAttemptGetAllInCourse(c.Request.Context(), userId, courseId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(assignmentAttempts))
	}
}
