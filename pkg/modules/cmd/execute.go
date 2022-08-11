package cmd

import (
	"fmt"
	"os/exec"
	"syscall"
)

func parseCommandOutput(command *exec.Cmd) (result string) {
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		result = fmt.Sprintf("An error has occured while executing the command... Error: %s", err)
	}
	result = string(commandOutput)
	return
}

func Arkor(command string) (result string) {
	executable := `C:\windows\system32\cmd.exe`
	commandToExecute := fmt.Sprintf("cmd.exe /S /C %s", command)
	args := []string{commandToExecute}
	executeCommand := exec.Command(executable, args...)
	// Hide cmd window
	executeCommand.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	executeCommand.SysProcAttr.CmdLine = args[0]
	result = parseCommandOutput(executeCommand)
	return
}
