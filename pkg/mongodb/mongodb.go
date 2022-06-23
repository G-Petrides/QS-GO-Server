package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	core "quaysports.com/server/pkg/core"
)

func dbConnection(config core.Config, channel chan string) {
	var client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.Db.DbAddress + ":" + config.Db.DbPort + "/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("test").Collection("items")
	var result bson.M
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
	}
	cursor.Decode(&result)
	channel <- ""
}
