package main

import (
	"context"
	"fmt"

	"github.com/avshetty1980/good-growth-interview/application"
)

func main() {

	server := application.NewAPIServer(":5000")
	err := server.Run(context.TODO())

	if err != nil {
		fmt.Println("Failed to start server", err)
	}

}
