package gincourse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/common"
	"server/libs/appctx"
	assignmentrepo "server/modules/assignment/repository"
	coursebiz "server/modules/course/biz"
)

func GetAllAssignmentsInSection(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionId, err := uuid.Parse(c.Param("sectionId"))
		if err != nil {
			panic(err)
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := coursebiz.NewGetAssignmentsInSectionBiz(assignmentRepo)

		assignments, err := biz.GetAssignmentsInSection(c.Request.Context(), sectionId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(assignments))
	}
}
