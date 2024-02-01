package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentrepo "server/modules/assignment/repository"
)

func GetAssignmentAttemptResultById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := uuid.Parse(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptGetResultBiz(assignmentRepo)

		assignmentAttempt, err := biz.GetAssignmentAttemptResult(c.Request.Context(), uid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignmentAttempt))
	}
}
