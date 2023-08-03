package app

import (
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/app/dependence"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	handlers := dependence.NewWire()
	configureMappings(router, handlers)
	return router
}

func configureMappings(router *gin.Engine, handlers dependence.HandlerContainer) {
	// Student
	apiGroupStudent := router.Group("v1.0/student")
	apiGroupStudent.GET("/profile", handlers.ViewProfileHandler.Handler)

	// Courses
	apiGroupCourses := router.Group("v1.0/courses")
	apiGroupCourses.GET("/course/:course_id", handlers.ViewCourseHandler.Handler)
	apiGroupCourses.GET("/list/", handlers.ListCoursesHandler.Handler)
}
