package main

import (
	"log"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"github.com/daxpat2000/cricserver/internal/core/services"
	"github.com/daxpat2000/cricserver/internal/handlers"
	sqlite_repository "github.com/daxpat2000/cricserver/internal/repositories/sqlite"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize sqlite db
	sqliteDB, err := gorm.Open(sqlite.Open("cricket.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate
	sqliteDB.AutoMigrate(&domain.Player{}, &domain.Team{})

	// Set up repository
	sqliteRepository := sqlite_repository.NewSQLiteRepository(sqliteDB)

	// Create a new player handler
	playerHandler := handlers.NewPlayerHandler(services.NewPlayerService(sqliteRepository))

	// Create a new team handler
	teamHandler := handlers.NewTeamHandler(services.NewTeamService(sqliteRepository))

	// Create a new router
	router := gin.Default()
	v1 := router.Group("/v1")

	// Register Routes
	v1.POST("/players", playerHandler.CreatePlayer)
	v1.GET("/players/:id", playerHandler.GetPlayer)
	v1.GET("/players", playerHandler.ListPlayers)

	v1.POST("/teams", teamHandler.CreateTeam)
	v1.GET("/teams/:name", teamHandler.GetTeam)
	v1.GET("/teams", teamHandler.GetTeam)

	// Run service
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run players service: %v", err)
	}
}
