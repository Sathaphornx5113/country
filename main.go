package main

import (
	"context"
	"country/pkg/app"
	"country/pkg/database"
	"fmt"
)

func main() {
	client := database.DbConn()
	if client == nil {
		fmt.Println("Failed to connect to MongoDB")
		return
	}

	defer client.Disconnect(context.Background())
	fmt.Println("Connected to MongoDB!")

	countries, err := app.FindAllCountries()
	if err != nil {
		fmt.Printf("Failed to retrieve countries: %v\n", err)
		return
	}

	for _, country := range countries {
		fmt.Printf("Country: %s, Year: %d, Population: %d\n", country.Country, country.Year, country.Population)
	}
}
