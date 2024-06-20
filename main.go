package main

import (
	"fmt"
	"log"

	"github.com/avshetty1980/good-growth-interview/application"
	"github.com/avshetty1980/good-growth-interview/config"
	"github.com/avshetty1980/good-growth-interview/database"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI(config.Envs.DBURI)

	clientOptions.SetAuth(options.Credential{
		Username: config.Envs.Username,
		Password: config.Envs.Password,
	})
	mongoStorage := database.NewMessageStore(clientOptions)

	store := database.NewStore(mongoStorage.Client, config.Envs.DBName, config.Envs.Collection)

	defer func() {
		err := mongoStorage.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	server := application.NewAPIServer(":5000", store)
	err := server.Run()

	if err != nil {
		fmt.Println("Failed to start server", err)
	}

}
