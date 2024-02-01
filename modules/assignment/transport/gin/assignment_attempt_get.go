package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentdto "server/modules/assignment/dto"
	assignmentrepo "server/modules/assignment/repository"
)

func GetAssignmentAttemptById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := uuid.Parse(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptGetBiz(assignmentRepo)

		assignmentAttempt, err := biz.GetAssignmentAttempt(c.Request.Context(), uid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignmentdto.ToAssignmentAttemptResponse(*assignmentAttempt)))
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignmentAttempt))
	}
}
