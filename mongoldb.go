package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connString = "mongodb+srv://mainAleks:mongodb@testcluster1.wfmzc1o.mongodb.net/?retryWrites=true&w=majority"
)

func mongoGetAllDevices() []Device_asset {

	start := time.Now()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("testdb").Collection("devices")
	var results []Device_asset
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result Device_asset
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		fmt.Println("loaded data from devices_db  for:", time.Since(start))
	}

	return results

	// Connect to MongoDB

	//fmt.Printf("Found device: %+v\n", results)

}
func mongoSendData(device Device_asset) {

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("testdb").Collection("devices")

	_, err = collection.InsertOne(context.Background(), device)
	if err != nil {
		log.Fatal(err)
	}
}

func mongoGetAllShops() []Shops {

	start := time.Now()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("testdb").Collection("shops")
	var results []Shops
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result Shops
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		fmt.Println("loaded data from shops_db  for:", time.Since(start))
	}

	return results

	//fmt.Printf("Found device: %+v\n", results)

}

func mongoSendShop(device Shops) {

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("testdb").Collection("shops")

	_, err = collection.InsertOne(context.Background(), device)
	if err != nil {
		log.Fatal(err)
	}
}
