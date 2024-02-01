package assignmentgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentrepo "server/modules/assignment/repository"
)

func GetAllAssignmentAttemptByAssignmentId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("c.Query(\"assignment_id\"):: ", c.Query("assignment_id"))
		assignmentId, err := uuid.Parse(c.Query("assignment_id"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptTeacherGetAllBiz(assignmentRepo)

		assignmentAttempts, err := biz.GetAllAssignmentAttemptByAssignmentId(c.Request.Context(), assignmentId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignmentAttempts))
	}
}
