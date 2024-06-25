package cmd

import (
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	serverLogPath = "/opt/AxxonSoft/AxxonNext/Logs"
	clientLogPath = "/home/asker/.local/share/AxxonSoft/AxxonNext/Logs"
	extractLogAppPath = "/opt/AxxonSoft/AxxonNext/bin/support"
	dumpDstPath = "./"

	clientName = "UILauncher"
)


func ClearLogs() tea.Msg {
	msg := clearDir(serverLogPath)
	if msg != Successfully {
		return msg
	}
	// return clearDir(clientLogPath)
	return msg
}

func ExctractLogs() tea.Msg {
	cmd := exec.Command(extractLogAppPath, dumpDstPath)
	if err := cmd.Run(); err != nil {
		return createErrMsg(err)
	}

	return Successfully
}

func KillUI() tea.Msg {
	cmd := exec.Command("pidof", clientName)

	out, err := cmd.Output()
	if err != nil {
		return createErrMsg(err)
	}

	pid := strings.TrimSpace(string(out))

	cmd = exec.Command("kill", pid)
	if err = cmd.Run(); err != nil {
		return createErrMsg(err)
	}
	
	return Successfully
}