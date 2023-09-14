package students

import "github.com/gin-gonic/gin"

func SetupStudentRoutes(router *gin.Engine) {
	studentRoutes := router.Group("/students")
	{
		studentRoutes.POST("/addstudent", CreateStudent)
		studentRoutes.GET("/getstudents", GetStudents)
		studentRoutes.GET("/getsinglestudent", GetSingleStudent)
		studentRoutes.GET("/deletestudent", DeleteStudent)
		studentRoutes.GET("/updatestudent", UpdatStudent)
	}
}
