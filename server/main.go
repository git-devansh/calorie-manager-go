package main

import (
	"os"

	"example.com/calorie-manager-go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/ingredient/:ingredient", routes.GetEntryByIngredient)
	router.GET("/entry/:id", routes.EntryById)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.DELETE("/ingredient/update/:id", routes.UpdateIngredient)
	router.PUT("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":" + port)
}
