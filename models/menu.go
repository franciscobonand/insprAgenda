package models

import (
	"fmt"
)

// InitiateTaskManager prints a welcoming message to the user
func InitiateTaskManager() {
	fmt.Println("Brief description of the scheduler")
}

// DisplayMainMenu displays the main actions a user can do
func DisplayMainMenu() {
	fmt.Println("\nMAIN MENU\n" +
		"Please select one of the following:\n" +
		"1- Visualize task board\n" +
		"2- Manage tasks\n" +
		"3- Filter tasks\n" +
		"4- Show deliveries calendar\n" +
		"0- Exit\n")
}

// DisplayVisualizationMenu displays options for visualizing the board
func DisplayVisualizationMenu() {
	fmt.Println("\nVISUALIZATION MENU\n" +
		"Tasks will be separetad by their status\n" +
		"Select desired ordering:\n" +
		"1- By Deadline\n" +
		"2- By Priority\n" +
		"3- By Added time\n" +
		"Type anything else to return to the Main Menu\n")
}

// DisplayManagementMenu displays main actions to manage the board
func DisplayManagementMenu() {
	fmt.Println("\nMANAGEMENT MENU\n" +
		"Select desired action:\n" +
		"1- Create task\n" +
		"2- Remove task\n" +
		"3- Update task status\n" +
		"4- Show task details\n" +
		"Type anything else to return to the Main Menu\n")
}

// DisplayFilterMenu displays options for showing tasks by filter
func DisplayFilterMenu() {
	fmt.Println("\nFILTER  MENU\n" +
		"Tasks will be separetad by their status\n" +
		"Select desired filter:\n" +
		"1- By Deadline\n" +
		"2- By Priority\n" +
		"3- By Added time\n" +
		"Type anything else to return to the Main Menu\n")
}

// DisplayCalendar shows a calendar with active task's deadlines and estimated conclusion time
func DisplayCalendar() {
	fmt.Println("Calendar with To Do's and Working tasks")
}
