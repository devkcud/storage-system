package connection

import "go.mongodb.org/mongo-driver/mongo"

type Connection struct {
	Client   *mongo.Client
	Database *mongo.Database
}
