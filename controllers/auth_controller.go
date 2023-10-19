package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jellicle-Cats/user-service/config"
	"github.com/Jellicle-Cats/user-service/models"
	"github.com/Jellicle-Cats/user-service/services"
	"github.com/Jellicle-Cats/user-service/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}
func (authController AuthController) GoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")
	var pathUrl string = "/"

	if ctx.Query("state") != "" {
		pathUrl = ctx.Query("state")
	}

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	googleUser, err := utils.GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userData := &models.User{
		Name:      googleUser.Name,
		Email:     googleUser.Email,
		Photo:     googleUser.Picture,
		Provider:  "Google",
		Role:      "user",
		UpdatedAt: time.Now(),
	}
	updatedUser, err := authController.UserService.UpsertUser(googleUser.Email, userData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	config := config.ReadConfig()

	token, err := utils.GenerateToken(config.TokenExpiresIn, updatedUser.ID.Hex(), config.JWTTokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprint("http://localhost:3000", pathUrl))
}

func (authController AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
