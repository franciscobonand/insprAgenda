package models

import (
	"insprTaskScheduler/insprAgenda/db"
	"time"
)

// Task field composition
type Task struct {
	ID, Priority, Status, TimeEstimate                     int
	Title, Description                                     string
	Dependency                                             []int
	Deadline, WorkStart, WorkEnd, CreationDate, LastUpdate time.Time
}

// CreateTask inserts task into DB
func CreateTask(title, description string, priority, timeEstimate int, deadline time.Time) {
	creationDate := time.Now()

	db := db.ConnectWithDB()
	dbColumns := "priority, status, title, description, dependency, deadline, " +
		"workstart, workend, creationdate, lastupdate, timeestimate"
	insertIntoDB, err := db.Prepare("insert into tasks(" + dbColumns + ") " +
		"values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		panic(err.Error())
	}

	insertIntoDB.Exec(priority, 1, title, description, nil, deadline, nil, nil, creationDate, creationDate, timeEstimate)
	defer db.Close()
}

// ShowTasksByStatus recieves a list of status and displays all tasks
// with those status
func ShowTasksByStatus(status ...int) {
	db := db.ConnectWithDB()

	tasks, err := db.Query("select * from tasks where status in $1", status)
	if err != nil {
		panic("Error while selecting tasks from DB:" + err.Error())
	}

	for tasks.Next() {
		// var ID, priority, effort int
		// var title string
		// var deadline time.Time

		err = tasks.Scan()
		if err != nil {
			panic("Error while scanning DB:" + err.Error())
		}
	}

	// Still need to adjust the input parameters and DB scanning
}
