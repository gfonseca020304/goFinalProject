package main

import (
	"fmt"
	"net/http"

	"github.com/gfonseca020304/goFinalProject/db"
	"github.com/gfonseca020304/goFinalProject/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("Some test")
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event                //This is an event instantiation to create one.
	err := context.ShouldBindJSON(&event) //Binds the JSON req to the event struct.

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error:": err})
		fmt.Println(err)
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save() //This method will save the event itself

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) //This is how you return a response to the client.
}
