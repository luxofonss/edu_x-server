package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentdto "server/modules/assignment/dto"
	assignmentmodel "server/modules/assignment/model"
	assignmentrepo "server/modules/assignment/repository"
)

func CreateFeedbackAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data *assignmentdto.FeedbackQuestionAnswerRequest
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewCustomError(err, "Invalid payload", "CreateFeedbackAnswer"))
		}

		questionAnswerId, err := uuid.Parse(c.Param("questionAnswerId"))
		if err != nil {
			panic(err)
		}

		feedbackId, err := uuid.Parse(data.FeedbackId)
		if err != nil {
			panic(err)
		}

		feedback := &assignmentmodel.Feedback{
			QuestionAnswerId: questionAnswerId,
			UserId:           requester.GetUserId(),
			Message:          data.Message,
			FeedbackId:       &feedbackId,
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewFeedbackAnswerBiz(assignmentRepo)

		feedbackAnswer, err := biz.CreateFeedbackAnswer(c.Request.Context(), feedback)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(feedbackAnswer))
	}
}
