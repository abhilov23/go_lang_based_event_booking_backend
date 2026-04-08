package routes

import (
	"net/http"

	"github.com/abhilov23/gin_project/models"
	"github.com/gin-gonic/gin"
)


func signup(context *gin.Context) {
   var user models.User

   err := context.ShouldBindJSON(&user)

   if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
	return
   }
   
   err = user.Save()

   if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user, try again later"})
	return
   }

   context.JSON(http.StatusCreated, gin.H{"user created": user})
}