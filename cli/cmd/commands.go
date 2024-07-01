package cmd

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	serverLogPath     = "/opt/AxxonSoft/AxxonNext/Logs"
	clientLogPath     = getHome()
	extractLogAppPath = "/opt/AxxonSoft/AxxonNext/bin/support"
	dumpDstPath       = "./"
	confFilePath      = "/opt/AxxonSoft/AxxonNext/instance.conf"

	clientName = "UILauncher"

	restartServer = "axxon-next restart"
	stopServer    = "axxon-next stop"
	startServer   = "axxon-next start"
)

func ClearLogs() tea.Msg {
	msg := clearDir(serverLogPath)
	if msg != Successfully {
		return msg
	}
	return clearDir(clientLogPath)
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
