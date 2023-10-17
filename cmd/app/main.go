package main

import (
	_ "github.com/fur1ouswolf/crud/docs"
	"github.com/fur1ouswolf/crud/internal/handlers"
	"github.com/fur1ouswolf/crud/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db := models.ConnectDB()

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	handler := handlers.NewHandler(db, logger)

	router := gin.Default()

	// Create region
	router.POST("/region", handler.CreateRegion)
	// Get all regions
	router.GET("/region", handler.AllRegions)
	// Delete region
	router.DELETE("/region/:id", handler.DeleteRegion)
	// Get all persons
	router.GET("/person", handler.AllPersons)
	// Get all persons in the region
	router.GET("/person/region/:id", handler.Residents)
	// Create new person
	router.POST("/person", handler.NewPerson)
	// Get information about person
	router.GET("/person/:id", handler.Person)
	// Update information about person
	router.POST("/person/:id", handler.UpdatePerson)
	// Delete person
	router.DELETE("/person/:id", handler.DeletePerson)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		handler.Logger.Error(err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	handler.Logger.Info("Server stopped")
}
