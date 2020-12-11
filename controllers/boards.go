package controllers

import (
	"os"
	"path/filepath"
)

// SetupBoardEnviroment checks for existence of boards file and csv's.
// If they don't exist, creates them
func SetupBoardEnviroment() {
	programPath, err := os.Getwd()

	if err != nil {
		panic("Error while setting up the enviroment:" + err.Error())
	}

	programPath = filepath.Join(programPath, "boards")

	_, errPath := os.Stat(programPath)
	if os.IsNotExist(errPath) {
		os.Mkdir("boards", 0755)
	}
}
