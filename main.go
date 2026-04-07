package main

import (
	"net/http"

	"github.com/abhilov23/gin_project/db"
	"github.com/abhilov23/gin_project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	// here the default() is the gin router
	server := gin.Default()

	// http support different kinds of requests: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS etc.
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost 8080
}

// here the context is the request and response object which provides different methods and functions
// to handle the request and response
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events, try again later"})
		return
	}
	// here the JSON method is used to send the response,
	// and gin.H is a shortcut for map[string]interface{}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// this will only accept the json data defined in the event struct
	// but it is not strictly required and can be ignored
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err =event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event created": event})
}
