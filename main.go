package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/ad/config"
	handlers "github.com/tanush-128/openzo_backend/ad/internal/api"
	"github.com/tanush-128/openzo_backend/ad/internal/pb"
	"github.com/tanush-128/openzo_backend/ad/internal/repository"
	"github.com/tanush-128/openzo_backend/ad/internal/service"
	"google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

type User2 struct {
}

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := connectToDB(cfg) // Implement database connection logic
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database: %w", err))
	}

	imageConn, err := grpc.Dial(cfg.ImageGrpc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer imageConn.Close()
	imageClient := pb.NewImageServiceClient(imageConn)

	adRepository := repository.NewAdRepository(db)

	AdService := service.NewAdService(adRepository, imageClient)

	// Initialize HTTP server with Gin
	router := gin.Default()
	handler := handlers.NewHandler(&AdService)

	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong ads",
		})
	})

	// router.Use(middlewares.JwtMiddleware(c))
	router.POST("/", handler.CreateAd)
	router.GET("/:id", handler.GetAdByID)
	router.GET("/pincode/:pincode", handler.GetAdsByPincode)
	// router.Use(middlewares.JwtMiddleware)
	router.Run(fmt.Sprintf(":%s", cfg.HTTPPort))

}
