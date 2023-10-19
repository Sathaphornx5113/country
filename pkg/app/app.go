package app

import (
	"context"
	"country/pkg/database"
	"country/pkg/domain"
	"fmt"
)

func FindAllCountries() ([]domain.CountryData, error) {
	client := database.DbConn()
	if client == nil {
		return nil, fmt.Errorf("Failed to connect to MongoDB")
	}

	defer client.Disconnect(context.Background())

	collection := client.Database("country_project").Collection("country")

	ctx := context.Background()
	cur, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var countries []domain.CountryData
	for cur.Next(ctx) {
		var country domain.CountryData
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
