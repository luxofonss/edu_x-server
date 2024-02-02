package assignmentgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentmodel "server/modules/assignment/model"
	assignmentrepo "server/modules/assignment/repository"
)

func FeedbackEditQuestionAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *assignmentmodel.QuestionAnswer
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewCustomError(err, "Invalid payload", "SubmitQuestionAnswer"))
		}

		fmt.Println("data:: ", data.Id)

		questionId, err := uuid.Parse(c.Param("questionId"))
		if err != nil {
			panic(err)
		}

		assignmentAttemptId, err := uuid.Parse(c.Param("assignmentAttemptId"))
		if err != nil {
			panic(err)
		}

		data.QuestionId = questionId
		data.AssignmentAttemptId = assignmentAttemptId

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewFeedbackEditQuestionAnswer(assignmentRepo)

		questionAnswer, err := biz.FeedbackEditQuestionAnswer(c.Request.Context(), data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(questionAnswer))
	}
}
