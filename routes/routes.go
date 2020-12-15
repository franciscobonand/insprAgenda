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
		controllers.ShowTasksByOrder(1)
	case 2: // Priority
		controllers.ShowTasksByOrder(2)
	case 3: // Added time
		controllers.ShowTasksByOrder(3)
	default:
		fmt.Println("Returning to Main Menu...")
	}
}

// HandleFilterChoice handles the input given by the user in this menu
func HandleFilterChoice() {
	var userCommand int
	fmt.Scan(&userCommand)

	switch userCommand {
	case 1: // Deadline
		controllers.ShowFilteredOptions(1)
	case 2: // Priority
		controllers.ShowFilteredOptions(2)
	case 3: // Added time
		controllers.ShowFilteredOptions(3)
	default:
		fmt.Println("Returning to Main Menu...")
	}
}
