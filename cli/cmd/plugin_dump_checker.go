package cmd

import (
	"errors"
	"fmt"
	"strings"
	"t_kt/internal/telnet"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const pluginDumpPath = "/tmp/CWerkAgent/dumps/"

var pluginDumpList []string

var (
	ErrInitPluginDumpCheckerErr = errors.New("init plugin dump checker error")
	ErrPluginDumpCkeckerErr     = errors.New("plugin dump checker error")
	ErrPluginNoAuth             = errors.New("plugin telnet authenticate error")
)

func IninPluginChecker() tea.Msg {
	if len(pluginDumpList) > 0 {
		pluginDumpList = []string{}
	}

	dumps, err := getPluginDirFiles(pluginDumpPath, "", "", "", 0)
	if err != nil {
		return createErrMsg(errors.Join(ErrInitPluginDumpCheckerErr, err))
	}

	pluginDumpList = append(pluginDumpList, dumps...)

	return Successfully
}

func CheckPluginDump() tea.Msg {
	time.Sleep(time.Second * 5)
	dumps, err := getPluginDirFiles(pluginDumpPath, "", "", "", 0)
	if err != nil {
		return createErrMsg(errors.Join(ErrPluginDumpCkeckerErr, err))
	}

	if len(dumps) > len(pluginDumpList) {
		return PluginDumpDrop
	}

	return NoPluginDumps
}

func getPluginDirFiles(path, username, password, host string, port uint) ([]string, error) {
	client, _ := telnet.GetClient(host, port)
	defer client.Close()
	if client.Auth(username, password) {
		command := fmt.Sprintf("ls -l %s", path)
		client.WriteString(command)
		rawFileList, err := client.ReadString("$")
		if err != nil {
			return nil, err
		}

		files := parsePluginDumpFiles(rawFileList)

		return files, nil
	}

	return nil, ErrPluginNoAuth
}

func parsePluginDumpFiles(rawFileList string) []string {
	files := strings.Split(rawFileList, "\n")
	files = files[2 : len(files)-1]
	for idx := range files {
		line := strings.Split(files[idx], " ")
		files[idx] = line[len(line)-1]
	}

	return files
}
