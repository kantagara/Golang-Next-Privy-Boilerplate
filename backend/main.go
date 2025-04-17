package main

import (
	"backend/internal/auth"
	"backend/internal/common"
	"backend/internal/routes"
	"backend/internal/user"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	loadEnv()
	println("Succesfully loaded environment variables")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mongoClient := connectToMongo(ctx)
	println("Connected to MongoDB")

	go func() {
		startGin(mongoClient)
	}()

	<-ctx.Done()
	println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mongoClient.Disconnect(shutdownCtx); err != nil {
		log.Fatal("Failed to disconnect Mongo: ", err)
	}

	println("Shutdown complete.")
}

func connectToMongo(ctx context.Context) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DATABASE_URL")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("Mongo connect failed: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping failed: ", err)
	}

	return client
}

func startGin(mongoClient *mongo.Client) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userRepo := user.NewRepository(mongoClient.Database("defree"))

	dependencies := &common.Dependencies{
		AuthHandler: &auth.Handler{
			Service: auth.NewService(userRepo),
		},
		UserHandler: &user.Handler{
			UserService: user.NewService(userRepo),
		},
	}

	routes.RegisterRoutes(router, dependencies)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Gin server failed: ", err)
	}
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
