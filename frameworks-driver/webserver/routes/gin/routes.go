package ginroutes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/libs/appctx"
	"server/middleware"
	assignmentgin "server/modules/assignment/transport/gin"
	"server/modules/auth/transport/ginauth"
	gincourse "server/modules/course/transport/gin"
	uploadgin "server/modules/upload/transport/gin"
	"server/modules/user/transport/ginuser"
)

func SetupRoutes(ctx appctx.AppContext, r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth/register", ginauth.Register(ctx))
	r.POST("/auth/login", ginauth.Login(ctx))

	r.Use(middleware.RequiredAuth(ctx))
	r.GET("/user/profile", ginuser.GetProfile(ctx))

	// Upload files
	r.POST("/upload", middleware.RequiredAuth(ctx), uploadgin.Upload(ctx))

	// Course
	r.POST("/course", middleware.RequiredTeacher(ctx), gincourse.CreateCourse(ctx))
	r.POST("/course/:courseId/section", middleware.RequiredTeacher(ctx), gincourse.CreateSection(ctx))
	r.POST("/course/:courseId/section/:sectionId/lecture", middleware.RequiredTeacher(ctx), gincourse.CreateLecture(ctx))

	r.GET("/course/:courseId/assignments", gincourse.GetAllAssignments(ctx))
	// Assigment
	r.GET("/assignment/:id", middleware.RequiredAuth(ctx), assignmentgin.GetAssignment(ctx))
	r.POST("/assignment", middleware.RequiredTeacher(ctx), assignmentgin.CreateAssignment(ctx))
}
