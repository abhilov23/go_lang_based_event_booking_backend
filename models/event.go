package models

import (
	"time"

	"github.com/abhilov23/gin_project/db"
)

type Event struct {
	ID          int64
	Name        string  `binding:"required"` // this binding defines that this field is required
	Description string  `binding:"required"`
	Location    string  `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int 
}

// this is a global variable
var events = []Event{}


func (e Event) Save()  error {

	// IN this query we are using the ? placeholder to avoid sql injection
  query := `
  INSERT INTO events(name, description, location, dateTime, user_id)
   VALUES (?, ?, ?, ?, ?) `
   
   stmt, err := db.DB.Prepare(query) // this will prepare the query which means it will be executed only once 
   // and then the statement will be cached in memory for future use
   
   if err !=nil {
	return err
   }
   
   defer stmt.Close() // this will be executed when the function returns because we are using the defer keyword
   result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
   if err != nil {
	return err
   }
   id, err := result.LastInsertId()
   e.ID = id

   return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // this will also execute the query but the difference is that results will not be stored in memory
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event 
	for rows.Next(){
     var event Event 
	 rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	
    
	if err != nil {
		return nil, err
	}
    events = append(events, event)
}
	return events, nil
}