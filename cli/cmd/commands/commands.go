package commands

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	serverLogPath = "/opt/AxxonSoft/AxxonNext/Logs"
	clientLogPath = "/home/asker/.local/share/AxxonSoft/AxxonNext/Logs"
	extractLogAppPath = "/opt/AxxonSoft/AxxonNext/bin/support"
	logDstPath = "/home/asker/"
)


func ClearLogs() tea.Msg {
	var err error
	if err = os.RemoveAll(serverLogPath); err != nil {
		return Unsuccessfully
	}
	if err = os.RemoveAll(clientLogPath); err != nil {
		return Unsuccessfully
	}
	
	return Successfully
}

func ExctractLogs() tea.Msg {
	cmd := exec.Command(extractLogAppPath, logDstPath)
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
		return Unsuccessfully
	}

	return Successfully
}