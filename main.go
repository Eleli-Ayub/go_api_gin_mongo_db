package main

import (
	"example.com/apis_db/database"
	"example.com/apis_db/students"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to database
	client, ctx, cancel, err := database.ConnectToDB("mongodb+srv://lesosioayub:ED0zc4tNdxQGy25c@golang.brga42o.mongodb.net/?retryWrites=true&w=majority")

	// initialize gin for api calling
	router := gin.Default()

	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)
	defer database.PingDB(client, ctx)

	students.SetupStudentRoutes(router)

	router.Run("localhost:8080")
}
