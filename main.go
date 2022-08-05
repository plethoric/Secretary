package main


import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/events", getEvents)

    router.Run("localhost:8080")
}

// An Event represents an event that will occur.
type Event struct {
	ID			int    	`json:"id"`
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

// example events
var events = []Event{
	{
		ID: 0, 
		Title: "Event 0", Description: "This is event 0", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: 1, 
		Title: "Event 1", Description: "This is event 1", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: 2, 
		Title: "Event 2", Description: "This is event 2", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: 3, 
		Title: "Event 3", Description: "This is event 3", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: 4, 
		Title: "Event 4", Description: "This is event 4", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
	{
		ID: 5, 
		Title: "Event 5", Description: "This is event 5", 
		Day: "1", Month: "January", Year: "2019", Time: "8:00", Duration: "1 hour", 
		Location: "Room 1", 
		Recurring: false, Completed: false,
	},
}

// getEvents responds with the list of all events as JSON.
func getEvents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, events)
}