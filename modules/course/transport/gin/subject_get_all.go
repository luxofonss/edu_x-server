package gincourse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
	"server/libs/appctx"
	coursebiz "server/modules/course/biz"
	coursepg "server/modules/course/repository/postgresql"
)

func GetAllSubject(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainSQLDbConnection()
		subjectRepo := coursepg.NewCourseRepo(db)
		biz := coursebiz.NewGetAllSubjectBiz(subjectRepo)

		subjects, err := biz.GetAllSubject(c.Request.Context())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(subjects))
	}
}
