package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const API_PORT = 8080

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	if err := router.Run(fmt.Sprintf(":%d", API_PORT)); err != nil {
		log.Fatal(err)
	}
}
