package collection

import (
	"github.com/devkcud/storage-system/api/internal/connection"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Generic struct {
	collection *mongo.Collection
}

type Types interface {
	Item
}

func New[T Types](connection *connection.Connection, name string) *T {
	return &T{Generic{connection.Database.Collection(name)}}
}

func (*Generic) Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, world!"})
}
