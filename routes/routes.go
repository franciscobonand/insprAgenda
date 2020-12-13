package routes

import (
	"fmt"
	"insprTaskScheduler/insprAgenda/controllers"
)

// HandleManagementChoice handles the input given by the user
func HandleManagementChoice() {
	var userCommand int
	fmt.Scan(&userCommand)

	switch userCommand {
	case 1: // Create
		controllers.GenerateNewTask()
	case 2: // Remove
		controllers.MoveTaskOnBoard(true)
	case 3: // Update status
		controllers.MoveTaskOnBoard(false)
	case 4: // Show details

	default:
		fmt.Println("Returning to Main Menu...")
	}
}
