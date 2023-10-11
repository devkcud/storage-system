package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/devkcud/storage-system/api/internal/collection"
	"github.com/devkcud/storage-system/api/internal/connection"
	"github.com/gin-gonic/gin"
)

const API_PORT = 8080

func main() {
	client, err := connection.Open("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	client.SetDatabase("storage")

	log.Println("Connected to MongoDB")

	collectionItems := collection.New[collection.Item](client, "items")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", collectionItems.GetAll)
	router.GET("/:id", collectionItems.GetSpecific)
	router.POST("/", collectionItems.Post)
    router.PATCH("/:id", collectionItems.Patch)
    router.DELETE("/:id", collectionItems.Delete)

	go func() {
		log.Println("Starting API server")

		if err := router.Run(fmt.Sprintf(":%d", API_PORT)); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan // So it properly exits (CTRL + C)

	log.Println("Goodbye! Closing MongoDB connection")
	defer client.Close()
}
