package models

import (
	"fmt"
	"insprTaskScheduler/insprAgenda/controllers"
)

// InitiateTaskManager prints a welcoming message to the user
func InitiateTaskManager() {
	fmt.Println("Brief description of the scheduler")
	controllers.SetupBoardEnviroment()
}

// DisplayMainMenu displays the main actions a user can do
func DisplayMainMenu() {
	fmt.Println("Show main menu, where you can:\n" +
		"1-See Boards\n" +
		"2-Manage Boards\n" +
		"3-Show callendar of deadline/time estimation\n" +
		"0-Exit\n")
}

// DisplayBoardsMenu displays options for visualizing the board
func DisplayBoardsMenu() {
	fmt.Println("Show listing options (priority/deadline/added time)")
}

// DisplayTaskManagerMenu displays main actions to manage the board
func DisplayTaskManagerMenu() {
	fmt.Println("Display menu for creating, removing, moving and showing task details")
}

// DisplayCalendar shows a calendar with active task's deadlines and estimated conclusion time
func DisplayCalendar() {
	fmt.Println("Calendar with To Do's and Working tasks")
}
