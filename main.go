package main

import (
	"log"
	"net/http"

	"gits-assignment/config"
	"gits-assignment/handlers"
	"gits-assignment/middleware"
	"gits-assignment/migration"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadENV()
	config.InitDatabase()
	migration.SyncDatabase()
	migration.SeedDatabase()
}

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Status OK",
		})

		return
	})

	app.POST("/signup", handlers.Register)
	app.POST("/login", handlers.Login)

	// author routes
	app.GET("/author", handlers.GetAllAuthors)
	app.POST("/author", middleware.Authorized(), handlers.AddNewAuthor)
	app.PATCH("/author/:id", middleware.Authorized(), handlers.ModifyAuthor)
	app.DELETE("/author/:id", middleware.Authorized(), handlers.RemoveAuthor)

	// book routes
	app.GET("/book", handlers.GetBooks)
	app.POST("/book", middleware.Authorized(), handlers.AddNewBook)
	app.PATCH("/book/:id", middleware.Authorized(), handlers.ModifyBook)
	app.DELETE("/book/:id", middleware.Authorized(), handlers.RemoveBook)

	// publisher routes
	app.GET("/publisher", handlers.GetAllPublishers)
	app.POST("/publisher", middleware.Authorized(), handlers.AddNewPublisher)
	app.PATCH("/publisher/:id", middleware.Authorized(), handlers.ModifyPublisher)
	app.DELETE("/publisher/:id", middleware.Authorized(), handlers.RemovePublisher)

	err := app.Run(":3000")
	if err != nil {
		log.Fatalf("APP::failed starting the server-> err: %v", err)
	}
}
