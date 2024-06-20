package database

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	client   *mongo.Client
	dbname   string
	collname string
}

func NewStore(c *mongo.Client, dbn string, cname string) *Storage {
	return &Storage{
		client:   c,
		dbname:   dbn,
		collname: cname,
	}
}

func (s *Storage) CreateMessage(content string) (string, error) {
	coll := s.client.Database(s.dbname).Collection(s.collname)

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
	coll := s.client.Database(s.dbname).Collection(s.collname)

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
