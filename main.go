package main


import (
	"fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/events", getEvents) // Get all dvents in JSON format.
	router.GET("/event/:id", getEventByID) // Get an event by ID.
	router.POST("/event", postEvent) // Add an event.
	router.POST("/events", postEvents) // Add multiple events.

    router.Run("localhost:8080")
}

// An Event represents an event that will occur.
type Event struct {
	ID			string  `json:"id"`
	Title	 	string 	`json:"title"`
	Description	string 	`json:"description"`
	Day			string 	`json:"day"`
	Month		string 	`json:"month"`
	Year		string 	`json:"year"`
	Time		string 	`json:"time"`
	Duration	string	`json:"duration"`
	Location	string 	`json:"location"`
	Recurring 	bool	`json:"recurring"`
	Completed 	bool	`json:"completed"`
}

// example events array
var events = []Event {
	{
		ID: "0", 
		Title: "Event 0", Description: "This is event 0", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: "1", 
		Title: "Event 1", Description: "This is event 1", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: "2", 
		Title: "Event 2", Description: "This is event 2", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: "3", 
		Title: "Event 3", Description: "This is event 3", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: "4", 
		Title: "Event 4", Description: "This is event 4", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: "5", 
		Title: "Event 5", Description: "This is event 5", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
}

// getEvents responds with the list of all events as JSON.
// ex: curl http://localhost:8080/events --include --request "GET"
func getEvents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, events)
}

// getEventByID locates the event whose ID value matches the id
// parameter sent by the client, then returns that event json as a response.
// ex: curl http://localhost:8080/event/2 --include --request "GET"
func getEventByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of events, looking for
    // an event whose ID value matches the parameter.
    for _, e := range events {
        if e.ID == id {
            c.IndentedJSON(http.StatusOK, e) // Returns the event if found.
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": 404}) // Resource (the event) not found. 
}

// postEvent adds an Event from the JSON received in the request body.
// ex: curl http://localhost:8080/event --include --request "POST" --header "Content-Type: application/json" --data '{"ID": "6", "Title": "Event 6", "Description": "This is event 6", "Day": "1", "Month": "January", "Year": "2019", "Time": "8:00", "Duration": "1 hour", "Location": "Room 1", "Recurring": false, "Completed": false}'
func postEvent(c *gin.Context) {
    var newEvent Event

    // Call BindJSON to bind the received JSON to
    // newEvent.
    if err := c.BindJSON(&newEvent); err != nil {
		fmt.Println(c)
		fmt.Println(err)
        return
    }

    // Add the new event to the slice.
    events = append(events, newEvent)
    c.IndentedJSON(http.StatusCreated, newEvent)
}

// Doesnt work currently
// ERROR: 400, invalid character '{' looking for beginning of object key string.
// // postEvents adds several Events from the JSON received in the request body.
// // ex: curl http://localhost:8080/events --include --request "POST" --header "Content-Type: application/json" --data '{{"ID": 6, "Title": "Event 6", "Description": "This is event 6", "Day": "1", "Month": "January", "Year": "2019", "Time": "8:00", "Duration": "1 hour", "Location": "Room 1", "Recurring": false, "Completed": false}, {"ID": "7", "Title": "Event 7", "Description": "This is event 7", "Day": "1", "Month": "January", "Year": "2019", "Time": "8:00", "Duration": "1 hour", "Location": "Room 1", "Recurring": false, "Completed": false}}'
// func postEvents(c *gin.Context) {
// 	var newEvents []Event

// 	// Call BindJSON to bind the received JSON to
// 	// newEvents.
// 	if err := c.BindJSON(&newEvents); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Add the new events to the slice.
// 	events = append(events, newEvents...)
// 	c.IndentedJSON(http.StatusCreated, newEvents)
// }