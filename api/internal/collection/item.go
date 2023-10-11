package collection

import (
	"net/http"

	"github.com/devkcud/storage-system/api/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func (i *Item) GetAll(c *gin.Context) {
	var items []bson.M

	cursor, err := i.collection.Find(c.Request.Context(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting items: " + err.Error()})
		return
	}

	if err = cursor.All(c.Request.Context(), &items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting items: " + err.Error()})
		return
	}

	// Add the 'id' field
	itemsJson := make([]map[string]interface{}, len(items))

	for i, item := range items {
		itemsJson[i] = item
		itemsJson[i]["id"] = item["_id"]

		delete(itemsJson[i], "_id")
	}

	c.JSON(http.StatusOK, items)
}

func (i *Item) GetSpecific(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var item model.Item

	if err := i.collection.FindOne(c.Request.Context(), bson.M{"_id": id}).Decode(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}
