package main

import (
	"fmt"
	"insprTaskScheduler/insprAgenda/models"
	"os"
)

func main() {
	models.InitiateTaskManager()

	for {
		models.DisplayMainMenu()
		var userCommand int
		fmt.Scan(&userCommand)

		switch userCommand {
		case 1:
			models.DisplayBoardsMenu()
		case 2:
			models.DisplayTaskManagerMenu()
		case 3:
			models.DisplayCalendar()
		case 0:
			fmt.Println("Shutting down...")
			os.Exit(0)

		default:
			fmt.Println("Unknown command, please try again!")
		}
	}
}
