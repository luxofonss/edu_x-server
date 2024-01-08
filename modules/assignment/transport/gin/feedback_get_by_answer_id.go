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

func GetFeedbacksByAnswerId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		answerId, err := uuid.Parse(c.Param("answerId"))
		fmt.Println("answerId:: ", answerId)
		if err != nil {
			panic(err)
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		getFeedbacksByAnswerIdBiz := assignmentbiz.NewGetFeedbackByAnswerIdBiz(assignmentRepo)

		feedbacks, err := getFeedbacksByAnswerIdBiz.GetFeedbackByAnswerId(c.Request.Context(), answerId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(feedbacks))
	}
}
