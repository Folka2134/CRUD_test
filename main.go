package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	http.HandleFunc("/", handler)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	log.Fatal(router.Run(":" + PORT))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Hello, World!"}`))
}
