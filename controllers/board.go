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

	fmt.Println("Task deadline (dd/MM/yyyy):")
	deadline := getValidDeadline()
	fmt.Scan(&deadline)

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
	var filterChoice string
	switch filterKind {
	case 1: // Deadline
		fmt.Println("Input deadline date (dd/MM/yyyy):")
		fmt.Scan(&filterChoice)
		fmt.Println("")
		models.DisplayByFilter(filterKind, filterChoice)
	case 2: //Priority
		fmt.Println("Input priority (1 to 10):")
		fmt.Scan(&filterChoice)
		fmt.Println("")
		models.DisplayByFilter(filterKind, filterChoice)
	case 3: // Added time
		fmt.Println("Input creation date (dd/MM/yyyy):")
		fmt.Scan(&filterChoice)
		fmt.Println("")
		models.DisplayByFilter(filterKind, filterChoice)
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

func getValidDeadline() string {
	var deadline string
	var invalidDeadline bool
	fmt.Scan(&deadline)
	date := strings.Split(deadline, "/")

	invalidDeadline, day := checkDay(strconv.Atoi(date[0]))
	invalidDeadline = checkMonth(day, date[1])
	invalidDeadline = checkYear(strconv.Atoi(date[2]))

	for invalidDeadline {
		fmt.Println("Invalid deadline. Insert a valid date (dd/MM/yyyy)")
		fmt.Scan(&deadline)
		date := strings.Split(deadline, "/")

		invalidDeadline, day = checkDay(strconv.Atoi(date[0]))
		invalidDeadline = checkMonth(day, date[1])
		invalidDeadline = checkYear(strconv.Atoi(date[2]))
	}
	return deadline
}

func checkDay(day int, err error) (bool, int) {
	if err != nil || day < 1 || day > 31 {
		return true, 0
	}
	return false, day
}

func checkMonth(day int, monthStr string) bool {
	// checkDay returns day = 0 if invalid day
	if day != 0 {
		month, err := strconv.Atoi(monthStr)
		if err != nil || month < 1 || month > 12 {
			return true
		}
		longerMonths := []int{1, 3, 5, 7, 8, 10, 12}
		normalMonths := []int{4, 6, 9, 11}
		if month == 2 && (day != 28 && day != 29) {
			return true
		} else if sort.SearchInts(longerMonths, month) >= len(longerMonths) {
			return true
		} else if sort.SearchInts(normalMonths, month) >= len(normalMonths) {
			return true
		}
		return false
	}
	return true
}

func checkYear(year int, err error) bool {
	if err != nil || year < 1 || year > 9999 {
		return true
	}
	return false
}

func getValidDependency() string {
	validIDs := models.GetActiveTaskIDs()
	var dependency int

	fmt.Println("List of active tasks:")
	models.ShowTasksByStatus("none", 1, 2)
	fmt.Println("")
	fmt.Println("Insert ID of task depended on:")
	fmt.Scan(&dependency)
	invalidDependency := sort.SearchInts(validIDs, dependency) >= len(validIDs)

	for invalidDependency {
		fmt.Println("Invalid active task ID. Insert valid active task ID:")
		fmt.Scan(&dependency)
		invalidDependency = sort.SearchInts(validIDs, dependency) >= len(validIDs)
	}

	return strconv.Itoa(dependency)
}
