package controllers

import (
	"net/http"
	"strconv"

	"github.com/Jellicle-Cats/user-service/models"
	"github.com/Jellicle-Cats/user-service/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (userController UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}

func (userController UserController) GetUserHistory(ctx *gin.Context) {
	userId := ctx.Query("userId")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "userId is required",
		})
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userId format",
		})
		return
	}
	bookingList, err := userController.UserService.RequestUserHistory(int64(userIdInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user history",
		})
	} else {
		ctx.JSON(http.StatusOK, bookingList)
	}

}
