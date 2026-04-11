package routes

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/abhilov23/gin_project/models"
	"github.com/abhilov23/gin_project/utils"
	"github.com/gin-gonic/gin"
)

// here the context is the request and response object which provides different methods and functions
// to handle the request and response
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later"})
		return
	}
	// here the JSON method is used to send the response,
	// and gin.H is a shortcut for map[string]interface{}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //converting string into int64

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	var event models.Event
	// this will only accept the json data defined in the event struct
	// but it is not strictly required and can be ignored
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
		return
	}

	event.UserID = int(userID)

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event created": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //converting string into int64

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event, try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //converting string into int64

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later"})
		return
	}

	 err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event, try again later"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
