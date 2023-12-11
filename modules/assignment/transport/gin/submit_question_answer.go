package assignmentgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

		assignmentAttemptId, err := strconv.Atoi(c.Param("assignmentAttemptId"))
		if err != nil {
			panic(err)
		}

		if assignmentAttemptId == 0 {
			panic(common.NewCustomError(nil, "Invalid assignment attempt id", "SubmitQuestionAnswer"))
		}

		questionId, err := strconv.Atoi(c.Param("questionId"))
		if err != nil {
			panic(err)
		}
		if questionId == 0 {
			panic(common.NewCustomError(nil, "Invalid question id", "SubmitQuestionAnswer"))
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
