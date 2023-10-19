package main

import (
	"context"
	"country/pkg/database"
	"fmt"
)

func main() {
	client := database.DbConn()
	if client == nil {
		fmt.Println("Failed to connect to MongoDB")
		return
	}

	// Your code to work with 'client' here

	defer client.Disconnect(context.Background())
	fmt.Println("Connected to MongoDB!")
}
