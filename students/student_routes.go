package students

import "github.com/gin-gonic/gin"

func SetupStudentRoutes(router *gin.Engine) {
	studentRoutes := router.Group("/students")
	{
		studentRoutes.POST("/addstudent", CreateStudent)
		studentRoutes.GET("/getstudents/", GetStudents)
		studentRoutes.DELETE("/deletestudent/id", DeleteStudent)
		studentRoutes.PATCH("/updatestudent/id", UpdatStudent)
	}
}
