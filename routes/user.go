package routes

import (
	"net/http"

	"example.com/go-api/models"
	"example.com/go-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request while signup"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error, could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "sign up success"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request while login"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "failed authentication"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "failed token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
