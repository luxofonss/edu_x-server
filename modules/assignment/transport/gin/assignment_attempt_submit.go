package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentrepo "server/modules/assignment/repository"
)

func SubmitAssignment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var assignmentAttemptId = uuid.MustParse(c.Param("assignmentAttemptId"))
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptSubmitBiz(assignmentRepo)

		response, err := biz.UpdateAssignmentAttempt(c.Request.Context(), assignmentAttemptId, requester.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(response))
	}
}
