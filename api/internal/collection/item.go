package collection

import (
	"net/http"

	"github.com/devkcud/storage-system/api/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Generic
}

func (i *Item) Create(c *gin.Context) {
	var item model.Item

	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding JSON: " + err.Error()})
		return
	}

	item.ID = primitive.NewObjectID()

	i.collection.InsertOne(c.Request.Context(), item)

	c.JSON(http.StatusOK, gin.H{"message": "Item created"})
}
