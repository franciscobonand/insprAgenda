package routes

import (
	"fmt"
	"insprTaskScheduler/insprAgenda/controllers"
)

// HandleManagementChoice handles the input given by the user in this menu
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
		controllers.ShowTaskDetails()
	default:
		fmt.Println("Returning to Main Menu...")
	}
}

// HandleVisualizationChoice handles the input given by the user in this menu
func HandleVisualizationChoice() {
	var userCommand int
	fmt.Scan(&userCommand)

	switch userCommand {
	case 1: // Deadline
		controllers.ShowTasksByFilter(1)
	case 2: // Priority
		controllers.ShowTasksByFilter(2)
	case 3: // Added time
		controllers.ShowTasksByFilter(3)
	default:
		fmt.Println("Returning to Main Menu...")
	}
}
