package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	serverLogPath = "/opt/AxxonSoft/AxxonNext/Logs"
	clientLogPath = "%s/.local/share/AxxonSoft/AxxonNext/Logs"
	extractLogAppPath = "/opt/AxxonSoft/AxxonNext/bin/support"
	dumpDstPath = "./"
	confFilePath = "/opt/AxxonSoft/AxxonNext/instance.conf"

	clientName = "UILauncher"

	restartServer = "axxon-next restart"
	stopServer = "axxon-next stop"
	startServer = "axxon-next start"
)


func ClearLogs() tea.Msg {
	msg := clearDir(serverLogPath)
	if msg != Successfully {
		return msg
	}
	dir, err := getHomeDir()
	if err != nil {
		return createErrMsg(err)
	}

	path := fmt.Sprintf(clientLogPath, dir)
	
	return clearDir(path)
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

func RestartSrv() tea.Msg {
	cmd := exec.Command("service", strings.Split(restartServer, " ")...)
	if err := cmd.Run(); err != nil {
		fmt.Print(err.Error())
		return createErrMsg(err)
	}
	
	return Successfully
}

func StopSrv() tea.Msg {
	cmd := exec.Command("service", strings.Split(stopServer, " ")...)
	if err := cmd.Run(); err != nil {
		return createErrMsg(err)
	}
	
	return Successfully
}

func StartSrv() tea.Msg {
	cmd := exec.Command("service", strings.Split(startServer, " ")...)
	if err := cmd.Run(); err != nil {
		return createErrMsg(err)
	}

	return Successfully
}

func SwitchToDebug() tea.Msg {
	file, err := os.OpenFile(confFilePath, os.O_RDWR, 0)
	if err != nil {
		return createErrMsg(err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return createErrMsg(err)
	}

	buf = bytes.Replace(buf, []byte("INFO"), []byte("DEBUG"), 1)

	if err = file.Truncate(0); err != nil {
		return createErrMsg(err)
	}
	if _, err = file.Seek(0, 0); err != nil {
		return createErrMsg(err)
	}

	if _, err = file.Write(buf); err != nil {
		return createErrMsg(err)
	}

	return Successfully
}

