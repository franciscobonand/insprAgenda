package controllers

import (
	"bufio"
	"fmt"
	"insprTaskScheduler/insprAgenda/models"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// GenerateNewTask gets input for a new Task and send it to be created in the DB
func GenerateNewTask() {
	var title, description, dependency string
	var timeEstimate, haveDependencies int

	input := bufio.NewReader(os.Stdin)
	// Skips '\n' originated from previous Scan
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
	priority := getValidPriority()

	fmt.Println("Estimated time for task conclusion (in hours - integer):")
	fmt.Scan(&timeEstimate)
	invalidTimeEstimate := timeEstimate < 1
	for invalidTimeEstimate {
		fmt.Println("Invalid estimated time. Insert integer that represents number of hours:")
		fmt.Scan(&timeEstimate)
		invalidTimeEstimate = timeEstimate < 1
	}

	fmt.Println("Task deadline (dd/MM/yyyy):")
	deadline := getValidDate()

	fmt.Println("Does this task have dependencies? (1-yes/2-no)")
	fmt.Scan(&haveDependencies)
	invalidDependencyOption := haveDependencies != 1 && haveDependencies != 2
	for invalidDependencyOption {
		fmt.Println("Invalid option. Task have dependencies? (1-yes/2-no)")
		fmt.Scan(&haveDependencies)
		invalidDependencyOption = haveDependencies != 1 && haveDependencies != 2
	}
	if haveDependencies == 1 {
		dependency = getValidDependency()
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
	validIDs := models.GetActiveTaskIDs()

	if len(validIDs) != 0 {
		fmt.Println("List of active tasks:")
		fmt.Println("")
		models.ShowTasksByStatus("none", 1, 2)

		var taskID int
		if remove {
			fmt.Println("Inform the ID of the task to be removed:")
		} else {
			fmt.Println("Inform the ID of the task to be updated:")
		}
		fmt.Scan(&taskID)

		invalidID := sort.SearchInts(validIDs, taskID) >= len(validIDs)
		for invalidID {
			fmt.Println("Invalid task ID. Insert valid active task ID:")
			fmt.Scan(&taskID)
			invalidID = sort.SearchInts(validIDs, taskID) >= len(validIDs)
		}

		models.UpdateTask(taskID, true, remove)
	} else {
		fmt.Println("There are no active tasks to be updated!")
	}
}

// ShowTaskDetails lists tasks and select the one to have its detals shown
func ShowTaskDetails() {
	models.ShowTasksByStatus("none", 1, 2, 3, 4)

	var taskID int
	fmt.Println("Inform the ID of the task to display it's details:")
	fmt.Scan(&taskID)

	models.DisplayTask(taskID)
}

// ShowTasksByOrder tells method ShowTasksByStatus which filter to use
func ShowTasksByOrder(orderOption int) {
	switch orderOption {
	case 1:
		models.ShowTasksByStatus("dl", 1, 2, 3, 4)
	case 2:
		models.ShowTasksByStatus("prt", 1, 2, 3, 4)
	case 3:
		models.ShowTasksByStatus("at", 1, 2, 3, 4)
	default:
		fmt.Println("This order option doesn't exist")
	}
}

// ShowFilteredOptions asks user to input the filters value
func ShowFilteredOptions(filterKind int) {
	switch filterKind {
	case 1: // Deadline
		fmt.Println("Input deadline date (dd/MM/yyyy):")
		deadline := getValidDate()
		fmt.Println("")
		models.DisplayByFilter(filterKind, deadline)
	case 2: //Priority
		fmt.Println("Input priority (1 to 10):")
		priority := getValidPriority()
		fmt.Println("")
		models.DisplayByFilter(filterKind, strconv.Itoa(priority))
	case 3: // Added time
		fmt.Println("Input creation date (dd/MM/yyyy):")
		createdAt := getValidDate()
		fmt.Println("")
		models.DisplayByFilter(filterKind, createdAt)
	default:
		fmt.Println("This filter doesn't exist")
	}
}

// AUXILIAR FUNCTIONS:

func getValidPriority() int {
	var priority int
	fmt.Scan(&priority)
	invalidPriority := priority < 1 || priority > 10
	for invalidPriority {
		fmt.Println("Please insert a valid priority (1 to 10)")
		fmt.Scan(&priority)
		invalidPriority = priority < 1 || priority > 10
	}
	return priority
}

func getValidDate() string {
	var deadline string
	var invalidDeadline bool
	fmt.Scan(&deadline)
	date := strings.Split(deadline, "/")

	if len(date) == 3 {
		invalidDeadline = checkDate(date)
	} else {
		invalidDeadline = true
	}

	for invalidDeadline {
		fmt.Println("Invalid deadline. Insert a valid date (dd/MM/yyyy)")
		fmt.Scan(&deadline)
		date = strings.Split(deadline, "/")

		if len(date) == 3 {
			invalidDeadline = checkDate(date)
		} else {
			invalidDeadline = true
		}
	}
	return deadline
}

// Returns true if invalid date
func checkDate(date []string) bool {
	dateDay, errDay := strconv.Atoi(date[0])
	dateMonth, errMonth := strconv.Atoi(date[1])
	dateYear, errYear := strconv.Atoi(date[2])
	if errDay != nil {
		panic("Day convertion error:" + errDay.Error())
	}
	if errMonth != nil {
		panic("Month convertion error:" + errMonth.Error())
	}
	if errYear != nil {
		panic("Year convertion error:" + errYear.Error())
	}
	if dateDay < 1 || dateDay > 31 {
		return true
	}
	if dateMonth < 1 || dateMonth > 12 {
		return true
	}
	if dateYear < 2020 || dateYear > 2999 {
		return true
	}
	return false
}

func getValidDependency() string {
	validIDs := models.GetActiveTaskIDs()
	var dependency int

	if len(validIDs) != 0 {
		fmt.Println("List of active tasks:")
		models.ShowTasksByStatus("none", 1, 2)
		fmt.Println("")
		fmt.Println("Insert ID of task depended on:")
		fmt.Scan(&dependency)
		invalidDependency := sort.SearchInts(validIDs, dependency) >= len(validIDs)

		for invalidDependency {
			fmt.Println("Invalid task ID. Insert valid active task ID:")
			fmt.Scan(&dependency)
			invalidDependency = sort.SearchInts(validIDs, dependency) >= len(validIDs)
		}
	} else {
		fmt.Println("There are no active tasks to be depended on!")
		dependency = 0
	}

	return strconv.Itoa(dependency)
}
