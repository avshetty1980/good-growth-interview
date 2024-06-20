package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/avshetty1980/good-growth-interview/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content"`
}

type Store interface {
	CreateMessage(string) (string, error)
	GetMessage(string) (string, error)
}

type Storage struct {
	client *mongo.Client
}

func NewStore(c *mongo.Client) *Storage {
	return &Storage{
		client: c,
	}
}

func (s *Storage) CreateMessage(content string) (string, error) {
	coll := s.client.Database(config.Envs.DBName).Collection(config.Envs.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newMessage := Message{
		ID:      primitive.NewObjectID(),
		Content: content,
	}
	res, err := coll.InsertOne(ctx, newMessage)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (s *Storage) GetMessage(id string) (string, error) {
	coll := s.client.Database(config.Envs.DBName).Collection(config.Envs.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messageId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	var message Message
	err = coll.FindOne(ctx, bson.M{"_id": messageId}).Decode(&message)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("message not found")
		}
		return "", err
	}
	fmt.Printf("message recieved: %#v", message)
	return message.Content, nil
}
