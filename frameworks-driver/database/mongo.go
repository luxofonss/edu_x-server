package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	DB mongo.Database
}

func (MongoDatabase) Connect(config DBConfig) (DBConnection, error) {
	dsn := fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("has err")
		return nil, err
	}
	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	mongoDb := client.Database(config.DB)
	fmt.Println("You successfully connected to MongoDB!")
	return mongoDb, nil
}
