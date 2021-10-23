package db

import (
	"context"
	"fmt"
	setting "go-template/pkg/settings"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database contains mongo.Database
type DB struct {
	Session *mongo.Client
	Keys    *mongo.Collection // TODO: Rename to KeysCollection
}

var MongoDB *DB

// Connect to MongoDB
func ConnectMongo() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(setting.MongoSetting.Address))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoDB = &DB{
		Session: client,
		Keys:    client.Database(setting.MongoSetting.Database).Collection(setting.MongoSetting.KeysCollection),
	}
	fmt.Println("Connected successfully to Mongo")
}

// Disconnect to MongoDB
func CloseMongo() {
	err := MongoDB.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
