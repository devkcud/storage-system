package connection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Open(mongoURI string) (*Connection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &Connection{client, nil}, nil
}

func (c *Connection) Close() {
	if err := c.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
