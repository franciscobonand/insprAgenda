package models

import (
	"database/sql"
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
	Title, Description, Dependency                         string
	Deadline, WorkStart, WorkEnd, CreationDate, LastUpdate time.Time
}

// CreateTask inserts task into DB
func CreateTask(title, description string, priority, timeEstimate int, deadline time.Time, dependency string) {
	creationDate := time.Now()
	if dependency == "" {
		dependency = "0"
	}

	db := db.ConnectWithDB()
	dbColumns := "priority, status, title, description, dependency, deadline, " +
		"workstart, workend, creationdate, lastupdate, timeestimate"
	insertIntoDB, err := db.Prepare("insert into tasks(" + dbColumns + ") " +
		"values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		panic(err.Error())
	}

	insertIntoDB.Exec(priority, 1, title, description, dependency, deadline, time.Time{}, time.Time{}, creationDate, creationDate, timeEstimate)
	fmt.Println("\nA new task was created!")
	defer db.Close()
}

// ShowTasksByStatus recieves a list of status and displays all tasks
// with those status
func ShowTasksByStatus(order string, statusList ...int) {
	// Joins array of ints into a string of numbers separeted by ','
	statusStr := strings.Trim(strings.Replace(fmt.Sprint(statusList), " ", ",", -1), "[]")
	query := generateQuery(order, statusStr)

	db := db.ConnectWithDB()

	tasks, err := db.Query(query)
	if err != nil {
		panic("Error while selecting tasks from DB:" + err.Error())
	}

	printDBRows(tasks, statusList)
}

// UpdateTask can move the task on the board if it doesn't have pendent dependency
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
			var dependency string
			task, err := db.Query("select status, dependency from tasks where id=" + id)
			if err != nil {
				panic("Error while selecting task from DB:" + err.Error())
			}

			task.Next()
			err = task.Scan(&statusNow, &dependency)
			if err != nil {
				panic("Error while getting task status:" + err.Error())
			}

			var newStatus, query string
			var hasDependency bool
			if statusNow == 1 {
				hasDependency = false
				newStatus = strconv.Itoa(statusNow + 1)
				query = "update tasks set status=" + newStatus + ", workstart='" + time.Now().Format("2006-01-02 15:04:05") + "' where id=" + id
			} else {
				hasDependency = checkForDependency(dependency)
				newStatus = strconv.Itoa(statusNow + 2)
				query = "update tasks set status=" + newStatus + ", workend='" + time.Now().Format("2006-01-02 15:04:05") + "' where id=" + id
			}

			if hasDependency {
				fmt.Println("Task (ID:" + id + ") can't be concluded because it depends on another task (ID:" + dependency + ")")
			} else {
				updateTaskDB, err := db.Prepare(query)
				if err != nil {
					panic("Error when updating task:" + err.Error())
				}
				updateTaskDB.Exec()
				fmt.Println("Task (ID:" + id + ") status has been updated!")
			}
		}
	}

	defer db.Close()
}

// DisplayTask gets task with informed ID and display it's information
func DisplayTask(taskID int) {
	db := db.ConnectWithDB()
	id := strconv.Itoa(taskID)

	task, err := db.Query("select * from tasks where id=" + id)
	if err != nil {
		panic("Error while selecting task from DB:" + err.Error())
	}

	task.Next()

	var ID, priority, status, effort int
	var title, description, dependency string
	var deadline, workstart, workend, creationdate, lastupdate time.Time

	err = task.Scan(&ID, &priority, &status, &title, &description, &dependency, &deadline,
		&workstart, &workend, &creationdate, &lastupdate, &effort)
	if err != nil {
		panic("Error while scanning DB:" + err.Error())
	}

	strPriority := strconv.Itoa(priority)
	strTimeEstimate := strconv.Itoa(effort)
	strStatus := []string{"To Do", "Working", "Closed", "Done"}
	timeSpent := getTimeSpent(workstart, workend)

	fmt.Println("\nID: " + id + " | Title: " + title +
		"Description: " + description +
		"Status: " + strStatus[status-1] +
		"\nPriority: " + strPriority + " | Time estimate: " + strTimeEstimate + "h" +
		"\nDeadline: " + deadline.Format("02/01/2006") +
		"\nCreation date: " + creationdate.Format("02/01/2006") +
		// "\nLast update: " + lastupdate.Format("02/01/2006") +
		"\nTime spent on task: " + timeSpent)
}

// DisplayByFilter displays filtered tasks by status
func DisplayByFilter(filterKind int, filter string) {
	db := db.ConnectWithDB()
	allStatus := []int{1, 2, 3, 4}

	switch filterKind {
	case 1: // Deadline
		date := formatDate(filter)
		tasks, err := db.Query("select * from tasks where deadline='" + date + "'")
		if err != nil {
			panic("Error while selecting task from DB:" + err.Error())
		}

		printDBRows(tasks, allStatus)

	case 2: //Priority
		tasks, err := db.Query("select * from tasks where priority=" + filter)
		if err != nil {
			panic("Error while selecting task from DB:" + err.Error())
		}

		printDBRows(tasks, allStatus)

	case 3: // Added time
		date := formatDate(filter)
		tasks, err := db.Query("select * from tasks where creationdate='" + date + "'")
		if err != nil {
			panic("Error while selecting task from DB:" + err.Error())
		}

		printDBRows(tasks, allStatus)

	default:
		fmt.Println("Filter error")
	}
}

// GetActiveTaskIDs returns slice of ints with ids of the active tasks (To Do/Working)
func GetActiveTaskIDs() []int {
	var ids []int
	db := db.ConnectWithDB()

	tasks, err := db.Query("select * from tasks where status in (1,2) order by id asc")
	if err != nil {
		panic("Error while selecting tasks from DB:" + err.Error())
	}

	for tasks.Next() {
		var ID, priority, status, effort int
		var title, description, dependency string
		var deadline, workstart, workend, creationdate, lastupdate time.Time

		err := tasks.Scan(&ID, &priority, &status, &title, &description, &dependency, &deadline,
			&workstart, &workend, &creationdate, &lastupdate, &effort)
		if err != nil {
			panic("Error while scanning DB:" + err.Error())
		}

		ids = append(ids, ID)
	}
	return ids
}

// AUXILIAR/CONTEXT FUNCTIONS:

func printTaskInfo(task Task) {
	id := strconv.Itoa(task.ID)
	priority := strconv.Itoa(task.Priority)
	fmt.Println("ID: " + id + " | Title: " + task.Title +
		" | Priority: " + priority + " | Deadline: " + task.Deadline.Format("02/01/2006") +
		" | Creation date: " + task.CreationDate.Format("02/01/2006"))
}

func getTimeSpent(start, end time.Time) string {
	startEditTimezone := start.Add(time.Hour * 3)
	startYear, _, _ := start.Date()

	if startYear != 1 {
		endYear, _, _ := end.Date()
		if endYear != 1 {
			duration := end.Sub(start)
			return duration.String()
		}
		duration := (time.Now()).Sub(startEditTimezone)
		return duration.String()
	}
	return "Task hasn't been started"
}

func generateQuery(filter, statusList string) string {
	if filter == "dl" {
		return "select * from tasks where status in (" + statusList + ") order by deadline asc"
	} else if filter == "prt" {
		return "select * from tasks where status in (" + statusList + ") order by priority desc"
	} else if filter == "at" {
		return "select * from tasks where status in (" + statusList + ") order by creationdate asc"
	}
	return "select * from tasks where status in (" + statusList + ")"
}

func checkForDependency(dependency string) bool {
	if dependency != "0" {
		var status int
		db := db.ConnectWithDB()
		task, err := db.Query("select status from tasks where id=" + dependency)
		if err != nil {
			panic("Error while selecting task from DB:" + err.Error())
		}

		task.Next()
		err = task.Scan(&status)
		if err != nil {
			panic("Error while getting task status:" + err.Error())
		}

		if status < 3 {
			return true
		}
		return false
	}
	return false
}

func printDBRows(tasks *sql.Rows, statusList []int) {
	var toDoTasks []Task
	var workingTasks []Task
	var closedTasks []Task
	var doneTasks []Task

	for tasks.Next() {
		var ID, priority, status, effort int
		var title, description, dependency string
		var deadline, workstart, workend, creationdate, lastupdate time.Time

		err := tasks.Scan(&ID, &priority, &status, &title, &description, &dependency, &deadline,
			&workstart, &workend, &creationdate, &lastupdate, &effort)
		if err != nil {
			panic("Error while scanning DB:" + err.Error())
		}

		auxTask := Task{
			ID:           ID,
			Title:        title,
			Priority:     priority,
			Status:       status,
			Deadline:     deadline,
			CreationDate: creationdate,
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
		fmt.Println("\nTO DO:")
		if len(toDoTasks) > 0 {
			for _, item := range toDoTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 2) < len(statusList) {
		fmt.Println("\nWORKING:")
		if len(workingTasks) > 0 {
			for _, item := range workingTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 3) < len(statusList) {
		fmt.Println("\nCLOSED:")
		if len(closedTasks) > 0 {
			for _, item := range closedTasks {
				printTaskInfo(item)
			}
		} else {
			fmt.Println("No tasks with this status")
		}
	}
	if sort.SearchInts(statusList, 4) < len(statusList) {
		fmt.Println("\nDONE:")
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

func formatDate(date string) string {
	dateParts := strings.Split(date, "/")
	return dateParts[2] + "-" + dateParts[1] + "-" + dateParts[0]
}
