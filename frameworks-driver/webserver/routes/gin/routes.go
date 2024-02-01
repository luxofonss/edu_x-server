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

	// Subject apis
	r.GET("/subjects", gincourse.GetAllSubject(ctx))

	// Courses
	r.GET("/courses/active", gincourse.GetAllActiveCourse(ctx))
	r.GET("/courses/:courseId", gincourse.GetCourseById(ctx))

	// REQUIRED AUTH APIS
	r.Use(middleware.RequiredAuth(ctx))
	r.GET("/user/profile", ginuser.GetProfile(ctx))

	// Upload files
	r.POST("/upload", middleware.RequiredAuth(ctx), uploadgin.Upload(ctx))

	// Course student management
	r.POST("/courses/attempt", gincourse.UserEnrollCourse(ctx))

	// Course
	r.POST("/courses", middleware.RequiredTeacher(ctx), gincourse.CreateCourse(ctx))
	r.POST("/courses/:courseId/section", middleware.RequiredTeacher(ctx), gincourse.CreateSection(ctx))
	r.POST("/courses/:courseId/section/:sectionId/lecture", middleware.RequiredTeacher(ctx), gincourse.CreateLecture(ctx))

	r.GET("/courses/mine", middleware.RequiredTeacher(ctx), gincourse.GetAllCourses(ctx))
	r.GET("/courses/:courseId/enrollments", middleware.RequiredTeacher(ctx), gincourse.GetAllCourseEnrollments(ctx))

	r.GET("/courses/:courseId/assignments", gincourse.GetAllAssignments(ctx))

	r.GET("/courses/:courseId/sections", gincourse.GetCourseSectionLecture(ctx))
	r.GET("/courses/:courseId/section/:sectionId/assignments", gincourse.GetAllAssignmentsInSection(ctx))
	r.GET("/courses/:courseId/section/:sectionId/lecture/:lectureId/assignments", gincourse.GetAllAssignmentsInLecture(ctx))

	r.GET("/courses/my-enrolled", gincourse.GetAllMyEnrolledCourses(ctx))

	// Course enroll manage
	r.PUT("/courses/enrolls", gincourse.UpdateCourseEnroll(ctx))

	// Assigment
	r.GET("/assignment/:id", middleware.RequiredAuth(ctx), assignmentgin.GetAssignment(ctx))
	r.POST("/assignment", middleware.RequiredTeacher(ctx), assignmentgin.CreateAssignment(ctx))

	// recognize assignment
	r.POST("/assignment/recognize", assignmentgin.RecognizeAssignment(ctx))

	r.GET("/teacher/assignment/attempt/get-all-attempts", middleware.RequiredAuth(ctx), assignmentgin.GetAllAssignmentAttemptByAssignmentId(ctx))
	r.GET("/assignment/attempt/:id", middleware.RequiredAuth(ctx), assignmentgin.GetAssignmentAttemptById(ctx))
	r.GET("/assignment/attempt-result/:id", middleware.RequiredAuth(ctx), assignmentgin.GetAssignmentAttemptResultById(ctx))
	// attempt an assignment in any course/lecture/section or in any public assignment
	r.POST("/assignment/attempt", middleware.RequiredAuth(ctx), assignmentgin.AttemptAssignment(ctx))

	// do assignment
	r.POST("/assignment-attempt/:assignmentAttemptId/question/:questionId/answer", middleware.RequiredAuth(ctx), assignmentgin.SubmitQuestionAnswer(ctx))

	// feedback
	r.POST("/assignment-attempt/:assignmentAttemptId/answer/:questionAnswerId/feedback", middleware.RequiredAuth(ctx), assignmentgin.CreateFeedbackAnswer(ctx))
	r.GET("/answer/:answerId/feedbacks", middleware.RequiredAuth(ctx), assignmentgin.GetFeedbacksByAnswerId(ctx))
}
