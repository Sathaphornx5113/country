package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url = "mongodb://localhost:27017"

func DbConn() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("Connect to MongoDb: %v failed: %v", url, err)
	}

	return client
}

func FindAllCountries() ([]CountryData, error) {
	client := database.DbConn()
	if client == nil {
		return nil, fmt.Errorf("Failed to connect to MongoDB")
	}

	defer client.Disconnect(context.Background())

	collection := client.Database("your-database-name").Collection("your-collection-name")

	ctx := context.Background()
	cur, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var countries []CountryData
	for cur.Next(ctx) {
		var country CountryData
		err := cur.Decode(&country)
		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return countries, nil
}
