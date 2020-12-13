package models

import (
	"fmt"
	"insprTaskScheduler/insprAgenda/db"
	"sort"
	"strconv"
	"strings"
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

	insertIntoDB.Exec(priority, 1, title, description, "{0}", deadline, time.Time{}, time.Time{}, creationDate, creationDate, timeEstimate)
	fmt.Println("A new task was created!")
	defer db.Close()
}

// ShowTasksByStatus recieves a list of status and displays all tasks
// with those status
func ShowTasksByStatus(statusList ...int) {
	statusStr := strings.Trim(strings.Replace(fmt.Sprint(statusList), " ", ",", -1), "[]")
	db := db.ConnectWithDB()

	tasks, err := db.Query("select * from tasks where status in (" + statusStr + ")")
	if err != nil {
		panic("Error while selecting tasks from DB:" + err.Error())
	}

	var toDoTasks []Task
	var workingTasks []Task
	var closedTasks []Task
	var doneTasks []Task

	for tasks.Next() {
		var ID, priority, status, effort int
		var title, description string
		var deadline, workstart, workend, creationdate, lastupdate time.Time
		var dependency []uint8

		err = tasks.Scan(&ID, &priority, &status, &title, &description, &dependency, &deadline,
			&workstart, &workend, &creationdate, &lastupdate, &effort)
		if err != nil {
			panic("Error while scanning DB:" + err.Error())
		}

		auxTask := Task{
			ID:       ID,
			Title:    title,
			Priority: priority,
			Status:   status,
			Deadline: deadline,
		}

		switch status {
		case 2:
			workingTasks = append(workingTasks, auxTask)
		case 3:
			closedTasks = append(closedTasks, auxTask)
		case 4:
			doneTasks = append(doneTasks, auxTask)
		default:
			toDoTasks = append(toDoTasks, auxTask)
		}
	}

	if sort.SearchInts(statusList, 1) < len(statusList) {
		fmt.Println("TO DO:")
		if len(toDoTasks) > 0 {
			for _, item := range toDoTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 2) < len(statusList) {
		fmt.Println("WORKING:")
		if len(workingTasks) > 0 {
			for _, item := range workingTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 3) < len(statusList) {
		fmt.Println("CLOSED:")
		if len(closedTasks) > 0 {
			for _, item := range closedTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 4) < len(statusList) {
		fmt.Println("DONE:")
		if len(doneTasks) > 0 {
			for _, item := range doneTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	fmt.Println("")
}

// UpdateTask can move the task on the board
// or uptade it's title/description (yet to be implemented)
func UpdateTask(taskID int, statusUpdate, remove bool) {
	db := db.ConnectWithDB()
	id := strconv.Itoa(taskID)

	if statusUpdate {
		if remove {
			updateTaskDB, err := db.Prepare("update tasks set status=3 where id=" + id)
			if err != nil {
				panic("Error when updating task:" + err.Error())
			}
			updateTaskDB.Exec()
			fmt.Println("Task (ID:" + id + ") has been removed!")
		} else {
			var statusNow int
			task, err := db.Query("select status from tasks where id=" + id)
			if err != nil {
				panic("Error while selecting task from DB:" + err.Error())
			}

			task.Next()
			err = task.Scan(&statusNow)
			if err != nil {
				panic("Error while getting task's status:" + err.Error())
			}

			var newStatus string
			if statusNow == 1 {
				newStatus = strconv.Itoa(statusNow + 1)
			} else {
				newStatus = strconv.Itoa(statusNow + 2)
			}
			updateTaskDB, err := db.Prepare("update tasks set status=" + newStatus + " where id=" + id)
			if err != nil {
				panic("Error when updating task:" + err.Error())
			}
			updateTaskDB.Exec()
			fmt.Println("Task's (ID:" + id + ") status has been updated!")
		}
	}

	defer db.Close()
}

// Auxiliar functions:

func printTaskInfo(task Task) {
	id := strconv.Itoa(task.ID)
	priority := strconv.Itoa(task.Priority)
	fmt.Println("ID: " + id + " | Title: " + task.Title +
		" | Priority: " + priority + " | Deadline: " + task.Deadline.Format("02/01/2006"))
}
