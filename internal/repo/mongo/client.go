package mongo

// import (
// 	"context"
// 	"os"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// func init() {

// 	uri := os.Getenv("MONGO_URI")

// 	c, err := mongo.Connect(
// 		context.Background(),
// 		options.Client().ApplyURI(uri),
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	client = c
// }

// func Client() *mongo.Client {
// 	return client
// }
