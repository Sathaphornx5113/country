package app

import (
	"context"
	"country/pkg/database"
	"fmt"
)

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
