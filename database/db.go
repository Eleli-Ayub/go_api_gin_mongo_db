package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CloseDB(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("Database disconected")
	}()

}

func ConnectToDB(mongo_url string) (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)

	fmt.Println("----------------------------------------")
	fmt.Println("---------Connection established---------")
	fmt.Println("----------------------------------------")

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongo_url))
	return

}

func PingDB(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("----------------------------------------")
	fmt.Println("---DB Pinged!! Connection established---")
	fmt.Println("----------------------------------------")
	return nil
}

func DbConnect() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {
	// connect to database
	client, ctx, cancel, err = ConnectToDB("mongodb+srv://lesosioayub:ED0zc4tNdxQGy25c@golang.brga42o.mongodb.net/?retryWrites=true&w=majority")
	return
}
