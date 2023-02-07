package configs

import (
	"context"
	"fiber_mongo_zap/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		logger.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //nolint
	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("task").Collection(collectionName)
	return collection
}

func CreateIndex() {
	var declarationCollection *mongo.Collection = GetCollection(DB, "declaration")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := declarationCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "decs.items.items.price", Value: 1},
				{Key: "decs.items.items.quantity", Value: 1},
			},
		},
	)
	if err != nil {
		logger.Error("create index err ", err)
	}
}
