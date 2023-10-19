package main

import (
	"/pkg/database/database.go"
	"context"
	"fmt"
)

func main() {
	client, err := database.DbConn()
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

}
