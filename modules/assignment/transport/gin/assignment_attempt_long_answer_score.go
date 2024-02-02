package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentdto "server/modules/assignment/dto"
	assignmentrepo "server/modules/assignment/repository"
)

func LongAnswerScore(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var LongAnswerScore assignmentdto.LongAnswerScoreRequest

		if err := c.ShouldBind(&LongAnswerScore); err != nil {
			panic(err)
		}

		var assignmentAttemptId = uuid.MustParse(c.Param("assignmentAttemptId"))
		var questionAnswerId = uuid.MustParse(c.Param("questionAnswerId"))

		db := appCtx.GetMainSQLDbConnection()
		repo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentAttemptLongAnswerScoreBiz(repo)

		err := biz.ScoreAssignmentAttemptLongAnswer(c.Request.Context(), assignmentAttemptId, questionAnswerId, LongAnswerScore.Point)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": "Scored long answer successfully",
		})

	}
}
