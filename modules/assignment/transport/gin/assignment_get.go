package assignmentgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentrepo "server/modules/assignment/repository"
)

func GetAssignment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := uuid.Parse(c.Param("id"))
		if err != nil {
			panic(err)
		}

		if err != nil {
			panic(err)
		}
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentGetBiz(assignmentRepo)

		assignment, err := biz.GetAssignment(c.Request.Context(), uid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignment))
	}
}
