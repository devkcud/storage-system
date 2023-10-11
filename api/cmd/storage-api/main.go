package main

import (
	"fmt"
	"log"

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

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	if err := router.Run(fmt.Sprintf(":%d", API_PORT)); err != nil {
		log.Fatal(err)
	}

	defer client.Close()
}
