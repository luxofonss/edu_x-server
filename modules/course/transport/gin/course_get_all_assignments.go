package gincourse

import (
	"github.com/gin-gonic/gin"
	"server/libs/appctx"
)

func GetAllAssignments(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//ownerId := c.MustGet(common.CurrentUser).(common.Requester)

	}
}
