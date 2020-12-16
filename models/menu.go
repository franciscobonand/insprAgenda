package models

import (
	"fmt"
)

// InitiateTaskManager prints a welcoming message to the user
func InitiateTaskManager() {
	fmt.Println("\nWelcome to Inspr's Task Manager v1.0!\n" +
		"It's higly recomended to read the README file for\n" +
		"more details on how to use it and implemented functionalities")
}

// DisplayMainMenu displays the main actions a user can do
func DisplayMainMenu() {
	fmt.Println("      _________________" +
		"\n     |    MAIN MENU    |\n" +
		"Please select one of the following:\n" +
		"1- Visualize task board\n" +
		"2- Manage tasks\n" +
		"3- Filter tasks\n" +
		// "4- Show deliveries calendar\n" +
		"0- Exit\n")
}

// DisplayVisualizationMenu displays options for visualizing the board
func DisplayVisualizationMenu() {
	fmt.Println("       __________________" +
		"\n      |VISUALIZATION MENU|\n" +
		"Tasks will be separetad by their status\n" +
		"Select desired ordering:\n" +
		"1- By Deadline\n" +
		"2- By Priority\n" +
		"3- By Added time\n" +
		"Type anything else to return to the Main Menu\n")
}

// DisplayManagementMenu displays main actions to manage the board
func DisplayManagementMenu() {
	fmt.Println("      _______________" +
		"\n     |MANAGEMENT MENU|\n" +
		"Select desired action:\n" +
		"1- Create task\n" +
		"2- Remove task\n" +
		"3- Update task status\n" +
		"4- Show task details\n" +
		"Type anything else to return to the Main Menu\n")
}

// DisplayFilterMenu displays options for showing tasks by filter
func DisplayFilterMenu() {
	fmt.Println("         ____________" +
		"\n        |FILTER  MENU|\n" +
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
