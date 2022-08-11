package file

import "os"

func CheckFileExistence(file string) (fileExists bool) {
	// Check if chrome_update.exe exists if it exists return true
	// If it doesn't exist return false
	_, err := os.Stat(file)
	if err != nil {
		fileExists = false
	} else {
		fileExists = true
	}
	return
}
