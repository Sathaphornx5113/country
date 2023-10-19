package main

import (
	"context"
	"fmt"

	"github.com/Sathaphornx5113/country/pkg/database"
)

func main() {
	client, err := database.DbConn()
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

}
