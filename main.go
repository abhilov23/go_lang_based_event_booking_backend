package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	// here the default() is the gin router
    server:= gin.Default()
    
	// http support different kinds of requests: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS etc.
	server.GET("/events", getEvents) 

	server.Run(":8080") // localhost 8080
}

// here the context is the request and response object which provides different methods and functions 
// to handle the request and response
func getEvents(context *gin.Context){
    
	// here the JSON method is used to send the response, 
	// and gin.H is a shortcut for map[string]interface{} 
    context.JSON(http.StatusOK, gin.H{"message": "Hello from the events endpoint"})
}