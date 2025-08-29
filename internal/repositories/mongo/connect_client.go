package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbURI string) (client *mongo.Client, err error) {
	//Set a timeout  to allow  the connection process to abort if it takes to long
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Connect to the MongoDB server
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))

	if err != nil {
		return nil, err
	}

	//Call the Ping method to verify that connection has been established successfully
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
