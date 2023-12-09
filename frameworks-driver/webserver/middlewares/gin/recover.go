package ginrecover

import (
	"github.com/gin-gonic/gin"
	"server/common"
	"server/libs/appctx"
)

func Recover(ctx appctx.AppContext) gin.HandlerFunc {
	print("recover middleware")
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				print("err in recover:: ", err)

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appError := common.ErrInternal(err.(error))

				c.AbortWithStatusJSON(appError.StatusCode, appError)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
