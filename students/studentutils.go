package students

import (
	"context"

	"example.com/apis_db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateStudentDocumment(client *mongo.Client, ctx context.Context, studentDoc interface{}) (result *mongo.InsertOneResult, err error) {

	collection := client.Database("school").Collection("students")
	result, err = collection.InsertOne(ctx, studentDoc)
	return
}

func QueryStudents(client *mongo.Client, ctx context.Context, query, field interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database("school").Collection("students")
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return

}
func QuerySingleStudent(client *mongo.Client, ctx context.Context, query, field interface{}) (result bson.M, err error) {

	collection := client.Database("school").Collection("students")
	err = collection.FindOne(ctx, query, options.FindOne().SetProjection(field)).Decode(&result)
	return
}

func UpdateStudentQuery(client *mongo.Client, ctx context.Context, filter bson.D, update models.Student) (result *mongo.UpdateResult, err error) {

	collection := client.Database("school").Collection("students")
	updateDoc := bson.D{{"$set", update}}
	result, err = collection.UpdateOne(ctx, filter, updateDoc)
	return

}
func DeleteStudentQuery(client *mongo.Client, ctx context.Context, filter interface{}) (result *mongo.DeleteResult, err error) {

	collection := client.Database("school").Collection("students")

	result, err = collection.DeleteOne(ctx, filter)
	return

}
