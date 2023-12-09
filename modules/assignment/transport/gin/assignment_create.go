package assignmentgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
	assignmentmodel "server/modules/assignment/model"
	assignmentrepo "server/modules/assignment/repository"
)

func CreateAssignment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId := c.MustGet(common.CurrentUser).(common.Requester)

		var data *assignmentmodel.AssignmentCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// setup dependencies
		db := appCtx.GetMainSQLDbConnection()
		assignmentRepo := assignmentrepo.NewAssignmentRepo(db)
		biz := assignmentbiz.NewAssignmentCreateBiz(assignmentRepo)

		assignmentCreated, err := biz.CreateAssignment(c.Request.Context(), data, ownerId.GetUserId())
		if err != nil {
			panic(err)
		}

		assignmentCreated.Mask(true)
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(assignmentCreated))
	}
}
