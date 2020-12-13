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

// GenerateNewTask initializes a new Task instance and send it to the DB
func GenerateNewTask() {
	var title, description, deadline string
	var priority, timeEstimate int

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

	fmt.Println("Estimated time for task conclusion (in hours):")
	fmt.Scan(&timeEstimate)

	fmt.Println("Task deadline (dd/MM/yyyy):")
	fmt.Scan(&deadline)

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

	models.CreateTask(title, description, priority, timeEstimate, deadlineDate)
}

// MoveTaskOnBoard lists tasks and select the one to have it's status updated
func MoveTaskOnBoard(remove bool) {
	fmt.Println("Active tasks:")
	fmt.Println("")
	models.ShowTasksByStatus(1, 2)

	var taskID int
	if remove {
		fmt.Println("Inform the ID of the task to be removed:")
		fmt.Scan(&taskID)
	} else {
		fmt.Println("Inform the ID of the task to be updated:")
		fmt.Scan(&taskID)
	}
	// Still have to call the function that makes the update (undergoing implementation)
}
