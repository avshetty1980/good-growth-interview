package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MessageStore struct {
	Client *mongo.Client
}

func NewMessageStore(opts *options.ClientOptions) *MessageStore {

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatalf("mongodb disconnect error : %v", err)
		return nil

	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("mongodb disconnect error : %v", err)
		return nil
	}
	log.Println("ping to Mongo successfull")

	return &MessageStore{
		Client: client,
	}
}

func (ms *MessageStore) Close() error {

	return ms.Client.Disconnect(context.Background())
}
