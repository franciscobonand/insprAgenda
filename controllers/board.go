package controllers

import (
	"bufio"
	"fmt"
	"insprTaskScheduler/insprAgenda/models"
	"os"
	"strconv"
	"strings"
	"time"
)

// GenerateNewTask gets input for a new Task and send it to be created in the DB
func GenerateNewTask() {
	var title, description, deadline, dependency string
	var priority, timeEstimate, haveDependencies int

	input := bufio.NewReader(os.Stdin)
	input.ReadString('\n')

	fmt.Println("Task title:")
	title, err := input.ReadString('\n')
	if err != nil {
		panic("Error while reading title: " + err.Error())
	}

	fmt.Println("Task description:")
	description, err = input.ReadString('\n')
	if err != nil {
		panic("Error while reading description: " + err.Error())
	}

	fmt.Println("Task priority (1 to 10):")
	fmt.Scan(&priority)

	fmt.Println("Estimated time for task conclusion (in hours - whole number):")
	fmt.Scan(&timeEstimate)

	fmt.Println("Task deadline (dd/MM/yyyy):")
	fmt.Scan(&deadline)

	fmt.Println("Does this task have dependencies? (1-yes/2-no)")
	fmt.Scan(&haveDependencies)
	if haveDependencies == 1 {
		fmt.Println("List of active tasks:")
		models.ShowTasksByStatus("none", 1, 2)
		fmt.Println("")
		fmt.Println("Insert ID of task depended on:")
		fmt.Scan(&dependency)
	}

	deadlineInfo := strings.Split(deadline, "/")
	deadlineDay, errDay := strconv.Atoi(deadlineInfo[0])
	deadlineMonth, errMonth := strconv.Atoi(deadlineInfo[1])
	deadlineYear, errYear := strconv.Atoi(deadlineInfo[2])
	if errDay != nil {
		panic("Day convertion error:" + errDay.Error())
	}
	if errMonth != nil {
		panic("Month convertion error:" + errMonth.Error())
	}
	if errYear != nil {
		panic("Year convertion error:" + errYear.Error())
	}
	deadlineDate := time.Date(deadlineYear, time.Month(deadlineMonth), deadlineDay, 12, 0, 0, 0, time.UTC)

	models.CreateTask(title, description, priority, timeEstimate, deadlineDate, dependency)
}

// MoveTaskOnBoard lists tasks and select the one to have it's status updated
func MoveTaskOnBoard(remove bool) {
	fmt.Println("Active tasks:")
	fmt.Println("")
	models.ShowTasksByStatus("none", 1, 2)

	var taskID int
	if remove {
		fmt.Println("Inform the ID of the task to be removed:")
	} else {
		fmt.Println("Inform the ID of the task to be updated:")
	}
	fmt.Scan(&taskID)
	models.UpdateTask(taskID, true, remove)
}

// ShowTaskDetails lists tasks and select the one to have its detals shown
func ShowTaskDetails() {
	models.ShowTasksByStatus("none", 1, 2, 3, 4)

	var taskID int
	fmt.Println("Inform the ID of the task to display it's details:")
	fmt.Scan(&taskID)

	models.DisplayTask(taskID)
}

// ShowTasksByFilter tell method ShowTasksByStatus which filter to use
func ShowTasksByFilter(filterOption int) {
	switch filterOption {
	case 1:
		models.ShowTasksByStatus("dl", 1, 2, 3, 4)
	case 2:
		models.ShowTasksByStatus("prt", 1, 2, 3, 4)
	case 3:
		models.ShowTasksByStatus("at", 1, 2, 3, 4)
	default:
		fmt.Println("This filter doesn't exist")
	}
}
