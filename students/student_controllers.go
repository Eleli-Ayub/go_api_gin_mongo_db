package students

import (
	"net/http"

	"example.com/apis_db/database"
	"example.com/apis_db/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStudent(c *gin.Context) {

	client, ctx, cancel, err := database.DbConnect()

	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)

	var studentDetails models.Student

	if err := c.ShouldBindJSON(&studentDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inserDocResult, err := CreateStudentDocumment(client, ctx, studentDetails)

	if err != nil {
		panic(err)
	}

	data := gin.H{

		"success": true,
		"message": "student created successfully",
		"student": inserDocResult,
	}
	c.JSON(http.StatusOK, data)
}
func GetStudents(c *gin.Context) {

	client, ctx, cancel, err := database.DbConnect()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)

	var query, field interface{}

	field = bson.D{{"_id", 0}}
	query = bson.D{{}}

	var results []bson.D
	cursor, err := QueryStudents(client, ctx, query, field)

	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	response := gin.H{
		"success": true,
		"message": "Data fetched successfully",
		"data":    results,
	}
	c.JSON(http.StatusOK, gin.H{"students fetch": response})
}
func GetSingleStudent(c *gin.Context) {

	client, ctx, cancel, err := database.DbConnect()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)

	var query, field interface{}

	field = bson.D{{"_id", 0}}
	query = bson.D{
		{"gpa", bson.D{{"$gte", 2.0}}},
	}

	result, err := QuerySingleStudent(client, ctx, query, field)

	if err != nil {
		panic(err)
	}

	response := gin.H{
		"success": true,
		"message": "Data fetched successfully",
		"data":    result,
	}
	c.JSON(http.StatusOK, gin.H{"student fetch": response})
}
func DeleteStudent(c *gin.Context) {

	client, ctx, cancel, err := database.DbConnect()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)

	var query interface{}

	query = bson.D{
		{"gender", "male"},
	}

	result, err := DeleteStudentQuery(client, ctx, query)

	if err != nil {
		panic(err)
	}

	response := gin.H{
		"success": true,
		"message": "Data Deleted successfully",
		"result":  result,
	}
	c.JSON(http.StatusOK, gin.H{"students delete": response})
}

func UpdatStudent(c *gin.Context) {

	client, ctx, cancel, err := database.DbConnect()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB(client, ctx, cancel)

	var studentDetails models.Student

	if err := c.ShouldBindJSON(&studentDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter := bson.D{
		{"name", "Chie Brenda"},
	}

	result, err := UpdateStudentQuery(client, ctx, filter, studentDetails)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"student update result": result})

}
