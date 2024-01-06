package assignmentgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	assignmentbiz "server/modules/assignment/biz"
)

func RecognizeAssignment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//data := make(map[string]interface{})
		//if err := c.ShouldBindJSON(&data); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		data := map[string]interface{}{
			"assignment_id": 1,
		}

		biz := assignmentbiz.NewRecognizeAssignmentBiz(appCtx.GetRecognizeAssignmentProvider())
		result, err := biz.RecognizeAssignment(c.Request.Context(), data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
