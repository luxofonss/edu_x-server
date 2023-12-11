package gincourse

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	assignmentrepo "server/modules/assignment/repository"
	coursebiz "server/modules/course/biz"
)

func GetAllAssignmentsInSection(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionIdStr := c.Param("sectionId")
		sectionId, err := strconv.Atoi(sectionIdStr)
		if err != nil {
			panic(err)
		}

		fmt.Println(sectionId)
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
