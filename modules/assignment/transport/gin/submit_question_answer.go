package assignmentgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentmodel "server/modules/assignment/model"
	assignmentrepo "server/modules/assignment/repository"
)

func SubmitQuestionAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data *assignmentmodel.QuestionAnswer
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewCustomError(err, "Invalid payload", "SubmitQuestionAnswer"))
		}

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
		data.UserId = requester.GetUserId()

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewSubmitQuestionAnswerBiz(assignmentRepo)

		questionAnswer, err := biz.SubmitQuestionAnswer(c.Request.Context(), data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(questionAnswer))
	}
}
