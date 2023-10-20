package server

import (
	"context"
	"log"
	"net/http"

	"github.com/Jellicle-Cats/user-service/config"
	"github.com/Jellicle-Cats/user-service/controllers"
	"github.com/Jellicle-Cats/user-service/database"
	"github.com/Jellicle-Cats/user-service/middleware"
	"github.com/Jellicle-Cats/user-service/proto"
	"github.com/Jellicle-Cats/user-service/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServer() {
	ctx := context.TODO()
	server := gin.Default()
	config := config.ReadConfig()
	mongoClient := database.NewMongoConnection(config.MongoURL)
	defer database.CloseClientDB()

	cred := insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:30043", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	bookingClient := proto.NewBookingServiceClient(conn)
	userCollection := mongoClient.Database("golangAPI").Collection("users")
	userService := services.NewUserService(userCollection, ctx, bookingClient)
	authController := controllers.NewAuthController(userService)
	userController := controllers.NewUserController(userService)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Google OAuth2 in Golang"})
	})
	router.GET("/auth/logout", authController.LogoutUser)
	router.GET("/sessions/oauth/google", authController.GoogleOAuth)
	router.GET("/users/me", middleware.DeserializeUser(userService), userController.GetMe)
	router.GET("/users/bookinglist", userController.GetUserHistory)
	server.Run(":8000")
}
